package weather

import (
	"context"
	"fmt"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/weather/openweather"
)

type (
	Provider interface {
		GetCurrentWeather(ctx context.Context) (*db.Weather, error)
		CreateDailyDBRecord(ctx context.Context) (*db.DailyDatum, error)
	}

	Config struct {
		OpenWeather openweather.Config `json:"openWeather"`
	}
	Client struct {
		config    Config
		providers []Provider
	}
)

var DefaultConfig = Config{
	OpenWeather: openweather.Config{
		Timeout: 30 * time.Second,
	},
}

func New(config Config) (c Client) {
	c.config = config
	if config.OpenWeather.IsValid() {
		c.providers = append(c.providers, openweather.New(config.OpenWeather))
	}
	return c
}

func (c *Client) GetCurrentWeather(ctx context.Context) (*db.Weather, error) {
	for _, p := range c.providers {
		if r, err := p.GetCurrentWeather(ctx); err != nil {
			fmt.Printf("problem getting weather: %s", err)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no valid weather providers")
}

func (c *Client) CreateDailyDBRecord(ctx context.Context) (*db.DailyDatum, error) {
	for _, p := range c.providers {
		if r, err := p.CreateDailyDBRecord(ctx); err != nil {
			fmt.Printf("problem getting daily record: %s", err)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no valid daily providers")
}
