package transport

import "net/http"

func DeviceAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	// simply returning "OK" for now ...
	w.WriteHeader(200)
}
