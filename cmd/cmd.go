package cmd

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type contextKey string

var (
	loggerContextKey = contextKey("logger")
)

var rootCmd = &cobra.Command{
	Use:   "mstodoexport",
	Short: "Microsoft To Do Export",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger := zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse debug command flag")
		}
		if debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		cmd.SetContext(context.WithValue(context.TODO(), loggerContextKey, &logger))
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger := zerolog.New(os.Stderr).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		logger.Fatal().Err(err).Msg("Root command encountered an error")
	}
}

func getLogger(cmd *cobra.Command) *zerolog.Logger {
	return cmd.Context().Value(loggerContextKey).(*zerolog.Logger)
}
