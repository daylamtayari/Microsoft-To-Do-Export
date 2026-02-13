package superproductivity

// CompleteBackup represents the top-level structure of a Super Productivity backup file
type CompleteBackup struct {
	Timestamp         int64           `json:"timestamp"`
	LastUpdate        int64           `json:"lastUpdate"`
	CrossModelVersion float64         `json:"crossModelVersion"`
	Data              AppDataComplete `json:"data"`
}

// AppDataComplete contains all application data for import/export
type AppDataComplete struct {
	// Core data types (user will populate these)
	Task    TaskState    `json:"task"`
	Project ProjectState `json:"project"`
	Tag     TagState     `json:"tag"`

	// Required but minimal structure (empty for basic import)
	// These are typed as 'any' and constructed as maps in NewMinimalAppDataComplete
	GlobalConfig   any                `json:"globalConfig"`
	Note           any                `json:"note"`
	SimpleCounter  any                `json:"simpleCounter"`
	TaskRepeatCfg  TaskRepeatCfgState `json:"taskRepeatCfg"`
	Metric         any                `json:"metric"`
	Planner        any `json:"planner"`
	IssueProvider  any `json:"issueProvider"`
	Boards         any `json:"boards"`
	MenuTree       any `json:"menuTree"`
	TimeTracking   any `json:"timeTracking"`
	Reminders      any `json:"reminders"`
	PluginMetadata any `json:"pluginMetadata"`
	PluginUserData any `json:"pluginUserData"`
}

// EntityState is a generic structure for NgRx entity states
type EntityState[T any] struct {
	IDs      []string     `json:"ids"`
	Entities map[string]T `json:"entities"`
}

// TaskState extends EntityState with additional task-specific state fields
type TaskState struct {
	IDs                   []string        `json:"ids"`
	Entities              map[string]Task `json:"entities"`
	CurrentTaskID         *string         `json:"currentTaskId"`
	SelectedTaskID        *string         `json:"selectedTaskId"`
	LastCurrentTaskID     *string         `json:"lastCurrentTaskId,omitempty"`
	IsDataLoaded          bool            `json:"isDataLoaded,omitempty"`
	TaskDetailTargetPanel *string         `json:"taskDetailTargetPanel,omitempty"`
}

// Task represents a task or subtask
type Task struct {
	ID             string           `json:"id"`
	Title          string           `json:"title"`
	SubTaskIDs     []string         `json:"subTaskIds"`
	TimeSpentOnDay map[string]int64 `json:"timeSpentOnDay"`
	TimeSpent      int64            `json:"timeSpent"`
	TimeEstimate   int64            `json:"timeEstimate"`
	IsDone         bool             `json:"isDone"`
	TagIDs         []string         `json:"tagIds"`
	Created        int64            `json:"created"`
	ProjectID      string           `json:"projectId"`
	Attachments    []TaskAttachment `json:"attachments"`
	Notes          *string          `json:"notes,omitempty"`
	ParentID       *string          `json:"parentId,omitempty"`
	DueDay         *string          `json:"dueDay,omitempty"`
	DueWithTime    *int64           `json:"dueWithTime,omitempty"`
	HasPlannedTime *bool            `json:"hasPlannedTime,omitempty"`
	DoneOn         *int64           `json:"doneOn,omitempty"`
	Modified       *int64           `json:"modified,omitempty"`
	RemindAt       *int64           `json:"remindAt,omitempty"`
	ReminderID     *string          `json:"reminderId,omitempty"`
	RepeatCfgID    *string          `json:"repeatCfgId,omitempty"`
}

// TaskAttachment represents a file or link attachment on a task
type TaskAttachment struct {
	ID        string  `json:"id"`
	TaskID    string  `json:"taskId"`
	Type      string  `json:"type"`
	Path      *string `json:"path,omitempty"`
	Title     *string `json:"title,omitempty"`
	Icon      *string `json:"icon,omitempty"`
	CreatedAt *int64  `json:"createdAt,omitempty"`
}

// ProjectState is the entity state for projects
type ProjectState = EntityState[Project]

// Project represents a work context project
type Project struct {
	ID             string                 `json:"id"`
	Title          string                 `json:"title"`
	TaskIDs        []string               `json:"taskIds"`
	BacklogTaskIDs []string               `json:"backlogTaskIds"`
	NoteIDs        []string               `json:"noteIds"`
	Theme          WorkContextThemeCfg    `json:"theme"`
	AdvancedCfg    WorkContextAdvancedCfg `json:"advancedCfg"`
	Icon           *string                `json:"icon,omitempty"`
}

// TagState is the entity state for tags
type TagState = EntityState[Tag]

// Tag represents a work context tag
type Tag struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	TaskIDs     []string               `json:"taskIds"`
	Theme       WorkContextThemeCfg    `json:"theme"`
	AdvancedCfg WorkContextAdvancedCfg `json:"advancedCfg"`
	Icon        *string                `json:"icon,omitempty"`
}

// WorkContextThemeCfg defines theme configuration for projects and tags
type WorkContextThemeCfg struct {
	Primary                  string  `json:"primary"`
	Accent                   string  `json:"accent"`
	Warn                     string  `json:"warn"`
	IsAutoContrast           *bool   `json:"isAutoContrast,omitempty"`
	HuePrimary               *string `json:"huePrimary,omitempty"`
	HueAccent                *string `json:"hueAccent,omitempty"`
	HueWarn                  *string `json:"hueWarn,omitempty"`
	BackgroundImageDark      *string `json:"backgroundImageDark,omitempty"`
	BackgroundImageLight     *string `json:"backgroundImageLight,omitempty"`
	BackgroundOverlayOpacity *int    `json:"backgroundOverlayOpacity,omitempty"`
}

// WorkContextAdvancedCfg contains advanced configuration for work contexts
type WorkContextAdvancedCfg struct {
	WorklogExportSettings WorklogExportSettings `json:"worklogExportSettings"`
}

// WorklogExportSettings defines settings for exporting work logs
type WorklogExportSettings struct {
	Cols            []string `json:"cols"`
	GroupBy         string   `json:"groupBy"`
	SeparateTasksBy string   `json:"separateTasksBy"`
}

// TaskRepeatCfgState is the entity state for task repeat configurations
type TaskRepeatCfgState = EntityState[TaskRepeatCfg]

// TaskRepeatCfg represents a recurring task configuration
type TaskRepeatCfg struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	ProjectID       *string  `json:"projectId,omitempty"`
	TagIDs          []string `json:"tagIds"`
	RepeatCycle     string   `json:"repeatCycle"` // DAILY, WEEKLY, MONTHLY, YEARLY
	RepeatEvery     int32    `json:"repeatEvery"`
	StartDate       string   `json:"startDate"`       // YYYY-MM-DD format
	EndDate         *string  `json:"endDate,omitempty"` // YYYY-MM-DD format
	IsPaused        bool     `json:"isPaused"`
	Monday          bool     `json:"monday"`
	Tuesday         bool     `json:"tuesday"`
	Wednesday       bool     `json:"wednesday"`
	Thursday        bool     `json:"thursday"`
	Friday          bool     `json:"friday"`
	Saturday        bool     `json:"saturday"`
	Sunday          bool     `json:"sunday"`
	Order           int      `json:"order"`
	DefaultEstimate int64    `json:"defaultEstimate,omitempty"`
	QuickSetting    *string  `json:"quickSetting,omitempty"` // DAILY, WEEKLY_CURRENT_WEEKDAY, etc.
	StartTime       *string  `json:"startTime,omitempty"`
	RemindAt        *string  `json:"remindAt,omitempty"` // AtStart, etc.
	Notes           *string  `json:"notes,omitempty"`
	LastTaskCreation *int64   `json:"lastTaskCreation,omitempty"`
}

// NewEmptyEntityState creates an empty entity state
func NewEmptyEntityState[T any]() EntityState[T] {
	return EntityState[T]{
		IDs:      []string{},
		Entities: make(map[string]T),
	}
}

// createDefaultTheme creates a default theme configuration
func createDefaultTheme() WorkContextThemeCfg {
	return WorkContextThemeCfg{
		Primary: "#6495ED", // DEFAULT_TODAY_TAG_COLOR (for TODAY tag)
		Accent:  "#ff4081",
		Warn:    "#e11826",
	}
}

// createDefaultAdvancedCfg creates a default advanced configuration
func createDefaultAdvancedCfg() WorkContextAdvancedCfg {
	return WorkContextAdvancedCfg{
		WorklogExportSettings: WorklogExportSettings{
			Cols:            []string{"DATE", "START", "END", "TIME_CLOCK", "TITLES_INCLUDING_SUB"},
			GroupBy:         "DATE",
			SeparateTasksBy: " | ",
		},
	}
}

// NewMinimalAppDataComplete creates a minimal valid AppDataComplete structure
// Based on Super Productivity's test suite createMinimalValidBackup function
func NewMinimalAppDataComplete() AppDataComplete {
	// Create required INBOX_PROJECT with minimal fields
	inboxProject := Project{
		ID:             "INBOX_PROJECT",
		Title:          "Inbox",
		TaskIDs:        []string{},
		BacklogTaskIDs: []string{},
		NoteIDs:        []string{},
		Theme:          createDefaultTheme(),
		AdvancedCfg:    createDefaultAdvancedCfg(),
	}

	// Create required TODAY tag with minimal fields
	todayTag := Tag{
		ID:          "TODAY",
		Title:       "Today",
		TaskIDs:     []string{},
		Theme:       createDefaultTheme(),
		AdvancedCfg: createDefaultAdvancedCfg(),
	}

	return AppDataComplete{
		Task: TaskState{
			IDs:            []string{},
			Entities:       make(map[string]Task),
			CurrentTaskID:  nil,
			SelectedTaskID: nil,
		},
		Project: EntityState[Project]{
			IDs:      []string{"INBOX_PROJECT"},
			Entities: map[string]Project{"INBOX_PROJECT": inboxProject},
		},
		Tag: EntityState[Tag]{
			IDs:      []string{"TODAY"},
			Entities: map[string]Tag{"TODAY": todayTag},
		},
		// All minimal fields constructed as maps matching expected JSON structure
		GlobalConfig: map[string]any{
			"misc": map[string]any{"isDisableInitialDialog": true},
			"sync": map[string]any{"isEnabled": false, "syncProvider": nil},
		},
		Note: map[string]any{
			"ids":        []string{},
			"entities":   map[string]any{},
			"todayOrder": []string{},
		},
		SimpleCounter: map[string]any{
			"ids":      []string{},
			"entities": map[string]any{},
		},
		TaskRepeatCfg: NewEmptyEntityState[TaskRepeatCfg](),
		Metric: map[string]any{
			"ids":      []string{},
			"entities": map[string]any{},
		},
		IssueProvider: map[string]any{
			"ids":      []string{},
			"entities": map[string]any{},
		},
		Planner: map[string]any{
			"days": map[string]any{},
		},
		Boards: map[string]any{
			"boardCfgs": []any{},
		},
		MenuTree: map[string]any{
			"tagTree":     []string{},
			"projectTree": []string{},
		},
		TimeTracking: map[string]any{
			"project": map[string]any{},
			"tag":     map[string]any{},
		},
		Reminders:      []any{},
		PluginMetadata: []any{},
		PluginUserData: []any{},
	}
}
