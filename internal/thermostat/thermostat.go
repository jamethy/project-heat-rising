package thermostat

import (
	"context"
	"fmt"
	"time"

	"github.com/jamethy/project-rising-heat/internal/db"
)

type (
	Provider interface {
		CreateDBRecord(ctx context.Context) (*db.Thermostat, error)
	}

	Config struct {
		Carrier CarrierConfig `json:"carrier"`
	}

	Client struct {
		providers []Provider
	}
)

var DefaultConfig = Config{
	Carrier: CarrierConfig{
		BaseUrl: "https://www.myhome.carrier.com/home",
		Timeout: 30 * time.Second,
	},
}

func New(config Config) (c Client) {
	if config.Carrier.Username != "" && config.Carrier.Password != "" {
		c.providers = append(c.providers, newCarrierClient(config.Carrier))
	}
	return c
}

func (c *Client) CreateDBRecord(ctx context.Context) (*db.Thermostat, error) {
	for _, p := range c.providers {
		if r, err := p.CreateDBRecord(ctx); err != nil {
			fmt.Printf("problem getting thermostat: %s", err)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no valid thermostats")
}
