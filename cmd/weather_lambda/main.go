package main

import (
	"context"
	"log"

	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"

	"github.com/aws/aws-lambda-go/lambda"
)

type AppConfig struct {
	Weather weather.Config
}

var config = AppConfig{
	Weather: weather.DefaultConfig,
}

func main() {

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal("failed to get config: ", err)
	}
	w := weather.New(config.Weather)

	lambda.Start(func(context.Context) error {

		wrec, err := w.GetWeatherDBRecord(context.Background())
		if err != nil {
			return err
		}
		util.PrettyPrint(wrec)
		return nil
	})
}
