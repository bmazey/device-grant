package main

import (
	"crypto/rand"
	"crypto/rsa"
	"device-grant/internal/config"
	"device-grant/internal/grants/device"
	"device-grant/pkg/oauth"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// parse local config (could be added as cmd line arg)
	cfg := config.NewConfig("internal/config/local.yml")

	// create a new mux router
	router := NewRouter(cfg)

	// start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}

// NewRouter creates a new mux router with applied server, oauth, and device grant configurations
func NewRouter(cfg config.Config) *mux.Router {
	// generate an RSA key pair
	private, err := rsa.GenerateKey(rand.Reader, cfg.OAuth.RSABits)
	if err != nil {
		log.Fatal(err)
	}

	// create a simple oauth2 issuer which contains a JWT signer and matching JWKS
	// the name provided below becomes the 'iss' claim in minted access tokens
	// start time determines the 'nbf' claim
	// the TTL integer determines the lifetime of an access token in seconds
	name := "http://" + cfg.Server.Host + ":" + cfg.Server.Port
	hour, _ := time.ParseDuration(cfg.OAuth.TokenTTL)
	issuer := oauth.NewSimpleIssuer(private, name+cfg.OAuth.JWKS, time.Now(), hour)

	// create a device granter
	minutes, _ := time.ParseDuration(cfg.DeviceGrant.UserCode.TTL)
	granter := device.NewGranter(issuer, minutes, cfg.DeviceGrant.UserCode.Length, name+cfg.DeviceGrant.Registration)

	// add a default public client_id for testing, and log it to console
	client := granter.ClientStore.Create()
	log.Printf("created default public client_id: %v", client)

	// create gorilla mux router
	router := mux.NewRouter()

	// host oauth2 JWKS endpoint
	router.HandleFunc(cfg.OAuth.JWKS, issuer.JWKSHandler)

	// routes specific to RFC 8628 OAuth 2.0 Device Authorization Grant https://www.rfc-editor.org/rfc/rfc8628
	router.HandleFunc("/device", granter.RegistrationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", granter.AccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/device_authorization", granter.AuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)

	return router
}
