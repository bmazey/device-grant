package device

import (
	"encoding/json"
	"log"
	"net/http"
)

// AuthorizationResponse as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.2
type AuthorizationResponse struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int64  `json:"expires_in"`
}

// AuthorizationHandler as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.1
func (g *Granter) AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// start by checking to see if the client_id exists
	id := r.FormValue("client_id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check to see if the client is known (public clients only - no client confidential)
	if !g.ClientStore.Contains(id) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create an interaction
	i := g.CreateInteraction(id)

	// prepare response
	response := AuthorizationResponse{
		DeviceCode:              i.DeviceCode,
		UserCode:                i.UserCode,
		VerificationURI:         g.VerificationURI,
		VerificationURIComplete: g.VerificationURI + "?user_code=" + i.UserCode,
		ExpiresIn:               int64(g.CodeTTL.Seconds()),
	}

	content, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(content)
}
