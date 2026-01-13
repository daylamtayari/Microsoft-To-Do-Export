package cmd

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"

	"github.com/daylamtayari/Microsoft-To-Do-Export/pkg/mstodo"
	"github.com/spf13/cobra"
)

type attachmentTask struct {
	ListId      string
	TaskId      string
	Attachments []mstodo.Attachment
}

var attachmentCmd = &cobra.Command{
	Use:   "attachment",
	Short: "Retrieve task attachment",
	Run: func(cmd *cobra.Command, args []string) {
		logger := getLogger(cmd)
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse all command flag")
		}
		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse completed command flag")
		}
		outputDir, err := cmd.Flags().GetString("output")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse output command flag")
		}
		listId, err := cmd.Flags().GetString("list")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse list command flag")
		}
		taskId, err := cmd.Flags().GetString("task")
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to parse task command flag")
		}

		if !all {
			if listId == "" {
				logger.Fatal().Msg("List ID value cannot be empty if not using all")
			}
			if taskId == "" {
				logger.Fatal().Msg("Task ID value cannot be empty if not using all")
			}
		}

		err = os.Mkdir(outputDir, 0755)
		if err != nil && !errors.Is(err, os.ErrExist) {
			logger.Fatal().Err(err).Msg("Failed to create output directory")
		}

		client := getClient(cmd)

		attachmentTasks := make([]attachmentTask, 0)
		if all {
			taskLists, err := client.GetAllTasks(completed)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to retrieve all tasks")
			}
			for i := range taskLists {
				for j := range taskLists[i].Tasks {
					if taskLists[i].Tasks[j].HasAttachments {
						attachmentTasks = append(attachmentTasks, attachmentTask{
							ListId: taskLists[i].Id,
							TaskId: taskLists[i].Tasks[j].Id,
						})
					}
				}
			}
		} else {
			attachmentTasks = append(attachmentTasks, attachmentTask{
				ListId: listId,
				TaskId: taskId,
			})
		}

		for i := range attachmentTasks {
			attachmentTasks[i].Attachments, err = client.ListAttachments(attachmentTasks[i].ListId, attachmentTasks[i].TaskId)
			if err != nil {
				logger.Fatal().Err(err).Msgf("Failed to list attachments for task with ID %q", taskId)
			}
		}

		for i := range attachmentTasks {
			attachOutputDir := outputDir
			if all {
				// Use the last 40 characters of the IDs as a way to handle Windows file path limits
				// Using the last 40 characters as the first characters are identical across lists and tasks
				attachOutputDir = filepath.Join(outputDir, (attachmentTasks[i].ListId)[len(attachmentTasks[i].ListId)-40:], (attachmentTasks[i].TaskId)[len(attachmentTasks[i].TaskId)-40:])
				err = os.MkdirAll(attachOutputDir, 0755)
				if err != nil && !errors.Is(err, os.ErrExist) {
					logger.Fatal().Err(err).Msgf("Failed to create task specific output directory %q", attachOutputDir)
				}
			}

			for j := range attachmentTasks[i].Attachments {
				// Don't update the original attachmentTask as this new attachment
				// variable is the only one to contain the contents of the file
				// and it is written directly to disk, helping limit memory usage
				attachment, err := client.GetAttachment(
					attachmentTasks[i].ListId,
					attachmentTasks[i].TaskId,
					attachmentTasks[i].Attachments[j].Id,
				)
				if err != nil {
					logger.Fatal().Err(err).Msgf("Failed to retrieve attachment with ID %q", attachmentTasks[i].Attachments[j].Id)

				}

				decodedContents, err := base64.StdEncoding.DecodeString(attachment.ContentBytes)
				if err != nil {
					logger.Error().Err(err).Msgf("Failed to decode the contents of attachment with ID %q", attachment.Id)
					continue
				}
				err = os.WriteFile(filepath.Join(attachOutputDir, attachment.Name), decodedContents, 0644)
				if err != nil {
					logger.Error().Err(err).Msgf("Failed to write the contents of attachment with ID %q", attachment.Id)
					continue
				}
				logger.Info().Msgf("Exported attachment %q", attachment.Name)
			}
		}
	},
}

func initAttachmentsCmd() *cobra.Command {
	attachmentCmd.Flags().String("list", "", "ID of the task list")
	attachmentCmd.Flags().String("task", "", "ID of the task to retrieve attachments for")
	attachmentCmd.Flags().Bool("completed", false, "Include attachments from completed tasks")
	attachmentCmd.Flags().Bool("all", false, "Get attachments for all tasks")
	attachmentCmd.Flags().StringP("output", "o", "mstodo_attachments", "Output directory of attachments (if all attachments retrieved, subdirectories will be automatically created)")
	attachmentCmd.Flags().Bool("human", false, "Use list and task names for the nested file paths instead of IDs (NOTE: If illegal characters are present or the length of the names are too long, it can fail to create the directories and files)")
	attachmentCmd.MarkFlagsRequiredTogether("list", "task")
	attachmentCmd.MarkFlagsMutuallyExclusive("list", "all")
	attachmentCmd.MarkFlagsMutuallyExclusive("list", "completed")
	attachmentCmd.MarkFlagsMutuallyExclusive("list", "human")
	return attachmentCmd
}
