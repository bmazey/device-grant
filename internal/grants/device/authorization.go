package device

import (
	"net/http"
)

// AuthorizationResponse as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.2
type AuthorizationResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
}

// AuthorizationHandler as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.1
func (g *Granter) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
	}

	// set JSON Content-Type header
	w.Header().Set("Content-Type", "application/json")
}

func (g *Granter) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// TODO - implement
}
