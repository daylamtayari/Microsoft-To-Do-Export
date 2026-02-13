// Non-feature complete package that handles the Super Productivity JSON import format
package superproductivity

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"
)

// NewCompleteBackup creates a new complete backup with current timestamp
func NewCompleteBackup(data AppDataComplete) *CompleteBackup {
	now := time.Now().UnixMilli()
	return &CompleteBackup{
		Timestamp:         now,
		LastUpdate:        now,
		CrossModelVersion: 4.5,
		Data:              data,
	}
}

// generateID creates a random ID similar to Super Productivity's nanoid format
func generateID() string {
	// Generate 16 random bytes
	b := make([]byte, 16)
	rand.Read(b)

	// Encode to base64 URL-safe format and trim to 21 characters
	id := base64.URLEncoding.EncodeToString(b)
	id = strings.ReplaceAll(id, "=", "")
	if len(id) > 21 {
		id = id[:21]
	}
	return id
}

// AddTask adds a new task to the backup and returns the generated task ID
// The task parameter should have all desired fields set except ID and Created,
// which will be auto-generated. This method handles updating project and tag relationships.
func (b *CompleteBackup) AddTask(task Task) string {
	// Generate ID and timestamp
	task.ID = generateID()

	// Ensure required fields have values
	if task.SubTaskIDs == nil {
		task.SubTaskIDs = []string{}
	}
	if task.TimeSpentOnDay == nil {
		task.TimeSpentOnDay = make(map[string]int64)
	}
	if task.TagIDs == nil {
		task.TagIDs = []string{}
	}
	if task.Attachments == nil {
		task.Attachments = []TaskAttachment{}
	}

	taskID := task.ID

	// Add to task state
	b.Data.Task.IDs = append(b.Data.Task.IDs, taskID)
	b.Data.Task.Entities[taskID] = task

	// Add to project's task list
	if project, exists := b.Data.Project.Entities[task.ProjectID]; exists {
		project.TaskIDs = append(project.TaskIDs, taskID)
		b.Data.Project.Entities[task.ProjectID] = project
	}

	// Add to each tag's task list
	for _, tagID := range task.TagIDs {
		if tag, exists := b.Data.Tag.Entities[tagID]; exists {
			tag.TaskIDs = append(tag.TaskIDs, taskID)
			b.Data.Tag.Entities[tagID] = tag
		}
	}

	// If this is a subtask, add it to the parent's SubTaskIDs
	if task.ParentID != nil {
		if parentTask, exists := b.Data.Task.Entities[*task.ParentID]; exists {
			parentTask.SubTaskIDs = append(parentTask.SubTaskIDs, taskID)
			b.Data.Task.Entities[*task.ParentID] = parentTask
		}
	}

	return taskID
}

// AddProject adds a new project to the backup and returns the generated project ID
func (b *CompleteBackup) AddProject(title string) string {
	projectID := generateID()

	project := Project{
		ID:             projectID,
		Title:          title,
		TaskIDs:        []string{},
		BacklogTaskIDs: []string{},
		NoteIDs:        []string{},
	}

	b.Data.Project.IDs = append(b.Data.Project.IDs, projectID)
	b.Data.Project.Entities[projectID] = project

	return projectID
}

// AddTag adds a new tag to the backup and returns the tag ID
// If a tag with the same title already exists, returns the existing tag's ID
func (b *CompleteBackup) AddTag(title string) string {
	// Check if a tag with this title already exists
	for _, existingTag := range b.Data.Tag.Entities {
		if existingTag.Title == title {
			return existingTag.ID
		}
	}

	// Tag doesn't exist, create a new one
	tagID := generateID()

	tag := Tag{
		ID:      tagID,
		Title:   title,
		TaskIDs: []string{},
	}

	b.Data.Tag.IDs = append(b.Data.Tag.IDs, tagID)
	b.Data.Tag.Entities[tagID] = tag

	return tagID
}

// AddRepeatCfg adds a new task repeat configuration to the backup and returns the generated ID
// The repeatCfg parameter should have all desired fields set except ID.
func (b *CompleteBackup) AddRepeatCfg(repeatCfg TaskRepeatCfg) string {
	// Generate ID
	repeatCfg.ID = generateID()

	// Ensure required fields have values
	if repeatCfg.TagIDs == nil {
		repeatCfg.TagIDs = []string{}
	}

	repeatCfgID := repeatCfg.ID

	b.Data.TaskRepeatCfg.IDs = append(b.Data.TaskRepeatCfg.IDs, repeatCfgID)
	b.Data.TaskRepeatCfg.Entities[repeatCfgID] = repeatCfg

	return repeatCfgID
}
