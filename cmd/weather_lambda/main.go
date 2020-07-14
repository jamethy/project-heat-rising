package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/task"
	"github.com/jamethy/project-rising-heat/internal/weather"
)

type AppConfig struct {
	DB      db.Config
	Weather weather.Config
}

var config = AppConfig{
	DB:      db.DefaultConfig,
	Weather: weather.DefaultConfig,
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

	w := weather.New(config.Weather)

	lambda.Start(func(ctx context.Context) error {
		return task.Weather(ctx, d, w)
	})
}

