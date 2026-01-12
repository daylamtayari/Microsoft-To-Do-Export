package todoistcsv

// Completed handling of the Todoist CSV import format,
// allowing for ease of creating valid import CSV files.

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

type Section struct {
	Tasks []Task
}

func CreateSection(tasks []Task) *Section {
	return &Section{
		Tasks: tasks,
	}
}

func (s *Section) AddTask(task Task) {
	s.Tasks = append(s.Tasks, task)
}

type Task struct {
	Title        Content
	Description  Description
	Priority     *Priority
	Subtask      *Indent
	Author       *Author
	Responsible  *Responsible
	Date         *Date
	DateLang     *DateLang
	Timezone     *Timezone
	Duration     *Duration
	DurationUnit *DurationUnit
	Deadline     *Deadline
	DeadlineLang *DeadlineLang
	Notes        []Note
	ChildTasks   []Task
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
