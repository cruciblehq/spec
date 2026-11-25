package trust

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"testing"
)

func TestEncodeDecodeKey(t *testing.T) {
	_, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	der, err := EncodeKey(key)
	if err != nil {
		t.Fatalf("EncodeKey() = %v, want nil", err)
	}
	got, err := DecodeKey(der)
	if err != nil {
		t.Fatalf("DecodeKey() = %v, want nil", err)
	}
	if !key.Equal(got) {
		t.Error("DecodeKey() returned a different key")
	}
}

func TestDecodeKeyInvalid(t *testing.T) {
	if _, err := DecodeKey([]byte("garbage")); !errors.Is(err, ErrInvalidKey) {
		t.Errorf("DecodeKey() = %v, want %v", err, ErrInvalidKey)
	}
}

func TestDecodeKeyNotEd25519(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := DecodeKey(der); !errors.Is(err, ErrInvalidKey) {
		t.Errorf("DecodeKey() = %v, want %v", err, ErrInvalidKey)
	}
}
