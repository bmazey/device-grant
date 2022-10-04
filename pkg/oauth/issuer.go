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
	Name      string
	Signer    jose.Signer
	Keys      jose.JSONWebKeySet
	NotBefore time.Time
	TokenTTL  time.Duration
}

func NewSimpleIssuer(private *rsa.PrivateKey, name string, start time.Time, ttl time.Duration) SimpleIssuer {
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
		TokenTTL:  ttl,
	}
}

func (s *SimpleIssuer) IssueJWT(subject string, audience []string) (*AccessToken, error) {
	builder := jwt.Signed(s.Signer)

	now := time.Now()
	later := time.Unix(now.Unix()+int64(s.TokenTTL.Seconds()), 0)

	claims := jwt.Claims{
		Issuer:    s.Name,
		Subject:   subject,
		Audience:  audience,
		IssuedAt:  jwt.NewNumericDate(now),
		Expiry:    jwt.NewNumericDate(later),
		NotBefore: jwt.NewNumericDate(s.NotBefore),
		ID:        uuid.New().String(),
	}

	accessJWT, err := builder.Claims(claims).CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &AccessToken{
		JWT:       accessJWT,
		TokenType: BEARER,
		Expiry:    int64(s.TokenTTL.Seconds()),
	}, nil
}
