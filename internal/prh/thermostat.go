package prh

import (
	"context"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"log"
	"log/slog"

	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Thermostat(ctx context.Context, dbConfig db.Config, thermostatClientConfig thermostat.Config) error {

	d, err := db.Connect(dbConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer util.SafeClose(d)

	t := thermostat.New(thermostatClientConfig)

	trec, err := t.CreateDBRecord(ctx)
	if err != nil {
		return err
	}

	err = trec.Insert(ctx, d, boil.Infer())
	if err != nil {
		log.Fatal("failed to write to database: ", err)
	}

	slog.Info("thermostat data recorded", "data", trec)
	return nil
}
