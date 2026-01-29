package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/daylamtayari/Microsoft-To-Do-Export/v2/pkg/mstodo"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "List tasks",
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger(cmd)
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse all command flag")
		}
		rawOutput, err := cmd.Flags().GetBool("raw")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse raw command flag")
		}
		justNames, err := cmd.Flags().GetBool("just-names")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse just-names command flag")
		}
		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse completed command flag")
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse output command flag")
		}
		listId, err := cmd.Flags().GetString("list")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse list command flag")
		}
		listName, err := cmd.Flags().GetString("name")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse name command flag")
		}

		client := getClient(cmd)
		var lists []mstodo.List

		if cmd.Flags().Changed("list") {
			list, err := client.GetList(listId)
			if err != nil {
				logger.Fatal().Err(err).Msgf("Failed to retrieve list with ID %q", listId)
			}
			tasks, err := client.GetListTasks(list.Id, completed)
			if err != nil {
				logger.Fatal().Err(err).Msgf("Failed to get tasks for list with ID %q", list.Id)
			}
			list.Tasks = tasks
			lists = []mstodo.List{list}
		} else if cmd.Flags().Changed("name") {
			allLists, err := client.GetLists()
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to retrieve lists")
			}
			for _, list := range allLists {
				if list.DisplayName == listName {
					lists = append(lists, list)
				}
			}
			if len(lists) == 0 {
				logger.Info().Msgf("No lists were found to match the name %q", listName)
				return
			}

			for i := range lists {
				tasks, err := client.GetListTasks(lists[i].Id, completed)
				if err != nil {
					logger.Error().Err(err).Msgf("Failed toretrieve tasks for list %q", lists[i].Id)
				}
				lists[i].Tasks = tasks
			}
		} else if all {
			lists, err = client.GetAllTasks(completed)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to get all tasks")
			}
		} else {
			// Should never reach here due to thhe use of MarkFlagsOneRequired but catch all in case
			logger.Fatal().Msg("Required option of list, name, or all not specified")
		}

		if rawOutput {
			for _, list := range lists {
				for _, task := range list.Tasks {
					fmt.Print("\n" + task.Title)
				}
			}
			return
		}

		if cmd.Flags().Changed("output") {
			file, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to open output file %q", outputFile)
			}
			defer file.Close() //nolint:errcheck

			writer := bufio.NewWriter(file)
			defer writer.Flush() //nolint:errcheck

			for _, list := range lists {
				for _, task := range list.Tasks {
					_, err := writer.WriteString("\n" + task.Title + "\t" + task.Status + "\t" + task.DueDateTime.Datetime.String())
					if err != nil {
						logger.Error().Err(err).Msg("Failed to write task to output")
					}
				}
			}
		}

		for _, list := range lists {
			t := table.NewWriter()
			t.SetStyle(table.StyleRounded)
			t.SetTitle(list.DisplayName)
			if !justNames {
				t.AppendHeader(table.Row{
					"Title",
					"Status",
					"Due Date",
					"Body",
				})
			}

			for _, task := range list.Tasks {
				if justNames {
					t.AppendRow(table.Row{task.Title})
				} else {
					dueDate := ""
					if task.DueDateTime != nil {
						dueDate = task.DueDateTime.Time().String()
					}
					t.AppendRow(table.Row{
						task.Title,
						task.Status,
						dueDate,
						task.Body.Content,
					})
				}
			}

			fmt.Printf("%s\n", t.Render())
		}
	},
}

func initTaskCmd() *cobra.Command {
	taskCmd.Flags().BoolP("all", "a", false, "List out all tasks for all task lists")
	taskCmd.Flags().String("list", "", "ID of the list to retrieve tasks for")
	taskCmd.Flags().String("name", "", "Name of the list to retrieve tasks for NOTE: Lists can have the same name, this will return tasks from all lists that share the name")
	taskCmd.Flags().BoolP("raw", "r", false, "Raw listing of tasks to stdout with no table")
	taskCmd.Flags().StringP("output", "o", "t", "Output file name")
	taskCmd.Flags().Bool("just-names", false, "Only output task names")
	taskCmd.Flags().Bool("completed", false, "Include completed tasks")
	taskCmd.MarkFlagsOneRequired("list", "name", "all")
	taskCmd.MarkFlagsMutuallyExclusive("list", "name", "all")
	taskCmd.MarkFlagsMutuallyExclusive("output", "raw")
	return taskCmd
}
