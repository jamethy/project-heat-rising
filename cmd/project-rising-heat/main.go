package main

import (
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/prh"
	"github.com/spf13/cobra"
	"log"
)

func main() {

	// set in root persistent pre run
	var config prh.Config

	cmdFlags := struct {
		configPath string
	}{}

	var rootCmd = &cobra.Command{
		Use:   "project-rising-heat <command>",
		Short: "A set of commands for...",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := prh.ReadConfigFile(cmdFlags.configPath)
			if err != nil {
				return fmt.Errorf("failed to read config file: %w", err)
			}
			config = cfg
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	rootCmd.PersistentFlags().StringVar(&cmdFlags.configPath, "config-file", prh.GetDefaultConfigFilePath(), "json file with credentials and stuff")

	var installScheduleCmd = &cobra.Command{
		Use:   "install-schedule",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var pollSenseHatCmd = &cobra.Command{
		Use:   "poll-sense-hat",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var pollWeatherCmd = &cobra.Command{
		Use:   "poll-weather",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.Weather(cmd.Context(), config.DB, config.WeatherProvider)
		},
	}

	var pollThermostatCmd = &cobra.Command{
		Use:   "poll-thermostat",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.Thermostat(cmd.Context(), config.DB, config.ThermostatClient)
		},
	}

	var generateDailyDataCmd = &cobra.Command{
		Use:   "generate-daily-data",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.DailyData(cmd.Context(), config.DB, config.WeatherProvider)
		},
	}

	var versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Prints the version string",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}

	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(
		installScheduleCmd,
		pollSenseHatCmd,
		pollWeatherCmd,
		pollThermostatCmd,
		generateDailyDataCmd,
		versionCommand,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var version = "unknown" // filled in by goreleaser
