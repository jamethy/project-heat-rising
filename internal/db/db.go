package db

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/boil"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type (
	Config struct {
		Host       string `json:"host"`
		Port       int    `json:"port"`
		Name       string `json:"name"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		SSLDisable bool   `json:"sslDisable"`
		LogSQL     bool   `json:"logSql"`
		Migrate    bool   `json:"migrate"`
	}
)

var DefaultConfig = Config{
	Host:       "localhost",
	Port:       5432,
	Name:       "localdb",
	Username:   "localuser",
	Password:   "secret",
	SSLDisable: false,
	LogSQL:     false,
	Migrate:    true,
}

func Connect(config Config) (*sql.DB, error) {

	// if running db locally, just disable ssl for convenience
	if config.Host == "127.0.0.1" {
		config.SSLDisable = true
	}

	params := []string{
		fmt.Sprintf("host=%s", config.Host),
		fmt.Sprintf("port=%d", config.Port),
		fmt.Sprintf("dbname=%s", config.Name),
		fmt.Sprintf("user=%s", config.Username),
		fmt.Sprintf("password=%s", config.Password),
	}

	if config.SSLDisable {
		params = append(params, "sslmode=disable")
	}

	connStr := strings.Join(params, " ")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	db.SetMaxOpenConns(1)
	boil.DebugMode = config.LogSQL

	return db, nil
}
