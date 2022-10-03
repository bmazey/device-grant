package transport

import "net/http"

func DeviceAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	// I'm a teapot (for now) ...
	w.WriteHeader(418)
}
