package task

import (
	"context"
	"database/sql"
	"log"

	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/jamethy/project-rising-heat/internal/weather"
	"github.com/volatiletech/sqlboiler/boil"
)

func Weather(ctx context.Context, d *sql.DB, w weather.Client) error {
	wrec, err := w.CreateDBRecord(ctx)
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
