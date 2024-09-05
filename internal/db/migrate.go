package db

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embeddedMigrations embed.FS

// Migrate the database using the scripts in "migrations"
func Migrate(db *sql.DB) error {
	_ = goose.SetDialect("postgres")
	goose.SetBaseFS(embeddedMigrations)

	if err := goose.Up(db, "migrations", goose.WithAllowMissing()); err != nil {
		return fmt.Errorf("during goose migration: %w", err)
	}
	return nil
}
