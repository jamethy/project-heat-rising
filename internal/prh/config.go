package prh

import (
	"encoding/json"
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/thermostat"
	"github.com/jamethy/project-rising-heat/internal/weather"
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
	configDir, err := getConfigDir()
	if err != nil {
		fmt.Println("Failed to get config dir! Using ~/.config - ", err)
		configDir = "~/.config"
	}
	return configDir + "/project-rising-heat.json"
}

func getConfigDir() (string, error) {
	switch runtime.GOOS {
	case "linux":
		dir := os.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			usr, err := user.Current()
			if err != nil {
				return "", err
			}
			dir = filepath.Join(usr.HomeDir, ".config")
		}
		return dir, nil
	case "windows":
		dir := os.Getenv("APPDATA")
		if dir == "" {
			usr, err := user.Current()
			if err != nil {
				return "", err
			}
			dir = filepath.Join(usr.HomeDir, "AppData", "Roaming")
		}
		return dir, nil
	case "darwin": // macOS
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return filepath.Join(usr.HomeDir, "Library", "Preferences"), nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
