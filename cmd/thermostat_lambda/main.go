package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/task"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
)

type AppConfig struct {
	DB         db.Config
	Thermostat thermostat.Config
}

var config = AppConfig{
	DB:         db.DefaultConfig,
	Thermostat: thermostat.DefaultConfig,
}

func main() {

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal("failed to get config: ", err)
	}

	d, err := db.Connect(config.DB)
	if err != nil {
		log.Fatal("failed to connected to database: ", err)
	}

	t := thermostat.New(config.Thermostat)

	lambda.Start(func(ctx context.Context) error {
		return task.Thermostat(ctx, d, t)
	})
}

