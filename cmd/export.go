package cmd

import (
	"fmt"
	"os"

	"github.com/daylamtayari/Microsoft-To-Do-Export/pkg/mstodo"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export task lists",
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger(cmd)
		outputType, err := cmd.Flags().GetString("type")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse output type command flag")
		}
		if outputType != "json" && outputType != "todoist" && outputType != "csv" {
			logger.Fatal().Msgf("Invalid output type %q provided", outputType)
		}
		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse completed command flag")
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse output command flag")
		}
		raw, err := cmd.Flags().GetBool("raw")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse raw command flag")
		}

		client := getClient(cmd)
		taskLists, err := client.GetAllTasks(completed)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to retrieve all tasks")
		}

		var outputContents string
		switch outputType {
		case "json":
			jsonOutput, err := jsoniter.Marshal(taskLists)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to marshal task lists to JSON")
			}
			outputContents = string(jsonOutput)
		}

		if raw {
			fmt.Print(outputContents)
		} else {
			err = os.WriteFile(outputFile, []byte(outputContents), 0644)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to write output file")
			}
			logger.Debug().Msgf("Wrote output file at %q", outputFile)
		}
	},
}

func initExportCmd() *cobra.Command {
	exportCmd.Flags().Bool("completed", false, "Include completed tasks")
	exportCmd.Flags().StringP("output", "o", "mstdexport.{file_type}", "Output file name")
	exportCmd.Flags().String("type", "json", "Output type (accepted values: 'json', 'todoist', 'csv')")
	exportCmd.Flags().BoolP("raw", "r", false, "Output to stdout instead of a file and no table")
	exportCmd.MarkFlagsMutuallyExclusive("output", "raw")
	return exportCmd
}

func exportJson(taskLists mstodo.List) string {
	return ""
}
