package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/util"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"net/http"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
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
	uri := c.config.BaseUrl + "/data/2.5/onecall"
	params.APIKey = &c.config.APIKey

	uri, err := util.AddQueryParameters(uri, params)
	if err != nil {
		return nil, fmt.Errorf("bad query params: %w", err)
	}

	res, err := ctxhttp.Get(ctx, c.client, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	defer util.SafeClose(res.Body)

	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("non-200 response: %d - %s", res.StatusCode, string(b))
	}

	var oneCall OneCall
	err = json.NewDecoder(res.Body).Decode(&oneCall)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &oneCall, nil
}

func (c *Client) GetCurrentWeather(ctx context.Context) (*db.Weather, error) {

	// checked earlier
	w, err := c.GetCurrent(ctx, GetParams{
		Lat:     &c.config.Latitude,
		Lon:     &c.config.Longitude,
		Units:   util.Ptr("imperial"),
		Exclude: util.Ptr("minutely,hourly,daily"),
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
		Units:   util.Ptr("imperial"),
		Exclude: util.Ptr("minutely,hourly,daily"),
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
