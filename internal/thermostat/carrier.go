package thermostat

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/util/ctxhttp"
	"github.com/jamethy/project-rising-heat/internal/util/ptr"
)

type (
	CarrierRequest struct {
		Selection interface{} `json:"selection"`
	}

	Selection struct {
		SelectionType               string `json:"selectionType"`
		SelectionMatch              string `json:"selectionMatch"`
		IncludeEvents               bool   `json:"includeEvents"`
		IncludeProgram              bool   `json:"includeProgram"`
		IncludeSettings             bool   `json:"includeSettings"`
		IncludeRuntime              bool   `json:"includeRuntime"`
		IncludeAlerts               bool   `json:"includeAlerts"`
		IncludeWeather              bool   `json:"includeWeather"`
		IncludeExtendedRuntime      bool   `json:"includeExtendedRuntime"`
		IncludeLocation             bool   `json:"includeLocation"`
		IncludeHouseDetails         bool   `json:"includeHouseDetails"`
		IncludeNotificationSettings bool   `json:"includeNotificationSettings"`
		IncludeTechnician           bool   `json:"includeTechnician"`
		IncludePrivacy              bool   `json:"includePrivacy"`
		IncludeVersion              bool   `json:"includeVersion"`
		IncludeOemCfg               bool   `json:"includeOemCfg"`
		IncludeSecuritySettings     bool   `json:"includeSecuritySettings"`
		IncludeSensors              bool   `json:"includeSensors"`
		IncludeUtility              bool   `json:"includeUtility"`
		IncludeAudio                bool   `json:"includeAudio"`
	}

	CarrierParams struct {
		Format    string `url:"format"`
		JSON      string `url:"json"`
		Timestamp string `url:"_timestamp"`
	}

	CarrierSummary struct {
		ThermostatCount *int     `json:"thermostatCount"`
		RevisionList    []string `json:"revisionList"`
		StatusList      []string `json:"statusList"`
		Status          *struct {
			Code    *int    `json:"code"`
			Message *string `json:"message"`
		} `json:"status"`
	}

	SummarySelection struct {
		SelectionType          *string `json:"selectionType"` // registered
		IncludeEquipmentStatus *bool   `json:"includeEquipmentStatus"`
	}

	CarrierResponse struct {
		Page struct {
			Page       int `json:"page"`
			TotalPages int `json:"totalPages"`
			PageSize   int `json:"pageSize"`
			Total      int `json:"total"`
		} `json:"page"`
		ThermostatList []struct {
			Identifier     string `json:"identifier"`
			Name           string `json:"name"`
			ThermostatRev  string `json:"thermostatRev"`
			IsRegistered   bool   `json:"isRegistered"`
			ModelNumber    string `json:"modelNumber"`
			Brand          string `json:"brand"`
			Features       string `json:"features"`
			LastModified   string `json:"lastModified"`
			ThermostatTime string `json:"thermostatTime"`
			UtcTime        string `json:"utcTime"`
			Location       struct {
				TimeZoneOffsetMinutes int    `json:"timeZoneOffsetMinutes"`
				TimeZone              string `json:"timeZone"`
				IsDaylightSaving      bool   `json:"isDaylightSaving"`
				StreetAddress         string `json:"streetAddress"`
				City                  string `json:"city"`
				ProvinceState         string `json:"provinceState"`
				Country               string `json:"country"`
				PostalCode            string `json:"postalCode"`
				PhoneNumber           string `json:"phoneNumber"`
				MapCoordinates        string `json:"mapCoordinates"`
			} `json:"location"`

			Runtime struct {
				RuntimeRev         string `json:"runtimeRev"`
				Connected          bool   `json:"connected"`
				FirstConnected     string `json:"firstConnected"`
				ConnectDateTime    string `json:"connectDateTime"`
				DisconnectDateTime string `json:"disconnectDateTime"`
				LastModified       string `json:"lastModified"`
				LastStatusModified string `json:"lastStatusModified"`
				RuntimeDate        string `json:"runtimeDate"`
				RuntimeInterval    int    `json:"runtimeInterval"`
				ActualTemperature  int    `json:"actualTemperature"`
				ActualHumidity     int    `json:"actualHumidity"`
				RawTemperature     int    `json:"rawTemperature"`
				ShowIconMode       int    `json:"showIconMode"`
				DesiredHeat        int    `json:"desiredHeat"`
				DesiredCool        int    `json:"desiredCool"`
				DesiredHumidity    int    `json:"desiredHumidity"`
				DesiredDehumidity  int    `json:"desiredDehumidity"`
				DesiredFanMode     string `json:"desiredFanMode"`
				DesiredHeatRange   []int  `json:"desiredHeatRange"`
				DesiredCoolRange   []int  `json:"desiredCoolRange"`
			} `json:"runtime"`
			Weather struct {
				Timestamp      string `json:"timestamp"`
				WeatherStation string `json:"weatherStation"`
				Forecasts      []struct {
					WeatherSymbol    int    `json:"weatherSymbol"`
					DateTime         string `json:"dateTime"`
					Condition        string `json:"condition"`
					Temperature      int    `json:"temperature"`
					Pressure         int    `json:"pressure"`
					RelativeHumidity int    `json:"relativeHumidity"`
					Dewpoint         int    `json:"dewpoint"`
					Visibility       int    `json:"visibility"`
					WindSpeed        int    `json:"windSpeed"`
					WindGust         int    `json:"windGust"`
					WindDirection    string `json:"windDirection"`
					WindBearing      int    `json:"windBearing"`
					Pop              int    `json:"pop"`
					TempHigh         int    `json:"tempHigh"`
					TempLow          int    `json:"tempLow"`
					Sky              int    `json:"sky"`
				} `json:"forecasts"`
			} `json:"weather"`
		} `json:"thermostatList"`
		Status struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
	}

	CarrierLogin struct {
		Username string `url:"userName" env:"CARRIER_USERNAME"`
		Password string `url:"password" env:"CARRIER_PASSWORD"`
	}

	CarrierConfig struct {
		CarrierLogin
		BaseUrl string        `env:"CARRIER_BASE_URL"`
		Timeout time.Duration `env:"CARRIER_TIMEOUT"`
	}

	carrier struct {
		config  CarrierConfig
		cookies []*http.Cookie
		client  *http.Client
	}
)

func (c *carrier) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Println("Authorizing request...")
	if c.cookies == nil {
		// todo check expiration
		fmt.Println("Existing auth cookie not found, reauthenticating")
		var err error
		c.cookies, err = c.Login(req.Context(), c.config.CarrierLogin)
		if err != nil {
			return nil, err
		}
	}

	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
		if cookie.Name == "AUTHZ_TOKEN_COOKIE" {
			req.Header.Set("Authorization", "Bearer "+cookie.Value)
		}
	}

	return http.DefaultTransport.RoundTrip(req)
}

func newCarrierClient(config CarrierConfig) *carrier {

	c := &carrier{
		config: config,
	}
	client := &http.Client{
		Timeout:   config.Timeout,
		Transport: c,
	}
	c.client = client
	return c
}

func (c *CarrierConfig) IsValid() bool {
	return c.Username != "" && c.Password != ""
}

func (c *carrier) Login(ctx context.Context, login CarrierLogin) ([]*http.Cookie, error) {
	uri := c.config.BaseUrl + "/login/"

	uri, err := util.AddQueryParameters(uri, login)
	if err != nil {
		return nil, fmt.Errorf("failed to add query parameters: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Add("referer", "https://www.carrier.com/residential/en/us/for-owners/controller-remote-access/")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	h := *c.client
	h.Transport = http.DefaultTransport
	res, err := ctxhttp.Do(ctx, &h, req)
	if err != nil {
		return nil, fmt.Errorf("failed to make authentication call: %w", err)
	}
	defer util.SafeClose(res.Body)
	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		return nil, fmt.Errorf("bad response code: %d - %s", res.StatusCode, string(body))
	}

	return res.Cookies(), nil
}

func (c *carrier) GetThermostat(ctx context.Context, id string) (*CarrierResponse, error) {

	uri := c.config.BaseUrl + "/api/1/thermostat"

	b, err := json.Marshal(&CarrierRequest{
		Selection: Selection{
			SelectionType:               "thermostats",
			SelectionMatch:              id,
			IncludeEvents:               false,
			IncludeProgram:              false,
			IncludeSettings:             false,
			IncludeRuntime:              true,
			IncludeAlerts:               false,
			IncludeWeather:              false,
			IncludeExtendedRuntime:      false,
			IncludeLocation:             false,
			IncludeHouseDetails:         false,
			IncludeNotificationSettings: false,
			IncludeTechnician:           false,
			IncludePrivacy:              false,
			IncludeVersion:              false,
			IncludeOemCfg:               false,
			IncludeSecuritySettings:     false,
			IncludeSensors:              false,
			IncludeUtility:              false,
			IncludeAudio:                false,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	var therm CarrierResponse
	_, err = ctxhttp.Get(ctxhttp.GetParams{
		Ctx:        ctx,
		HttpClient: c.client,
		URL:        uri,
		Query: &CarrierParams{
			Format:    "json",
			JSON:      string(b),
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		},
		Destination: &therm,
	})
	if err != nil {
		return &therm, fmt.Errorf("failed to get: %w", err)
	}

	return &therm, nil
}

func (c *carrier) GetSummary(ctx context.Context) (*CarrierSummary, error) {
	uri := c.config.BaseUrl + "/api/1/thermostatSummary"

	b, err := json.Marshal(CarrierRequest{
		Selection: SummarySelection{
			SelectionType:          ptr.Str("registered"),
			IncludeEquipmentStatus: ptr.True(),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	params := CarrierParams{
		Format:    "json",
		JSON:      string(b),
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}

	var summary CarrierSummary
	_, err = ctxhttp.Get(ctxhttp.GetParams{
		Ctx:         ctx,
		HttpClient:  c.client,
		URL:         uri,
		Query:       &params,
		Destination: &summary,
	})
	return &summary, err
}

func (c *carrier) GetThermostatDBRecord(ctx context.Context) (*DBRecord, error) {

	summary, err := c.GetSummary(ctx)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(summary.StatusList[0], ":")
	thermostatId, statuses := parts[0], parts[1]

	var isHeating, isCooling bool
	if statuses != "" {
		isHeating = false                                   // todo
		isCooling = strings.Contains(statuses, "compCool1") // todo
	}

	data, err := c.GetThermostat(ctx, thermostatId)
	if err != nil {
		return nil, err
	}

	t := data.ThermostatList[0].Runtime

	return &DBRecord{
		Provider:     "Carrier",
		ThermostatId: thermostatId,
		ActualTemp:   c.toF(t.ActualTemperature),
		Humidity:     float64(t.ActualHumidity),
		TargetCool:   c.toF(t.DesiredCool),
		TargetHeat:   c.toF(t.DesiredHeat),
		IsHeating:    isHeating,
		IsCooling:    isCooling,
		Timestamp:    time.Now(),
	}, nil
}

func (c *carrier) toF(carrierTemp int) float64 {
	return float64(carrierTemp) / 10.0
}
