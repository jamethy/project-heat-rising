package task

import (
	"context"
	"database/sql"
	"log"

	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/volatiletech/sqlboiler/boil"
)

func Thermostat(ctx context.Context, d *sql.DB, t thermostat.Client) error {
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
