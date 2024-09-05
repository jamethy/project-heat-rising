package prh

import (
	"encoding/json"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/weather"
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

func ReadConfigFile(filePath string) (Config, error) {
	config := Config{
		DB:              db.DefaultConfig,
		WeatherProvider: weather.DefaultConfig,
	}

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
