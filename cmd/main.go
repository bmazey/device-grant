package main

import (
	"device-grant/internal/config"
	"device-grant/internal/transport"
	"log"
	"net/http"
)

func main() {
	// parse local config (could be added as cmd line arg)
	cfg := config.NewConfig("internal/config/local.yml")

	// create a new mux router
	router, granter := transport.NewRouter(cfg)

	// add a default public client_id for testing and log it to console
	client := granter.ClientStore.Create()
	log.Printf("created default public client_id: %v", client)

	// start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}
