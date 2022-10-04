package device

import "net/http"

// RegistrationHandler prompts the user to authorize devices on behalf of a third-party client
// normally we would require some form of basic user auth on this endpoint and present an interface of some kind
func (g *Granter) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// start by getting user code
	userCode := r.FormValue("user_code")
	if userCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// attempt to authorize device
	err := g.AuthorizeDevice(userCode)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
