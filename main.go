package main

import (
	"context"
	"log"

	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"
)

// todo use grafana

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

	//d, err := db.Connect(config.DB)
	//if err != nil {
	//	log.Fatal("failed to connected to database: ", err)
	//}

	ctx := context.Background()

	util.PrettyPrint(config)

	w := weather.New(config.Weather)
	wrec, err := w.CreateDBRecord(ctx)
	if err != nil {
		panic(err)
	}

	//err = wrec.Insert(ctx, d, boil.Infer())
	//if err != nil {
	//	log.Fatal("failed to write to database: ", err)
	//}
	util.PrettyPrint(wrec)

	c := thermostat.New(config.Thermostat)
	trec, err := c.CreateDBRecord(context.Background())
	if err != nil {
		panic(err)
	}
	//err = trec.Insert(ctx, d, boil.Infer())
	//if err != nil {
	//	log.Fatal("failed to write to database: ", err)
	//}
	util.PrettyPrint(trec)
}
