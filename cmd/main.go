package main

import (
	"crypto/rand"
	"crypto/rsa"
	"device-grant/internal/grants/device"
	"device-grant/pkg/oauth"
	"log"
	"net/http"
	"time"

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
	// start time determines the 'nbf' claim
	// the TTL integer determines the lifetime of an access token in seconds
	hour, _ := time.ParseDuration("60m")
	issuer := oauth.NewSimpleIssuer(private, "http://127.0.0.1:8081/jwks", time.Now(), hour)

	// create a device granter
	minutes, _ := time.ParseDuration("10m")
	granter := device.NewGranter(issuer, minutes, 8)
	// TODO - add default client to granter and print it

	// create gorilla mux router
	router := mux.NewRouter()

	// host oauth2 JWKS endpoint
	router.HandleFunc("/jwks", issuer.JWKSHandler)

	// routes specific to RFC 8628 OAuth 2.0 Device Authorization Grant https://www.rfc-editor.org/rfc/rfc8628
	router.HandleFunc("/device", granter.RegistrationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", granter.AccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/device_authorization", granter.AuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)

	// start the server
	log.Fatal(http.ListenAndServe(":8081", router))
}
