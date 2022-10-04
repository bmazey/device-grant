package oauth

import (
	"crypto/rsa"
	"log"

	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type SimpleIssuer struct {
	Name   string
	Signer jose.Signer
	Key    jose.JSONWebKey
}

func NewSimpleIssuer(private *rsa.PrivateKey, name string) SimpleIssuer {
	// generate a uuid to serve as the kid
	kid := uuid.New().String()

	// create a JWT signer & matching JWKS from the generated pair
	signer, err := NewSigner(private)
	if err != nil {
		log.Fatal(err)
	}

	return SimpleIssuer{
		Name:   name,
		Signer: signer,
		Key:    NewJSONWebKey(private.PublicKey, kid),
	}
}

func (s *SimpleIssuer) IssueJWT(subject string, audience []string) (string, error) {
	builder := jwt.Signed(s.Signer)

	claims := jwt.Claims{
		Issuer:   s.Name,
		Subject:  subject,
		Audience: audience,
	}

	return builder.Claims(claims).CompactSerialize()
}
