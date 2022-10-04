package oauth

import (
	"crypto/rsa"

	"gopkg.in/square/go-jose.v2"
)

const (
	JWT = "JWT"
	KID = "kid"
)

func NewSigner(private *rsa.PrivateKey, kid string) (jose.Signer, error) {
	// create signing key
	key := jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       private,
	}

	// specify JSON Web Token
	opts := jose.SignerOptions{}
	opts.WithType(JWT)
	opts.WithHeader(KID, kid)

	return jose.NewSigner(key, &opts)
}
