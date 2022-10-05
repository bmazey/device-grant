package internal

import (
	"device-grant/internal/config"
	"device-grant/internal/grants/device"
	"device-grant/internal/transport"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"testing"
)

// TestSuccessfulDeviceGrantFlow generates a random public client id, attempts to authorize a device, and produce a jwt
func TestSuccessfulDeviceGrantFlow(t *testing.T) {
	// parse test config
	cfg := config.NewConfig("config/test.yml")

	w := httptest.NewRecorder()
	router, granter := transport.NewRouter(cfg)

	// we need to add a new client_id for testing purposes
	client := granter.ClientStore.Create()

	// let's test first HTTP interaction, public client wants to request device authorization
	request := httptest.NewRequest("POST", "/device_authorization?client_id="+client, nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	log.Print(w.Body)

	response := device.AuthorizationResponse{}
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.Nil(t, err)

	// now we can simulate a user authorizing a device
	request = httptest.NewRequest("POST", "/device?user_code="+response.UserCode, nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)

	// device should now be trusted at this point, let's generate an access token using our original client_id
	request = httptest.NewRequest("POST", "/access_token?grant_type="+device.TYPE+"&client_id="+client+
		"&device_code="+response.DeviceCode, nil)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	log.Print(w.Body)
}

// TestMismatchedDeviceCodes server responds with HTTP 403 FORBIDDEN when a client presents incorrect device code
func TestMismatchedDeviceCodes(t *testing.T) {
	// TODO - implement ...
}
