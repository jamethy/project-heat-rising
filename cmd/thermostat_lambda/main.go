package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
)

type AppConfig struct {
	Thermostat thermostat.Config
}

var config = AppConfig{
	Thermostat: thermostat.DefaultConfig,
}

func main() {

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal("failed to get config: ", err)
	}
	t := thermostat.New(config.Thermostat)

	lambda.Start(func(context.Context) error {

		trec, err := t.GetThermostatDBRecord(context.Background())
		if err != nil {
			return err
		}
		util.PrettyPrint(trec)
		return nil
	})
}
