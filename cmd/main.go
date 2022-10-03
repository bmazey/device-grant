package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"oauth2/internal/transport"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/authorization", transport.DeviceAuthorizationHandler)
	router.HandleFunc("/access_token", transport.DeviceAccessTokenHandler)

	// create and host web server
	svr := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8081",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(svr.ListenAndServe())
}
