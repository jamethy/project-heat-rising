package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/volatiletech/sqlboiler/boil"
)

type AppConfig struct {
	DB db.Config
}

var config = AppConfig{
	DB: db.DefaultConfig,
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

	lambda.Start(func(ctx context.Context, event events.SQSEvent) error {

		errs := make([]error, 0)

		for _, record := range event.Records {
			if err := processSqsRecord(ctx, d, record); err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) > 0 {
			messages := make([]string, 0, len(errs))
			for _, e := range errs {
				messages = append(messages, e.Error())
			}
			return errors.New("failures: [" + strings.Join(messages, ", ") + "]")
		}
		return nil
	})
}

func processSqsRecord(ctx context.Context, d *sql.DB, msg events.SQSMessage) error {
	var dbRecord db.Upstair
	err := json.NewDecoder(strings.NewReader(msg.Body)).Decode(&dbRecord)
	if err != nil {
		return fmt.Errorf("failed to unmarshal sqs event: %w", err)
	}

	err = dbRecord.Insert(ctx, d, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to write to db: %w", err)
	}

	util.PrettyPrint(dbRecord)
	return nil
}
