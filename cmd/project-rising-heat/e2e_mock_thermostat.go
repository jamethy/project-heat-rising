package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)

// Mock server to use as Carrier thermostat API during unit tests.
// Currently just has hooks, but may want to expand to default responses.
func newCarrierMockAPIServer() *httptest.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /login/", func(writer http.ResponseWriter, request *http.Request) {
		carrierAPITestHooks.onLogin(writer, request)
	})

	mux.HandleFunc("POST /token", func(writer http.ResponseWriter, request *http.Request) {
		carrierAPITestHooks.onToken(writer, request)
	})

	mux.HandleFunc("GET /api/1/thermostat", func(writer http.ResponseWriter, request *http.Request) {
		carrierAPITestHooks.onThermostat(writer, request)
	})

	mux.HandleFunc("GET /api/1/thermostatSummary", func(writer http.ResponseWriter, request *http.Request) {
		carrierAPITestHooks.onThermostatSummary(writer, request)
	})

	return httptest.NewServer(mux)
}

// CarrierAPITestHooks can be used to validate data is correctly sent to the mock server
// these functions are nil by default, so if there is an unexpected NPE then you might unexpectedly be calling something
type CarrierAPITestHooks struct {
	onLogin             func(http.ResponseWriter, *http.Request)
	onToken             func(http.ResponseWriter, *http.Request)
	onThermostat        func(http.ResponseWriter, *http.Request)
	onThermostatSummary func(http.ResponseWriter, *http.Request)
}

var carrierAPITestHooks CarrierAPITestHooks

func resetCarrierAPITestHooks() {
	carrierAPITestHooks = CarrierAPITestHooks{}
}

// separate function instead of json file because we only care about the cookies
func generateCarrierLoginResponse(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "portalRememberMe",
		Value:   "deleteMe",
		Domain:  "www.myhome.carrier.com",
		Path:    "/",
		MaxAge:  0,
		Expires: time.Now().Add(-4 * time.Hour),
		Secure:  true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "AUTHZ_TOKEN_COOKIE",
		Value:    "7oUlAauthztokencookieLJP3oY",
		Domain:   "www.myhome.carrier.com",
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "REFRESH_TOKEN_COOKIE",
		Value: "QXfSnFdefinitelyarefreshcookiehKEgUo9mpsnG",
	})
}
