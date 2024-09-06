package db

import (
	"embed"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/util"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embeddedMigrations embed.FS

// Migrate the database using the scripts in "migrations"
func Migrate(dbConfig Config) error {
	d, err := Connect(dbConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer util.SafeClose(d)

	_ = goose.SetDialect("postgres")
	goose.SetBaseFS(embeddedMigrations)

	if err := goose.Up(d, "migrations", goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("during goose migration: %w", err)
	}
	return nil
}
