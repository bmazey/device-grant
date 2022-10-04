package oauth

import (
	"crypto/rsa"
	"gopkg.in/square/go-jose.v2"
	"log"
	"net/http"
)

const RS256 = "RS256"

func NewJSONWebKey(public rsa.PublicKey, kid string) jose.JSONWebKey {
	return jose.JSONWebKey{
		Key:       public,
		KeyID:     kid,
		Algorithm: RS256,
	}
}

func (s *SimpleIssuer) JWKSHandler(w http.ResponseWriter, r *http.Request) {
	response, err := s.Key.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
