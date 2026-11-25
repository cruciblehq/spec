package trust

import (
	"crypto/ed25519"
	"crypto/x509"

	"github.com/cruciblehq/utils-go/crex"
)

// Encodes key in PKCS #8 DER form.
func EncodeKey(key ed25519.PrivateKey) ([]byte, error) {
	der, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidKey, err)
	}
	return der, nil
}

// Decodes a PKCS #8 DER Ed25519 private key.
func DecodeKey(der []byte) (ed25519.PrivateKey, error) {
	parsed, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, crex.Wrap(ErrInvalidKey, err)
	}
	key, ok := parsed.(ed25519.PrivateKey)
	if !ok {
		return nil, crex.Newf(ErrInvalidKey, "key has type %T where an Ed25519 key is expected", parsed)
	}
	return key, nil
}
