package main

import (
	"context"
	"log"

	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/task"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"
	"github.com/volatiletech/sqlboiler/boil"
)

type AppConfig struct {
	DB         db.Config
	Thermostat thermostat.Config
	Weather    weather.Config
}

var config = AppConfig{
	DB:         db.DefaultConfig,
	Thermostat: thermostat.DefaultConfig,
	Weather:    weather.DefaultConfig,
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

	boil.DebugMode = true
	ctx := context.Background()

	util.PrettyPrint(config)

	w := weather.New(config.Weather)
	t := thermostat.New(config.Thermostat)

	err = task.Weather(ctx, d, w)
	if err != nil {
		panic(err)
	}

	err = task.Thermostat(ctx, d, t)
	if err != nil {
		panic(err)
	}

	err = task.DailyData(ctx, d, w)
	if err != nil {
		panic(err)
	}
}
