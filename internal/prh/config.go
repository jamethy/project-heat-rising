package prh

import (
	"encoding/json"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/weather"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

type Config struct {
	DB               db.Config
	WeatherProvider  weather.Config
	ThermostatClient thermostat.Config
}

var DefaultConfig = Config{
	DB:               db.DefaultConfig,
	WeatherProvider:  weather.DefaultConfig,
	ThermostatClient: thermostat.DefaultConfig,
}

func ReadConfigFromUserIntoFile(i io.Reader, filePath string) error {
	// todo actually read important values from user
	config, _ := ReadConfigFile(filePath)
	// error is ignored because we want to use defaults in case of missing file

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	if err = json.NewEncoder(f).Encode(config); err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	return nil
}

func ReadConfigFile(filePath string) (Config, error) {
	config := DefaultConfig

	f, err := os.Open(filePath)
	if err != nil {
		return config, fmt.Errorf("failed to open file: %w", err)
	}

	if err = json.NewDecoder(f).Decode(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	return config, nil
}

func GetDefaultConfigFilePath() string {
	if runtime.GOOS != "linux" {
		log.Fatalln("only linux is supported")
	}

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		usr, err := user.Current()
		if err != nil {
			fmt.Println("Failed to get config dir! Using ~/.config - ", err)
			configDir = "~"
		}
		configDir = filepath.Join(usr.HomeDir, ".config")
	}

	return configDir + "/project-rising-heat.json"
}
