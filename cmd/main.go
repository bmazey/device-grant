package main

import (
	"crypto/rand"
	"crypto/rsa"
	"device-grant/internal/transport"
	"device-grant/pkg/oauth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// generate an RSA key pair
	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	issuer := oauth.NewSimpleIssuer(private)

	// create gorilla mux router
	router := mux.NewRouter()

	// standard oauth2 JWKS endpoint
	router.HandleFunc("/jwks", jwks.JWKSHandler)

	// routes specific to RFC 8628 OAuth 2.0 Device Authorization Grant https://www.rfc-editor.org/rfc/rfc8628
	router.HandleFunc("/device_authorization", transport.DeviceAuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", transport.DeviceAccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8081", router))
}
