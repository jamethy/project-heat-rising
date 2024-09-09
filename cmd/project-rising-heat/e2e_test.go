package main

import (
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"testing"
	"time"
)

func TestVersionCommand(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		version = "test-version"

		res, err := executeSubCommand("version")

		assert.NoError(t, err)
		assert.Equal(t, res, version+"\n")
	})
}

func TestPollThermostatCommand(t *testing.T) {
	t.Run("auth", func(t *testing.T) {
		defer resetCarrierAPITestHooks()
		var capturedLoginParams = struct {
			username string
			password string
			headers  http.Header
		}{}
		carrierAPITestHooks.onLogin = func(writer http.ResponseWriter, request *http.Request) {
			capturedLoginParams.username = request.URL.Query().Get("userName")
			capturedLoginParams.password = request.URL.Query().Get("password")
			capturedLoginParams.headers = request.Header
			generateCarrierLoginResponse(writer, request)
		}

		var capturedTokenParams = struct {
			grantType string
			code      string
			headers   http.Header
		}{}
		carrierAPITestHooks.onToken = func(writer http.ResponseWriter, request *http.Request) {
			capturedTokenParams.grantType = request.URL.Query().Get("grant_type")
			capturedTokenParams.code = request.URL.Query().Get("code")
			capturedTokenParams.headers = request.Header
			loadJSONFileHandler("sample-carrier-token-response.json")(writer, request)
		}

		var capturedAPICallParams = struct {
			headers http.Header
		}{}
		carrierAPITestHooks.onThermostatSummary = func(writer http.ResponseWriter, request *http.Request) {
			capturedAPICallParams.headers = request.Header
			// no need for actual response here since we don't need to capture anything else
		}

		// don't really care what happens
		_, _ = executeSubCommand("poll-thermostat")

		assert.Equal(t, "carrier_username", capturedLoginParams.username)
		assert.Equal(t, "carrier_password", capturedLoginParams.password)
		assert.Equal(t, "https://www.carrier.com/residential/en/us/for-owners/controller-remote-access/", capturedLoginParams.headers.Get("referer"))
		assert.Equal(t, "application/x-www-form-urlencoded", capturedLoginParams.headers.Get("Content-Type"))

		assert.Equal(t, "refresh_token", capturedTokenParams.grantType)
		assert.Equal(t, "QXfSnFdefinitelyarefreshcookiehKEgUo9mpsnG", capturedTokenParams.code)
		assert.Equal(t, "https://www.myhome.carrier.com/carrier/consumerportal/index.html", capturedTokenParams.headers.Get("referer"))
		assert.Equal(t, "application/x-www-form-urlencoded", capturedTokenParams.headers.Get("Content-Type"))

		assert.Equal(t, "access definitelyanaccesstoken", capturedAPICallParams.headers.Get("Authorization"))
	})

	t.Run("happy-path", func(t *testing.T) {
		defer resetCarrierAPITestHooks()
		// happy path assumes everything sent to carrier is correct and the responses are good
		carrierAPITestHooks.onLogin = generateCarrierLoginResponse
		carrierAPITestHooks.onToken = loadJSONFileHandler("sample-carrier-token-response.json")
		carrierAPITestHooks.onThermostat = loadJSONFileHandler("sample-carrier-thermostat-response.json")
		carrierAPITestHooks.onThermostatSummary = loadJSONFileHandler("sample-carrier-thermostat-summary-response.json")

		dbCountBefore, _ := db.Thermostats().Count(ctx, testdb)

		_, err := executeSubCommand("poll-thermostat")
		assert.NoError(t, err)

		dbCountAfter, _ := db.Thermostats().Count(ctx, testdb)
		assert.Equal(t, dbCountBefore+1, dbCountAfter)

		dbRecord, _ := db.Thermostats(qm.OrderBy(db.ThermostatColumns.ID+" desc")).One(ctx, testdb)
		assertWithinASecond(t, dbRecord.CreatedAt, time.Now())
		assertWithinASecond(t, dbRecord.Timestamp, time.Now())
		assert.Equal(t, "Carrier", dbRecord.Provider)
		assert.Equal(t, "900000000000", dbRecord.ThermostatID)
		assertWithinDelta(t, dbRecord.TargetCool, 73.0, 0.01)
		assertWithinDelta(t, dbRecord.TargetHeat, 65.0, 0.01)
		assertWithinDelta(t, dbRecord.ActualTemp, 72.9, 0.01)
		assertWithinDelta(t, dbRecord.Humidity, 49, 0.1)
		assert.False(t, dbRecord.IsHeating)
		assert.True(t, dbRecord.IsCooling)
	})

	// todo test errors on auth, errors on thermostat endpoints, missing data, etc
}
