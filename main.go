package main

import (
	"context"
	"log"

	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"
)

// todo use grafana

type AppConfig struct {
	Thermostat thermostat.Config
	Weather    weather.Config
}

var config = AppConfig{
	Thermostat: thermostat.DefaultConfig,
	Weather:    weather.DefaultConfig,
}

func main() {

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal("failed to get config: ", err)
	}

	util.PrettyPrint(config)

	w := weather.New(config.Weather)
	wrec, err := w.GetWeatherDBRecord(context.Background())
	if err != nil {
		panic(err)
	}
	util.PrettyPrint(wrec)

	c := thermostat.New(config.Thermostat)
	trec, err := c.GetThermostatDBRecord(context.Background())
	if err != nil {
		panic(err)
	}
	util.PrettyPrint(trec)

}
