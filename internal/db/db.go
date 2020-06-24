package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

type Config struct {
	URL      string `env:"DATABASE_URL"`
	Username string `env:"DATABASE_USERNAME"`
	Password string `env:"DATABASE_PASSWORD"`
	LogSQL   bool   `env:"DATABASE_LOG_SQL"`
}

var DefaultConfig = Config{
	LogSQL: false,
}

func Connect(config Config) (*sql.DB, error) {

	connStr := "postgres://" + config.Username + ":" + config.Password + "@" + config.URL

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	db.SetMaxOpenConns(1)
	boil.DebugMode = config.LogSQL

	return db, nil
}
