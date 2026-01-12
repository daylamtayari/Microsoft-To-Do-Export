package mstodo

import (
	"strings"
	"time"
)

// MSTime wraps time.Time to handle Microsoft Graph API datetime format
type MSTime struct {
	time.Time
}

// UnmarshalJSON handles Microsoft datetime format
// Supports both UTC format (2026-01-10T18:36:12.0646821Z) and local format (2026-01-12T06:00:00.0000000)
func (t *MSTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		*t = MSTime{}
		return nil
	}

	var parsed time.Time
	var err error

	// Try parsing with timezone
	parsed, err = time.Parse(time.RFC3339Nano, s)
	if err != nil {
		// If no timezone specified, parse as local time
		parsed, err = time.Parse("2006-01-02T15:04:05.9999999", s)
		if err != nil {
			return err
		}
	}
	*t = MSTime{parsed}
	return nil
}

// MarshalJSON outputs in RFC3339Nano format
func (t MSTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte("\"" + t.Format(time.RFC3339Nano) + "\""), nil
}

// Note body of a task
type TaskBody struct {
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

// Time field a particular task
type TaskTime struct {
	Datetime MSTime `json:"dateTime"`
	Timezone string `json:"timeZone"`
}

// Time returns the datetime in the correct timezone
// The Datetime field contains a local time, and Timezone specifies which timezone it's in
func (tt *TaskTime) Time() (time.Time, error) {
	if tt == nil {
		return time.Time{}, nil
	}

	loc, err := time.LoadLocation(tt.Timezone)
	if err != nil {
		return time.Time{}, err
	}

	year, month, day := tt.Datetime.Date()
	hour, min, sec := tt.Datetime.Clock()
	nsec := tt.Datetime.Nanosecond()

	return time.Date(year, month, day, hour, min, sec, nsec, loc), nil
}

// Recurrence pattern
type TaskRecurrencePattern struct {
	DayOfMonth     int32    `json:"dayOfMonth"`
	DaysOfWeek     []string `json:"daysOfWeek"`
	FirstDayOfWeek string   `json:"firstDayOfWeek"`
	Index          string   `json:"index"`
	Interval       int32    `json:"interval"`
	Month          int32    `json:"month"`
	Type           string   `json:"type"`
}

// Range of a recurring task
type TaskRecurrenceRange struct {
	Type                string `json:"type"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	RecurrenceTimezone  string `json:"recurrenceTimeZone"`
	NumberOfOccurrences int32  `json:"numberOfOccurrences"`
}

// Recurrence configuration for a task
type TaskRecurrence struct {
	Pattern TaskRecurrencePattern `json:"pattern"`
	Range   TaskRecurrenceRange   `json:"range"`
}

// Checklist/sub-items of a task
type TaskChecklistItems struct {
	DisplayName     string `json:"displayName"`
	CreatedDatetime MSTime `json:"createdDatetime"`
	IsChecked       bool   `json:"isChecked"`
	Id              string `json:"id"`
}

// Task object
type Task struct {
	Id                   string                `json:"id"`
	Status               string                `json:"status"`
	Title                string                `json:"title"`
	Importance           string                `json:"importance"`
	IsReminderOn         bool                  `json:"isReminderOn"`
	CreatedDateTime      MSTime                `json:"createdDateTime"`
	LastModifiedDateTime MSTime                `json:"lastModifiedDateTime"`
	HasAttachments       bool                  `json:"hasAttachments"`
	Categories           []string              `json:"categories"`
	Body                 TaskBody              `json:"body"`
	CompletedDateTime    *TaskTime             `json:"completedDateTime"`
	DueDateTime          *TaskTime             `json:"dueDateTime"`
	Recurrence           *TaskRecurrence       `json:"recurrence"`
	ReminderDateTime     *TaskTime             `json:"reminderDateTime"`
	StartDateTime        *TaskTime             `json:"startDateTime"`
	ChecklistItems       *[]TaskChecklistItems `json:"checklistItems"`
}

// List object
// Tasks field that contains all tasks of the
// list, but only non-nil when retrieved from
// GetAllTasks
type List struct {
	DisplayName       string  `json:"displayName"`
	Id                string  `json:"id"`
	IsOwner           bool    `json:"isOwner"`
	IsShared          bool    `json:"isShared"`
	WellKnownListName string  `json:"wellKnownListName"`
	Tasks             *[]Task `json:"tasks"`
}
