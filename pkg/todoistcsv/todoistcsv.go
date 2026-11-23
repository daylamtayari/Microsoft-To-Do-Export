// Complete handling of the Todoist CSV import format,
// allowing for ease of creating valid import CSV files.
package todoistcsv

import (
	"encoding/csv"
	"strconv"
	"strings"
	"time"
)

// Represents a CSV export file
type Export struct {
	Meta     Meta
	Sections []Section
}

func CreateExport(meta Meta, sections []Section) *Export {
	return &Export{
		Meta:     meta,
		Sections: sections,
	}
}

func (e *Export) AddSection(section Section) {
	e.Sections = append(e.Sections, section)
}

func (e *Export) SetMeta(meta Meta) {
	e.Meta = meta
}

// Number of columns in a Todoist CSV import file
const columnCount = 14

// Returns the contents of a Todoist CSV import file
func (e *Export) CSV() string {
	var csvBuilder strings.Builder
	writer := csv.NewWriter(&csvBuilder)

	// Write header row and meta field
	writeRecord(writer, "TYPE", "CONTENT", "DESCRIPTION", "PRIORITY", "INDENT", "AUTHOR", "RESPONSIBLE", "DATE", "DATE_LANG", "TIMEZONE", "DURATION", "DURATION_UNIT", "DEADLINE", "DEADLINE_LANG")
	writeRecord(writer, "meta", string(e.Meta))
	writeRecord(writer)

	// Write all sections and tasks
	for i := range e.Sections {
		writeRecord(writer, string(SectionType), e.Sections[i].Title)
		for j := range e.Sections[i].Tasks {
			exportTask(writer, e.Sections[i].Tasks[j])
		}
	}

	writer.Flush()
	return strings.TrimSuffix(csvBuilder.String(), "\n")
}

// Writes a record padded with empty fields to the full column count.
// The csv.Writer quotes any field containing a comma, quote, or newline.
func writeRecord(writer *csv.Writer, fields ...string) {
	record := make([]string, columnCount)
	copy(record, fields)
	_ = writer.Write(record) // Writes to a strings.Builder cannot fail
}

func exportTask(writer *csv.Writer, task Task) {
	// Handling values from struct to export values
	author := ""
	if task.Author != nil {
		author = User(*task.Author).String()
	}
	responsible := ""
	if task.Responsible != nil {
		responsible = User(*task.Responsible).String()
	}
	date := ""
	if task.Date != nil {
		date = time.Time(*task.Date).Format("2006-01-02")
	}
	dateLang := ""
	if task.DateLang != nil {
		dateLang = string(*task.DateLang)
	}
	timezone := ""
	if task.Timezone != nil {
		tz := time.Location(*task.Timezone)
		timezone = tz.String()
	}
	duration := ""
	if task.Duration != nil {
		duration = strconv.Itoa(int(*task.Duration))
	}
	deadline := ""
	if task.Deadline != nil {
		deadline = time.Time(*task.Deadline).Format("2006-01-02")
	}
	deadlineLang := ""
	if task.DeadlineLang != nil {
		deadlineLang = string(*task.DeadlineLang)
	}

	// Write task
	writeRecord(writer, string(TaskType), normalizeNewlines(string(task.Title)), normalizeNewlines(string(task.Description)), strconv.Itoa(int(task.Priority)), strconv.Itoa(int(task.Subtask)), author, responsible, date, dateLang, timezone, duration, string(task.DurationUnit), deadline, deadlineLang)

	// Write all notes
	for i := range task.Notes {
		writeRecord(writer, string(NoteType), normalizeNewlines(string(task.Notes[i].Content)))
	}

	// Recursively call for any child tasks
	for i := range task.ChildTasks {
		exportTask(writer, task.ChildTasks[i])
	}
}

// Converts Windows-style line endings to plain newlines
func normalizeNewlines(s string) string {
	return strings.ReplaceAll(s, "\r\n", "\n")
}

type Section struct {
	Title string
	Tasks []Task
}

func CreateSection(title string, tasks []Task) *Section {
	return &Section{
		Title: title,
		Tasks: tasks,
	}
}

func (s *Section) AddTask(task Task) {
	s.Tasks = append(s.Tasks, task)
}

func (s *Section) AddTasks(tasks []Task) {
	s.Tasks = append(s.Tasks, tasks...)
}

type Task struct {
	Title        Content
	Description  Description
	Priority     Priority
	Subtask      Indent
	Author       *Author
	Responsible  *Responsible
	Date         *Date
	DateLang     *DateLang
	Timezone     *Timezone
	Duration     *Duration
	DurationUnit DurationUnit
	Deadline     *Deadline
	DeadlineLang *DeadlineLang
	Notes        []Note
	ChildTasks   []Task
}

func CreateTask(title string, desc string, prio Priority, indent Indent, durUnit *DurationUnit) Task {
	task := Task{
		Title:       Content(title),
		Description: Description(desc),
		Priority:    prio,
		Subtask:     indent,
	}
	if durUnit != nil {
		task.DurationUnit = DurationUnit(*durUnit)
	} else {
		task.DurationUnit = NoneDurationUnit
	}
	return task
}

func (t *Task) AddNote(note Note) {
	t.Notes = append(t.Notes, note)
}

func (t *Task) AddChildTask(task Task) {
	t.ChildTasks = append(t.ChildTasks, task)
}

type Note struct {
	Content Content
}
