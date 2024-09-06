package prh

import (
	"context"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"log"

	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Weather(ctx context.Context, dbConfig db.Config, weatherConfig weather.Config) error {

	d, err := db.Connect(dbConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer util.SafeClose(d)

	w := weather.New(weatherConfig)

	wrec, err := w.GetCurrentWeather(ctx)
	if err != nil {
		return err
	}
	err = wrec.Insert(ctx, d, boil.Infer())
	if err != nil {
		log.Fatal("failed to write to database: ", err)
	}
	util.PrettyPrint(wrec)
	return nil
}
