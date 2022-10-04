package oauth

import (
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2"
)

const JWT = "JWT"

func NewSigner(private *rsa.PrivateKey) (jose.Signer, error) {
	// create signing key
	key := jose.SigningKey{
		Algorithm: jose.RS256,
		Key:       private,
	}

	// specify JSON Web Token
	opts := jose.SignerOptions{}
	opts.WithType(JWT)

	return jose.NewSigner(key, &opts)
}
