package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jamesburns-rts/go-env"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/volatiletech/sqlboiler/boil"
)

type AppConfig struct {
	Lambda    string
	DB        db.Config
	OutputDir string
}

var config = AppConfig{
	DB:        db.DefaultConfig,
	OutputDir: "output",
}

func main() {
	ctx := context.Background()

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal("failed to get config: ", err)
	}

	d, err := db.Connect(config.DB)
	if err != nil {
		log.Fatal("failed to connected to database: ", err)
	}

	if config.Lambda == "TRUE" {
		lambda.Start(func(ctx context.Context, event events.SQSEvent) error {

			errs := make([]error, 0)

			for _, record := range event.Records {
				if err := processRecord(ctx, d, strings.NewReader(record.Body)); err != nil {
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
	} else {
		if err := syncFiles(ctx, d); err != nil {
			log.Println("Error with output files: ", err)
		}
	}
}

func syncFiles(ctx context.Context, d *sql.DB) error {
	files, err := ioutil.ReadDir(config.OutputDir)
	if err != nil {
		return fmt.Errorf("failed to read output dir: %s - %w", config.OutputDir, err)
	}
	for _, file := range files {
		filePath := config.OutputDir + "/" + file.Name()
		f, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to read output file: %s - %w", file.Name(), err)
		}
		err = processRecord(ctx, d, f)
		if err != nil {
			return fmt.Errorf("failed to process output file: %s - %w", file.Name(), err)
		}
		util.SafeClose(f)

		err = os.Remove(filePath)
		if err != nil {
			return fmt.Errorf("failed to delete output file: %s - %w", file.Name(), err)
		}
	}
	return nil
}

func processRecord(ctx context.Context, d *sql.DB, r io.Reader) error {
	var dbRecord db.Upstair
	err := json.NewDecoder(r).Decode(&dbRecord)
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
