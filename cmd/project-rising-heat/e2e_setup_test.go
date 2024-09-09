package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/prh"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ory/dockertest/v3/docker"
)

var testdb *sql.DB
var config prh.Config
var ctx = context.Background()

func TestMain(m *testing.M) {
	beingRunInDocker := os.Getenv("IN_DOCKER") == "1"

	config = prh.DefaultConfig

	mockCarrierAPI := newCarrierMockAPIServer()
	defer mockCarrierAPI.Close()
	config.ThermostatClient.Carrier.Username = "carrier_username"
	config.ThermostatClient.Carrier.Password = "carrier_password"
	config.ThermostatClient.Carrier.BaseUrl = mockCarrierAPI.URL

	config.DB.Name = "e2e_test"
	config.DB.SSLDisable = true
	config.DB.Username = "e2e_username"
	config.DB.Password = "e2e_secret"
	if beingRunInDocker {
		config.DB.Host = "host.docker.internal"
	}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("Could not connect to docker", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16",
		Env: []string{
			"POSTGRES_PASSWORD=" + config.DB.Password,
			"POSTGRES_USER=" + config.DB.Username,
			"POSTGRES_DB=" + config.DB.Name,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatal("Could not start resource", err)
	}
	ptemp := resource.GetPort("5432/tcp")
	config.DB.Port, _ = strconv.Atoi(ptemp)

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	log.Println("Connecting to database on url: " + databaseURL)

	_ = resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		log.Println("pinging database")
		testdb, err = sql.Open("pgx", databaseURL)
		if err != nil {
			log.Println(err)
			return err
		}
		err = testdb.Ping()
		if err != nil {
			log.Println(err)
		}
		return err
	}); err != nil {
		log.Fatal("Could not connect to postgres", err)
	}

	// find way to project root directory
	for {
		if _, err := os.Stat("cmd"); err == nil {
			break
		}
		if err := os.Chdir(".."); err != nil {
			panic(err)
		}
	}

	// execute setup sql
	//c, err := os.ReadFile(filepath.Join("internal", "testdata", ".sql"))
	//if err != nil {
	//	log.Fatal("could not read test sql file", err)
	//}

	//log.Println("Executing test sql file: integration_test_init.sql...")
	//_, err = testdb.Exec(string(c))
	//if err != nil {
	//	log.Fatal("could not execute test sql file", err)
	//}

	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatal("Could not purge resource", err)
	}

	os.Exit(code)
}

// stolen from cobra library
func executeSubCommand(args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd := setupCommand(func(filePath string) (prh.Config, error) {
		// todo assert filePath or at least check value
		return config, nil
	})
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err = cmd.Execute()
	return buf.String(), err
}

// creates a http handler that just returns the json of a file
// fileName assumes relative to test_data and that the tests are run in repo root (since TestMain navigates there)
func loadJSONFileHandler(fileName string) http.HandlerFunc {
	fileName = filepath.Join("cmd", "project-rising-heat", "test_data", fileName)

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(b)
	}
}

func assertWithinASecond(t *testing.T, expected, actual time.Time) {
	assert.Truef(t, expected.After(actual.Add(-500*time.Millisecond)), "expected is too early: %s", expected.Sub(actual))
	assert.Truef(t, expected.Before(actual.Add(500*time.Millisecond)), "expected is too late: %s", expected.Sub(actual))
}

func assertWithinDelta(t *testing.T, expected, actual, delta float64) {
	assert.Greaterf(t, expected, actual-delta, "expected was too less than actual")
	assert.Lessf(t, expected, actual+delta, "expected was too greater than actual")
}
