package mstodo

import "time"

type TaskBody struct {
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

type TaskTime struct {
	Datetime time.Time `json:"dateTime"`
	Timezone string    `json:"timeZone"`
}

type TaskRecurrencePattern struct {
	DayOfMonth     int32    `json:"dayOfMonth"`
	DaysOfWeek     []string `json:"daysOfWeek"`
	FirstDayOfWeek string   `json:"firstDayOfWeek"`
	Index          string   `json:"index"`
	Interval       int32    `json:"interval"`
	Month          int32    `json:"month"`
	Type           string   `json:"type"`
}

type TaskRecurrenceRange struct {
	Type                string `json:"type"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	RecurrenceTimezone  string `json:"recurrenceTimeZone"`
	NumberOfOccurrences int32  `json:"numberOfOccurrences"`
}

type TaskRecurrence struct {
	Pattern TaskRecurrencePattern `json:"pattern"`
	Range   TaskRecurrenceRange   `json:"range"`
}

type TaskChecklistItems struct {
	DisplayName     string    `json:"displayName"`
	CreatedDatetime time.Time `json:"createdDatetime"`
	IsChecked       bool      `json:"isChecked"`
	Id              string    `json:"id"`
}

type Task struct {
	Id                   string              `json:"id"`
	Status               string              `json:"status"`
	Title                string              `json:"title"`
	Importance           string              `json:"importance"`
	IsReminderOn         string              `json:"isReminderOn"`
	CreatedDateTime      time.Time           `json:"createdDateTime"`
	LastModifiedDateTime time.Time           `json:"lastModifiedDateTime"`
	HasAttachments       bool                `json:"hasAttachments"`
	Categories           []string            `json:"categories"`
	Body                 TaskBody            `json:"body"`
	CompletedDateTime    *TaskTime           `json:"completedDateTime"`
	DueDateTime          *TaskTime           `json:"dueDateTime"`
	Recurrence           *TaskRecurrence     `json:"recurrence"`
	ReminderDateTime     *TaskTime           `json:"reminderDateTime"`
	StartDateTime        *TaskTime           `json:"startDateTime"`
	ChecklistItems       *TaskChecklistItems `json:"checklistItems"`
}

type List struct {
	DisplayName       string `json:"displayName"`
	Id                string `json:"id"`
	IsOwner           bool   `json:"isOwner"`
	IsShared          bool   `json:"isShared"`
	WellKnownListName string `json:"wellKnownListName"`
}
