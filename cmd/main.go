package main

import (
	"log"
	"net/http"
	"oauth2/internal/transport"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/device_authorization", transport.DeviceAuthorizationHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/access_token", transport.DeviceAccessTokenHandler).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8081", router))
}
