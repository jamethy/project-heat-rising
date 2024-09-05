package openweather

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
	url := c.config.BaseUrl + "/onecall"
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
	lat, _ := strconv.ParseFloat(c.config.Latitude, 10)
	lon, _ := strconv.ParseFloat(c.config.Longitude, 10)

	w, err := c.GetCurrent(ctx, GetParams{
		Lat:     &lat,
		Lon:     &lon,
		Units:   ptr.Str("imperial"),
		Exclude: ptr.Str("minutely,hourly,daily"),
	})
	if err != nil {
		return nil, err
	}

	return &db.Weather{
		Provider:           null.StringFrom("OpenWeatherMaps"),
		Temperature:        null.Float32FromPtr(w.Current.Temp),
		FeelsLike:          null.Float32FromPtr(w.Current.FeelsLike),
		Pressure:           null.Float32FromPtr(w.Current.Pressure),
		Humidity:           null.Float32FromPtr(w.Current.Humidity),
		WindSpeed:          null.Float32FromPtr(w.Current.WindSpeed),
		WindDirection:      null.Float32FromPtr(w.Current.WindDeg),
		Clouds:             null.Float32FromPtr(w.Current.Clouds),
		UvIndex:            null.Float32FromPtr(w.Current.Uvi),
		RainLevel:          null.Float32FromPtr(w.GetRainLevel()),
		WeatherDescription: null.StringFromPtr(w.GetWeatherDescription()),
		Timestamp:          time.Now(),
	}, nil
}

func (c *Client) CreateDailyDBRecord(ctx context.Context) (*db.DailyDatum, error) {

	// checked earlier
	lat, _ := strconv.ParseFloat(c.config.Latitude, 10)
	lon, _ := strconv.ParseFloat(c.config.Longitude, 10)

	w, err := c.GetCurrent(ctx, GetParams{
		Lat:     &lat,
		Lon:     &lon,
		Units:   ptr.Str("imperial"),
		Exclude: ptr.Str("minutely,hourly,daily"),
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
