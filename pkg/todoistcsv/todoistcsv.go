// Complete handling of the Todoist CSV import format,
// allowing for ease of creating valid import CSV files.
package todoistcsv

import (
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

// Returns the contents of a Todoist CSV import file
func (e *Export) CSV() string {
	var csvBuilder strings.Builder

	// Write header row and meta field
	csvBuilder.WriteString("TYPE,CONTENT,DESCRIPTION,PRIORITY,INDENT,AUTHOR,RESPONSIBLE,DATE,DATE_LANG,TIMEZONE,DURATION,DURATION_UNIT,DEADLINE,DEADLINE_LANG")
	csvBuilder.WriteString("\nmeta," + string(e.Meta) + ",,,,,,,,,,,,")
	csvBuilder.WriteString("\n,,,,,,,,,,,,,")

	// Write all sections and tasks
	for i := range e.Sections {
		csvBuilder.WriteString("\n" + string(SectionType) + "," + e.Sections[i].Title + ",,,,,,,,,,,,")
		for j := range e.Sections[i].Tasks {
			exportTask(&csvBuilder, e.Sections[i].Tasks[j])
		}
	}

	return csvBuilder.String()
}

func exportTask(builder *strings.Builder, task Task) {
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
	builder.WriteString("\n" + string(TaskType) + "," + string(task.Title) + "," + string(task.Description) + "," + strconv.Itoa(int(task.Priority)) + "," + strconv.Itoa(int(task.Subtask)) + "," + author + "," + responsible + "," + date + "," + dateLang + "," + timezone + "," + duration + "," + string(task.DurationUnit) + "," + deadline + "," + deadlineLang)

	// Write all notes
	for i := range task.Notes {
		builder.WriteString("\n" + string(NoteType) + "," + string(task.Notes[i].Content) + ",,,,,,,,,,,,")
	}

	// Recursively call for any child tasks
	for i := range task.ChildTasks {
		exportTask(builder, task.ChildTasks[i])
	}
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
