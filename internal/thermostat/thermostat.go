package thermostat

import (
	"context"
	"fmt"
	"time"
)

type (
	DBRecord struct {
		Provider     string
		ThermostatId string
		ActualTemp   float64
		Humidity     float64
		TargetCool   float64
		TargetHeat   float64
		IsHeating    bool
		IsCooling    bool
		Timestamp    time.Time
	}

	Provider interface {
		GetThermostatDBRecord(ctx context.Context) (*DBRecord, error)
	}

	Config struct {
		Carrier CarrierConfig
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

func (c *Client) GetThermostatDBRecord(ctx context.Context) (*DBRecord, error) {
	for _, p := range c.providers {
		if r, err := p.GetThermostatDBRecord(ctx); err != nil {
			fmt.Printf("problem getting thermostat: %s", err)
		} else {
			return r, nil
		}
	}
	return nil, fmt.Errorf("no valid thermostats")
}
