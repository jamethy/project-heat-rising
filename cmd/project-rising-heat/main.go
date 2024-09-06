package main

import (
	"fmt"
	"github.com/jamethy/project-rising-heat/internal/db"
	"github.com/jamethy/project-rising-heat/internal/prh"
	"github.com/spf13/cobra"
	"io"
	"log"
	"log/slog"
	"strings"
)

func main() {
	cmd := setupCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func setupCommand() *cobra.Command {

	cmdFlags := struct {
		configPath    string
		logLevel      string
		skipMigration bool
	}{}

	var config prh.Config // set if readConfigPreRun put in cmd preRun
	readConfigPreRun := func(cmd *cobra.Command, args []string) error {
		cfg, err := prh.ReadConfigFile(cmdFlags.configPath)
		if err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}
		config = cfg
		return nil
	}

	migrationPreRun := func(cmd *cobra.Command, args []string) error {
		if config.DB.Migrate && !cmdFlags.skipMigration {
			if err := db.Migrate(config.DB); err != nil {
				return fmt.Errorf("failed to run migrations: %w", err)
			}
		} else {
			slog.Debug("skipping migrations")
		}
		return nil
	}

	var rootCmd = &cobra.Command{
		Use:   "project-rising-heat <command>",
		Short: "A set of commands for...",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			setupLogger(cmd.OutOrStdout(), cmdFlags.logLevel)
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	rootCmd.PersistentFlags().StringVar(&cmdFlags.configPath, "config-file", prh.GetDefaultConfigFilePath(), "json file with credentials and stuff")
	rootCmd.PersistentFlags().StringVar(&cmdFlags.logLevel, "log-level", "info", "debug, info, warn, or error")
	rootCmd.Flags().BoolVar(&cmdFlags.skipMigration, "skip-migrations", false, "Skip the database migrations") // I know, it's not used everywhere, but it's fine.

	var createConfigCmd = &cobra.Command{
		Use:   "create-config",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.ReadConfigFromUserIntoFile(cmd.InOrStdin(), cmdFlags.configPath)
		},
	}

	var migrationsCmd = &cobra.Command{
		Use:     "migrations",
		Short:   "",
		PreRunE: joinPreRuns(readConfigPreRun, migrationPreRun),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var installScheduleCmd = &cobra.Command{
		Use:     "install-schedule",
		Short:   "",
		PreRunE: readConfigPreRun,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var pollSenseHatCmd = &cobra.Command{
		Use:     "poll-sense-hat",
		Short:   "",
		PreRunE: joinPreRuns(readConfigPreRun, migrationPreRun),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var pollWeatherCmd = &cobra.Command{
		Use:     "poll-weather",
		Short:   "",
		PreRunE: joinPreRuns(readConfigPreRun, migrationPreRun),
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.Weather(cmd.Context(), config.DB, config.WeatherProvider)
		},
	}

	var pollThermostatCmd = &cobra.Command{
		Use:     "poll-thermostat",
		Short:   "",
		PreRunE: joinPreRuns(readConfigPreRun, migrationPreRun),
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.Thermostat(cmd.Context(), config.DB, config.ThermostatClient)
		},
	}

	var generateDailyDataCmd = &cobra.Command{
		Use:     "generate-daily-data",
		Short:   "",
		PreRunE: joinPreRuns(readConfigPreRun, migrationPreRun),
		RunE: func(cmd *cobra.Command, args []string) error {
			return prh.DailyData(cmd.Context(), config.DB, config.WeatherProvider)
		},
	}

	var versionCommand = &cobra.Command{
		Use:   "version",
		Short: "Prints the version string",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := cmd.OutOrStdout().Write([]byte(version + "\n"))
			return err
		},
	}

	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(
		createConfigCmd,
		migrationsCmd,
		installScheduleCmd,
		pollSenseHatCmd,
		pollWeatherCmd,
		pollThermostatCmd,
		generateDailyDataCmd,
		versionCommand,
	)

	return rootCmd
}

var version = "unknown" // filled in by goreleaser

// the logger is primarily used for the polling subcommands since they are run automatically
func setupLogger(w io.Writer, level string) {
	slogLevel := slog.LevelInfo
	switch strings.ToLower(level) {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	}

	h := slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slogLevel,
	})

	l := slog.New(h)
	l = l.With("app", slog.GroupValue(
		slog.String("name", "project-rising-heat"),
		slog.String("version", version),
	))
	slog.SetDefault(l)
}

// PreRunFunc copied from cobra.Command.PreRunE definition
// just defined for better clarity of joinPreRuns definition
type PreRunFunc func(cmd *cobra.Command, args []string) error

func joinPreRuns(fns ...PreRunFunc) PreRunFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, fn := range fns {
			if err := fn(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}
