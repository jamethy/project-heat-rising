package weather

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/util/ctxhttp"
	"github.com/jamethy/project-rising-heat/internal/util/ptr"
	"github.com/volatiletech/null"
)

//https://openweathermap.org/current

type OpenWeatherConfig struct {
	Latitude  string        `env:"OPEN_WEATHER_LAT"`
	Longitude string        `env:"OPEN_WEATHER_LON"`
	BaseUrl   string        `env:"OPEN_WEATHER_BASE_URL"`
	APIKey    string        `env:"OPEN_WEATHER_API_KEY"`
	Timeout   time.Duration `env:"OPEN_WEATHER_TIMEOUT"`
}

type openWeatherClient struct {
	client *http.Client
	config OpenWeatherConfig
}

type getParams struct {
	Query  *string  `url:"q,omitempty"` // e.g city name
	Lat    *float64 `url:"lat,omitempty"`
	Lon    *float64 `url:"lon,omitempty"`
	APIKey *string  `url:"appid,omitempty"`
	Units  *string  `url:"units,omitempty"` // imperial, metric, default
}

type (
	response struct {
		Coord      *coord     `json:"coord"`
		Weather    []*weather `json:"weather"`
		Base       *string    `json:"base"`
		Main       *main      `json:"main"`
		Visibility *int       `json:"visibility"`
		Wind       *wind      `json:"wind"`
		Clouds     *clouds    `json:"clouds"`
		Dt         *int       `json:"dt"`
		Sys        *sys       `json:"sys"`
		Timezone   *int       `json:"timezone"` // UTC offset in seconds
		ID         *int       `json:"id"`
		Name       *string    `json:"name"`
		Cod        *int       `json:"cod"`
	}
	coord struct {
		Lon *float32 `json:"lon"`
		Lat *float32 `json:"lat"`
	}
	weather struct {
		ID          *int    `json:"id"`
		Main        *string `json:"main"`
		Description *string `json:"description"`
		Icon        *string `json:"icon"`
	}
	main struct {
		Temp      *float32 `json:"temp"`
		FeelsLike *float32 `json:"feels_like"`
		TempMin   *float32 `json:"temp_min"`
		TempMax   *float32 `json:"temp_max"`
		Pressure  *float32 `json:"pressure"`
		Humidity  *float32 `json:"humidity"`
	}
	wind struct {
		Speed *float32 `json:"speed"`
		Deg   *float32 `json:"deg"`
	}
	clouds struct {
		All *float32 `json:"all"`
	}
	sys struct {
		Type    *int    `json:"type"`
		ID      *int    `json:"id"`
		Country *string `json:"country"`
		Sunrise *int    `json:"sunrise"`
		Sunset  *int    `json:"sunset"`
	}
)

func (w *response) GetSunrise() *time.Time {
	if w.Sys == nil || w.Sys.Sunrise == nil {
		return nil
	}
	t := time.Unix(int64(*w.Sys.Sunrise), 0)
	if w.Timezone != nil {
		t = t.In(time.FixedZone("local", *w.Timezone))
	}
	return &t
}

func (w *response) GetSunset() *time.Time {
	if w.Sys == nil || w.Sys.Sunset == nil {
		return nil
	}
	t := time.Unix(int64(*w.Sys.Sunset), 0)
	if w.Timezone != nil {
		t = t.In(time.FixedZone("local", *w.Timezone))
	}
	return &t
}

func (c *OpenWeatherConfig) IsValid() bool {
	_, latErr := strconv.ParseFloat(c.Latitude, 10)
	_, lonErr := strconv.ParseFloat(c.Longitude, 10)
	return c.APIKey != "" && c.BaseUrl != "" && latErr == nil && lonErr == nil
}

func newOpenWeatherClient(config OpenWeatherConfig) *openWeatherClient {
	return &openWeatherClient{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

func (c *openWeatherClient) GetOpenWeather(ctx context.Context, params getParams) (*response, error) {
	params.APIKey = &c.config.APIKey
	var weather response

	_, err := ctxhttp.Get(ctxhttp.GetParams{
		Ctx:         ctx,
		HttpClient:  c.client,
		URL:         c.config.BaseUrl,
		Query:       &params,
		Destination: &weather,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get weather: %w", err)
	}

	return &weather, nil
}

func (c *openWeatherClient) CreateDBRecord(ctx context.Context) (*db.Weather, error) {

	// checked earlier
	lat, _ := strconv.ParseFloat(c.config.Latitude, 10)
	lon, _ := strconv.ParseFloat(c.config.Longitude, 10)

	w, err := c.GetOpenWeather(ctx, getParams{
		Lat:   &lat,
		Lon:   &lon,
		Units: ptr.Str("imperial"),
	})
	if err != nil {
		return nil, err
	}

	return &db.Weather{
		Provider:      null.StringFrom("OpenWeatherMaps"),
		Temperature:   null.Float32FromPtr(w.Main.Temp),
		FeelsLike:     null.Float32FromPtr(w.Main.FeelsLike),
		Pressure:      null.Float32FromPtr(w.Main.Pressure),
		Humidity:      null.Float32FromPtr(w.Main.Humidity),
		WindSpeed:     null.Float32FromPtr(w.Wind.Speed),
		WindDirection: null.Float32FromPtr(w.Wind.Deg),
		Clouds:        null.Float32FromPtr(w.Clouds.All),
		Timestamp:     time.Now(),
	}, nil
}

func (c *openWeatherClient) CreateDailyDBRecord(ctx context.Context) (*db.DailyDatum, error) {

	// checked earlier
	lat, _ := strconv.ParseFloat(c.config.Latitude, 10)
	lon, _ := strconv.ParseFloat(c.config.Longitude, 10)

	w, err := c.GetOpenWeather(ctx, getParams{
		Lat:   &lat,
		Lon:   &lon,
		Units: ptr.Str("imperial"),
	})
	if err != nil {
		return nil, err
	}

	return &db.DailyDatum{
		Date:    time.Now(),
		Sunrise: null.TimeFromPtr(w.GetSunrise()),
		Sunset:  null.TimeFromPtr(w.GetSunset()),
	}, nil
}
