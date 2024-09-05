package prh

import (
	"context"
	"github.com/jamethy/project-rising-heat/internal/db"
	"log"

	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/volatiletech/sqlboiler/boil"
)

func Thermostat(ctx context.Context, dbConfig db.Config, thermostatClientConfig thermostat.Config) error {
	d, err := db.Connect(dbConfig)
	if err != nil {
		log.Fatal("failed to connected to database: ", err)
	}

	t := thermostat.New(thermostatClientConfig)

	trec, err := t.CreateDBRecord(ctx)
	if err != nil {
		return err
	}

	err = trec.Insert(ctx, d, boil.Infer())
	if err != nil {
		log.Fatal("failed to write to database: ", err)
	}

	util.PrettyPrint(trec)
	return nil
}
