package device

import (
	"encoding/json"
	"log"
	"net/http"
)

// AccessTokenRequest as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.4
type AccessTokenRequest struct {
	GrantType  string `json:"grant_type,omitempty"`
	DeviceCode string `json:"device_code,omitempty"`
	ClientID   string `json:"client_id,omitempty"`
}

// AccessTokenHandler as defined in https://www.rfc-editor.org/rfc/rfc8628#section-3.4
func (g *Granter) AccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	//var request AccessTokenRequest
	//
	//err := json.NewDecoder(r.Body).Decode(&request)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//}

	w.Header().Set("Content-Type", "application/json")

	accessToken, err := g.Issuer.IssueJWT("test-subject", []string{"my-audience"})
	if err != nil {
		log.Fatal(err)
	}

	content, err := json.Marshal(accessToken)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(content)
}
