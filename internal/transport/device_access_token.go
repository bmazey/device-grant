package transport

import (
	"encoding/json"
	"net/http"
)

// DeviceAccessTokenRequest as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.4
type DeviceAccessTokenRequest struct {
	GrantType  string `json:"grant_type,omitempty"`
	DeviceCode string `json:"device_code,omitempty"`
	ClientID   string `json:"client_id,omitempty"`
}

// DeviceAccessTokenHandler endpoint as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.4
func DeviceAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request DeviceAccessTokenRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
