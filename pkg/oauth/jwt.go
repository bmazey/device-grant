package oauth

import (
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2"
)

const (
	JWT    = "JWT"
	KID    = "kid"
	BEARER = "bearer"
)

type AccessToken struct {
	JWT       string `json:"access_token"`
	TokenType string `json:"token_type"`
	Expiry    int64  `json:"expires_in"`
}

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
