package main

import (
	"crypto/rand"
	"crypto/rsa"
	"device-grant/internal/grants/device"
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

	// create a simple oauth2 issuer which contains a JWT signer and matching JWKS
	// the name provided below becomes the 'iss' claim in minted access tokens
	issuer := oauth.NewSimpleIssuer(private, "127.0.0.1:8081/jwks")

	// create a device granter
	granter := device.NewGranter(issuer)

	// create gorilla mux router
	router := mux.NewRouter()

	// host oauth2 JWKS endpoint
	router.HandleFunc("/jwks", issuer.JWKSHandler)

	// routes specific to RFC 8628 OAuth 2.0 Device Authorization Grant https://www.rfc-editor.org/rfc/rfc8628
	router.HandleFunc("/device_authorization", granter.AuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", granter.AccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)

	// start the server
	log.Fatal(http.ListenAndServe(":8081", router))
}
