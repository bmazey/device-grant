package device

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
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
	//err := json.NewDecoder(r.Body).Decode(&request)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//}

	w.Header().Set("Content-Type", "application/json")

	jwt, err := g.Issuer.IssueJWT("test-subject", []string{"my-audience"})
	if err != nil {
		log.Fatal(err)
	}

	token := oauth2.Token{
		AccessToken: jwt,
	}

	content, err := json.Marshal(token)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(content)
}
