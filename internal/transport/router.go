package transport

import (
	"crypto/rand"
	"crypto/rsa"
	"device-grant/internal/config"
	"device-grant/internal/grants/device"
	"device-grant/pkg/oauth"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// NewRouter creates a new mux router with applied server, oauth, and device grant configurations
func NewRouter(cfg config.Config) (*mux.Router, *device.Granter) {
	// generate an RSA key pair
	private, err := rsa.GenerateKey(rand.Reader, cfg.OAuth.RSABits)
	if err != nil {
		log.Fatal(err)
	}

	// create a simple oauth2 issuer which contains a JWT signer and matching JWKS
	// the name provided below becomes the 'iss' claim in minted access tokens
	// start time determines the 'nbf' claim
	// the TTL integer determines the lifetime of an access token in seconds
	// we are using plain http here strictly for example purposes
	name := cfg.OAuth.Issuer
	hour, _ := time.ParseDuration(cfg.OAuth.TokenTTL)
	issuer := oauth.NewSimpleIssuer(private, name+cfg.OAuth.JWKS, cfg.OAuth.Audience, time.Now(), hour)

	// create a device granter
	name = cfg.Server.Host + ":" + cfg.Server.Port
	minutes, _ := time.ParseDuration(cfg.DeviceGrant.UserCode.TTL)
	granter := device.NewGranter(issuer, minutes, cfg.DeviceGrant.UserCode.Length, name+cfg.DeviceGrant.Registration)

	// create gorilla mux router
	router := mux.NewRouter()

	// host oauth2 JWKS endpoint
	router.HandleFunc(cfg.OAuth.JWKS, issuer.JWKSHandler)

	// routes specific to RFC 8628 OAuth 2.0 Device Authorization Grant https://www.rfc-editor.org/rfc/rfc8628
	router.HandleFunc("/device", granter.RegistrationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", granter.AccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/device_authorization", granter.AuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)

	return router, &granter
}
