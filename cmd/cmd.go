package cmd

import (
	"context"
	"os"
	"time"

	"github.com/daylamtayari/Microsoft-To-Do-Export/pkg/mstodo"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

type contextKey string

var (
	clientContextKey = contextKey("client")
	loggerContextKey = contextKey("logger")
)

var rootCmd = &cobra.Command{
	Use:   "mstodoexport",
	Short: "Microsoft To Do Export",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Logger handling
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
		cmd.SetContext(context.WithValue(cmd.Context(), loggerContextKey, &logger))
		// Token handling
		envToken := os.Getenv("MSTDEXPORT_TOKEN")
		if envToken != "" {
			createClient(cmd, envToken)
			logger.Debug().Msg("Using token from environment variable")
			return
		}
		tokenFlag, err := cmd.Flags().GetString("token")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse token command flag")
		}
		if tokenFlag != "" {
			createClient(cmd, tokenFlag)
			logger.Debug().Msg("Using token from token flag")
			return
		}
		tokenFile, err := cmd.Flags().GetString("token-file")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse token file command flag")
		}
		if tokenFile != "" {
			token, err := os.ReadFile(tokenFile)
			if err != nil {
				logger.Fatal().Err(err).Msgf("Failed to read token file %q")
				return
			}
			createClient(cmd, string(token))
			logger.Debug().Msg("Using token from token file")
			return
		}
		logger.Fatal().Msg("No token was provided\nPlease provide one via the MSTDEXPORT_TOKEN environment value or the token or token-file flags")
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token value")
	rootCmd.PersistentFlags().String("token-file", "", "File containing token")
	rootCmd.MarkFlagsMutuallyExclusive("token", "token-file")
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

func getClient(cmd *cobra.Command) *mstodo.Client {
	return cmd.Context().Value(clientContextKey).(*mstodo.Client)
}

func createClient(cmd *cobra.Command, token string) {
	cmd.SetContext(context.WithValue(cmd.Context(), clientContextKey, mstodo.NewClient(nil, &token)))
}
