package agent

// Requests installation of the agent's certificate.
//
// The client sends this to [EndpointCertify] to finish enrollment. The root
// signs the certificate for the life of the machine, so it is never renewed.
// The agent checks that the certificate is signed by the root, names its
// machine, and holds its key. Installing it consumes the enrollment token.
type CertifyRequest struct {
	Certificate []byte `codec:"certificate"` // DER encoded agent certificate.
}

// Confirms that the agent's certificate was installed.
type CertifyResponse struct {
}
