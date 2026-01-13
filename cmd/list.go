package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List task lists",
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger(cmd)
		jsonOutput, err := cmd.Flags().GetBool("json")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse json command flag")
		}
		rawOutput, err := cmd.Flags().GetBool("raw")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse raw command flag")
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse output command flag")
		}

		client := getClient(cmd)
		lists, err := client.GetLists()
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to retrieve lists")
		}

		if cmd.Flags().Changed("output") {
			var outputData []byte
			if jsonOutput {
				outputData, err = jsoniter.Marshal(lists)
				if err != nil {
					logger.Fatal().Err(err).Msg("Failed to marshal lists to JSON")
				}
			} else {
				var outputBuilder strings.Builder
				outputBuilder.WriteString("Microsoft To Do Lists")
				for _, l := range lists {
					outputBuilder.WriteString("\n" + l.DisplayName)
				}
				outputData = []byte(outputBuilder.String())
			}
			err = os.WriteFile(outputFile, outputData, 0644)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to write output file")
			}
			logger.Debug().Msgf("Wrote output file at %q", outputFile)
			return
		}

		if jsonOutput {
			outputData, err := jsoniter.MarshalIndent(lists, "", "    ")
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to marshal lists to JSON")
			}
			fmt.Print(string(outputData))
		} else if rawOutput {
			for i := range lists {
				if i == 0 {
					fmt.Print(lists[i].DisplayName)
				} else {
					fmt.Print("\n" + lists[i].DisplayName)
				}
			}
		} else {
			t := table.NewWriter()
			t.SetStyle(table.StyleRounded)
			t.AppendHeader(table.Row{"Lists"})
			for i := range lists {
				t.AppendRow(table.Row{lists[i].DisplayName})
			}
			fmt.Printf("%s\n", t.Render())
		}
	},
}

func initListCmd() *cobra.Command {
	listCmd.Flags().BoolP("json", "j", false, "JSON output")
	listCmd.Flags().BoolP("raw", "r", false, "Raw output to stdout with no table")
	listCmd.Flags().StringP("output", "o", "", "Output file name")
	return listCmd
}
