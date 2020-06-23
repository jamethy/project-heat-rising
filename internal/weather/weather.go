package weather

import (
	"context"
	"fmt"
	"time"
)

type (
	DBRecord struct {
		Provider      string
		Temperature   float64
		FeelsLike     float64
		Pressure      float64
		Humidity      float64
		WindSpeed     float64
		WindDirection float64
		Clouds        float64
		Timestamp     time.Time
	}

	DailyDBRecord struct {
		Date    time.Time
		Sunrise time.Time
		Sunset  time.Time
	}

	Provider interface {
		GetWeatherDBRecord(ctx context.Context) (*DBRecord, error)
	}

	Config struct {
		OpenWeather OpenWeatherConfig
	}
	Client struct {
		config    Config
		providers []Provider
	}
)

var DefaultConfig = Config{
	OpenWeather: OpenWeatherConfig{
		Timeout: 30 * time.Second,
	},
}

func New(config Config) (c Client) {
	c.config = config
	if config.OpenWeather.IsValid() {
		c.providers = append(c.providers, newOpenWeatherClient(config.OpenWeather))
	}
	return c
}

func (c *Client) GetWeatherDBRecord(ctx context.Context) (*DBRecord, error) {
	for _, p := range c.providers {
		if r, err := p.GetWeatherDBRecord(ctx); err != nil {
			fmt.Printf("problem getting weather: %s", err)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no valid weather providers")
}
