package trust

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/cruciblehq/utils-go/pki"
)

type session struct {
	client    tls.ConnectionState
	server    tls.ConnectionState
	clientErr error
	serverErr error
}

func handshake(t *testing.T, serverCfg, clientCfg *tls.Config) session {
	t.Helper()
	now := func() time.Time { return testNow }
	serverCfg.Time = now
	clientCfg.Time = now

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	var s session
	done := make(chan struct{})
	go func() {
		defer close(done)
		conn, err := ln.Accept()
		if err != nil {
			s.serverErr = err
			return
		}
		defer conn.Close()
		server := tls.Server(conn, serverCfg)
		s.serverErr = server.Handshake()
		s.server = server.ConnectionState()
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := tls.Client(conn, clientCfg)
	s.clientErr = client.Handshake()
	s.client = client.ConnectionState()
	<-done
	return s
}

func tlsCert(key ed25519.PrivateKey, chain ...*x509.Certificate) tls.Certificate {
	cert := tls.Certificate{PrivateKey: key}
	for _, c := range chain {
		cert.Certificate = append(cert.Certificate, c.Raw)
	}
	return cert
}

func selfSigned(t *testing.T) tls.Certificate {
	t.Helper()
	pub, key, _ := ed25519.GenerateKey(rand.Reader)
	return tlsCert(key, createCert(t, pki.LeafTemplate("machine", validFor(ClientLifetime), x509.ExtKeyUsageServerAuth), nil, pub, key))
}

func serverConfig(ca testCA, cert tls.Certificate, crl *x509.RevocationList) *tls.Config {
	return ServerConfig(func() *tls.Certificate { return &cert }, ca.root, func() *x509.RevocationList { return crl })
}

func TestBaseConfig(t *testing.T) {
	cfg := baseConfig()
	if cfg.MinVersion != tls.VersionTLS13 {
		t.Errorf("MinVersion = %x, want %x", cfg.MinVersion, tls.VersionTLS13)
	}
	if len(cfg.CurvePreferences) != 1 || cfg.CurvePreferences[0] != tls.X25519MLKEM768 {
		t.Errorf("CurvePreferences = %v, want [%v]", cfg.CurvePreferences, tls.X25519MLKEM768)
	}
}

func TestServerAndClientConfig(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	client, clientKey := ca.client(t)
	crl := ca.crl(t, 1)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), crl),
		ClientConfig(tlsCert(clientKey, client, ca.issuer), ca.root, crl, testMachine),
	)
	if s.serverErr != nil || s.clientErr != nil {
		t.Fatalf("handshake = (%v, %v), want (nil, nil)", s.serverErr, s.clientErr)
	}
	if s.client.Version != tls.VersionTLS13 {
		t.Errorf("Version = %x, want %x", s.client.Version, tls.VersionTLS13)
	}
	if s.client.CurveID != tls.X25519MLKEM768 {
		t.Errorf("CurveID = %v, want %v", s.client.CurveID, tls.X25519MLKEM768)
	}
}

func TestServerConfigRevokedIssuer(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	client, clientKey := ca.client(t)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), ca.crl(t, 2, ca.issuer)),
		ClientConfig(tlsCert(clientKey, client, ca.issuer), ca.root, ca.crl(t, 1), testMachine),
	)
	if !errors.Is(s.serverErr, ErrRevoked) {
		t.Errorf("server handshake = %v, want %v", s.serverErr, ErrRevoked)
	}
}

func TestServerConfigClientSignedByRoot(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	pub, clientKey, _ := ed25519.GenerateKey(rand.Reader)
	client := createCert(t, leafTemplate(x509.ExtKeyUsageClientAuth, nil), ca.root, pub, ca.rootKey)
	crl := ca.crl(t, 1)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), crl),
		ClientConfig(tlsCert(clientKey, client), ca.root, crl, testMachine),
	)
	if !errors.Is(s.serverErr, ErrInvalidChain) {
		t.Errorf("server handshake = %v, want %v", s.serverErr, ErrInvalidChain)
	}
}

func TestServerConfigClientFromOtherRoot(t *testing.T) {
	ca := newTestCA(t)
	other := newTestCA(t)
	machine, machineKey := ca.machine(t)
	client, clientKey := other.client(t)
	crl := ca.crl(t, 1)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), crl),
		ClientConfig(tlsCert(clientKey, client, other.issuer), ca.root, crl, testMachine),
	)
	if s.serverErr == nil {
		t.Error("server handshake = nil, want an error")
	}
}

func TestServerConfigNoClientCertificate(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	crl := ca.crl(t, 1)
	clientCfg := ClientConfig(tls.Certificate{}, ca.root, crl, testMachine)
	clientCfg.Certificates = nil

	s := handshake(t, serverConfig(ca, tlsCert(machineKey, machine), crl), clientCfg)
	if s.serverErr == nil {
		t.Error("server handshake = nil, want an error")
	}
}

func TestClientConfigRevokedMachine(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	client, clientKey := ca.client(t)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), ca.crl(t, 1)),
		ClientConfig(tlsCert(clientKey, client, ca.issuer), ca.root, ca.crl(t, 2, machine), testMachine),
	)
	if !errors.Is(s.clientErr, ErrRevoked) {
		t.Errorf("client handshake = %v, want %v", s.clientErr, ErrRevoked)
	}
}

func TestClientConfigOtherMachine(t *testing.T) {
	ca := newTestCA(t)
	machine, machineKey := ca.machine(t)
	client, clientKey := ca.client(t)
	crl := ca.crl(t, 1)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine), crl),
		ClientConfig(tlsCert(clientKey, client, ca.issuer), ca.root, crl, MachineID("fedcba9876543210fedcba9876543210")),
	)
	if s.clientErr == nil {
		t.Error("client handshake = nil, want an error")
	}
}

func TestClientConfigMachineSignedByIssuer(t *testing.T) {
	ca := newTestCA(t)
	pub, machineKey, _ := ed25519.GenerateKey(rand.Reader)
	machine := createCert(t, leafTemplate(x509.ExtKeyUsageServerAuth, []string{testMachine.Name()}), ca.issuer, pub, ca.issuerKey)
	client, clientKey := ca.client(t)
	crl := ca.crl(t, 1)

	s := handshake(t,
		serverConfig(ca, tlsCert(machineKey, machine, ca.issuer), crl),
		ClientConfig(tlsCert(clientKey, client, ca.issuer), ca.root, crl, testMachine),
	)
	if !errors.Is(s.clientErr, ErrInvalidChain) {
		t.Errorf("client handshake = %v, want %v", s.clientErr, ErrInvalidChain)
	}
}

func TestEnrollConfig(t *testing.T) {
	ca := newTestCA(t)
	client, clientKey := ca.client(t)

	s := handshake(t,
		serverConfig(ca, selfSigned(t), ca.crl(t, 1)),
		EnrollConfig(tlsCert(clientKey, client, ca.issuer)),
	)
	if s.serverErr != nil || s.clientErr != nil {
		t.Fatalf("handshake = (%v, %v), want (nil, nil)", s.serverErr, s.clientErr)
	}
	if len(s.client.PeerCertificates) != 1 {
		t.Errorf("len(PeerCertificates) = %d, want 1", len(s.client.PeerCertificates))
	}
}
