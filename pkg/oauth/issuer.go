package oauth

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
)

type SimpleIssuer struct {
	Signer jose.Signer
	Key    jose.JSONWebKey
}

func NewSimpleIssuer(private *rsa.PrivateKey) SimpleIssuer {
	// generate a uuid to serve as the kid
	kid := uuid.New().String()

	// create a JWT signer & matching JWKS from the generated pair
	signer, err := NewSigner(private)
	if err != nil {
		log.Fatal(err)
	}

	return SimpleIssuer{
		Signer: signer,
		Key:    NewJSONWebKey(private.PublicKey, kid),
	}
}

func (s *SimpleIssuer) IssueJWT(issuer string, subject string, audience []string) (string, error) {
	builder := jwt.Signed(s.Signer)

	claims := jwt.Claims{
		Issuer:   issuer,
		Subject:  subject,
		Audience: audience,
	}

	builder.Claims(claims)

	return builder.CompactSerialize()
}
