package openweather

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/util/ctxhttp"
	"github.com/jamethy/project-rising-heat/internal/util/ptr"
)

//https://openweathermap.org/current

type Client struct {
	client *http.Client
	config Config
}

func New(config Config) *Client {
	return &Client{
		config: config,
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

type GetParams struct {
	Query   *string  `url:"q,omitempty"` // e.g city name
	Lat     *float64 `url:"lat,omitempty"`
	Lon     *float64 `url:"lon,omitempty"`
	APIKey  *string  `url:"appid,omitempty"`
	Units   *string  `url:"units,omitempty"`   // imperial, metric, default
	Exclude *string  `url:"exclude,omitempty"` // current,minutely,hourly,daily
}

func (c *Client) GetCurrent(ctx context.Context, params GetParams) (*OneCall, error) {
	url := c.config.BaseUrl + "/data/2.5/onecall"
	params.APIKey = &c.config.APIKey
	var oneCall OneCall

	_, err := ctxhttp.Get(ctxhttp.GetParams{
		Ctx:         ctx,
		HttpClient:  c.client,
		URL:         url,
		Query:       &params,
		Destination: &oneCall,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get weather: %w", err)
	}

	return &oneCall, nil
}

func (c *Client) GetCurrentWeather(ctx context.Context) (*db.Weather, error) {

	// checked earlier
	w, err := c.GetCurrent(ctx, GetParams{
		Lat:     &c.config.Latitude,
		Lon:     &c.config.Longitude,
		Units:   ptr.Str("imperial"),
		Exclude: ptr.Str("minutely,hourly,daily"),
	})
	if err != nil {
		return nil, err
	}

	return &db.Weather{
		Provider:           "OpenWeatherMaps",
		Temperature:        w.Current.Temp,
		FeelsLike:          w.Current.FeelsLike,
		Pressure:           w.Current.Pressure,
		Humidity:           w.Current.Humidity,
		WindSpeed:          w.Current.WindSpeed,
		WindDirection:      w.Current.WindDeg,
		Clouds:             w.Current.Clouds,
		UvIndex:            w.Current.Uvi,
		RainLevel:          w.GetRainLevel(),
		WeatherDescription: w.GetWeatherDescription(),
		Timestamp:          time.Now(),
	}, nil
}

func (c *Client) CreateDailyDBRecord(ctx context.Context) (*db.DailyDatum, error) {

	w, err := c.GetCurrent(ctx, GetParams{
		Lat:     &c.config.Latitude,
		Lon:     &c.config.Longitude,
		Units:   ptr.Str("imperial"),
		Exclude: ptr.Str("minutely,hourly,daily"),
	})
	if err != nil {
		return nil, err
	}

	return &db.DailyDatum{
		Date:    time.Now(),
		Sunrise: *w.GetSunrise(), // todo check
		Sunset:  *w.GetSunset(),
	}, nil
}
