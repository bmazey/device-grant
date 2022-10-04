package oauth

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"time"
)

type SimpleIssuer struct {
	Signer    jose.Signer
	Keys      jose.JSONWebKeySet
	NotBefore time.Time
	Name      string
	TTL       int64
}

func NewSimpleIssuer(private *rsa.PrivateKey, name string, start time.Time, ttl int64) SimpleIssuer {
	// generate a uuid to serve as the kid
	kid := uuid.New().String()

	// create a JWT signer & matching JWKS from the generated pair
	signer, err := NewSigner(private, kid)
	if err != nil {
		log.Fatal(err)
	}

	return SimpleIssuer{
		Signer:    signer,
		Keys:      NewJSONWebKeySet(private.PublicKey, kid),
		NotBefore: start,
		Name:      name,
		TTL:       ttl,
	}
}

func (s *SimpleIssuer) IssueJWT(subject string, audience []string) (string, error) {
	builder := jwt.Signed(s.Signer)

	now := time.Now()
	later := time.Unix(now.Unix()+s.TTL, 0)

	claims := jwt.Claims{
		Issuer:    s.Name,
		Subject:   subject,
		Audience:  audience,
		IssuedAt:  jwt.NewNumericDate(now),
		Expiry:    jwt.NewNumericDate(later),
		NotBefore: jwt.NewNumericDate(s.NotBefore),
		ID:        uuid.New().String(),
	}

	return builder.Claims(claims).CompactSerialize()
}
