package compute

// Installs a machine certificate.
//
// crux sends this to [EndpointCertify] to finish enrollment. The root signs the
// certificate for the life of the machine, so it is never renewed. cruxd checks
// that the certificate is signed by the root, names its machine, and holds its
// key. Installing it consumes the enrollment token.
type CertifyRequest struct {
	Certificate []byte `codec:"certificate"` // DER encoded machine certificate.
}

// Confirms that the machine certificate was installed.
type CertifyResponse struct {
}
