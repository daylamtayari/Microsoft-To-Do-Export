package todoistcsv

import (
	"strconv"
	"time"
)

// Types are accurate representation of the Todoist CSV import format:
// https://www.todoist.com/help/articles/import-or-export-a-project-as-a-csv-file-in-todoist-YC8YvN

type Type string

const (
	TaskType    = Type("task")
	SectionType = Type("section")
	NoteType    = Type("note")
)

type Content string

type Description string

type Priority int

const (
	Priority1 = iota + 1
	Priority2
	Priority3
	Priority4
)

type Indent int

const (
	Indent1 = iota + 1
	Indent2
	Indent3
	Indent4
)

type User struct {
	Username string
	UserID   int
}

func (u User) String() string {
	return u.Username + " (" + strconv.Itoa(u.UserID) + ")"
}

type Author User

type Responsible User

type Date time.Time

type LanguageCode string

const (
	CzechLang              = LanguageCode("cs")
	DanishLang             = LanguageCode("da")
	GermanLang             = LanguageCode("de")
	EnglishLang            = LanguageCode("en")
	SpanishLang            = LanguageCode("es")
	FinnishLang            = LanguageCode("fi")
	FrenchLang             = LanguageCode("fr")
	ItalianLang            = LanguageCode("it")
	JapaneseLang           = LanguageCode("ja")
	KoreanLang             = LanguageCode("ko")
	NorwegianLang          = LanguageCode("nb")
	DutchLang              = LanguageCode("nl")
	PolishLang             = LanguageCode("pl")
	BrazilianPortugeseLang = LanguageCode("pt_BR")
	RussianLang            = LanguageCode("ru")
	SwedishLang            = LanguageCode("sv")
	SimplifiedChineseLang  = LanguageCode("zh_CN")
	TraditionalChineseLang = LanguageCode("zh_TW")
)

type DateLang LanguageCode

type Timezone time.Location

type Duration int

type DurationUnit string

const (
	MinuteDurationUnit = DurationUnit("minute")
	NoneDurationUnit   = DurationUnit("None")
)

type Deadline time.Time

type DeadlineLang LanguageCode

type Meta string

const (
	ListLayout  = Meta("view_style=list")
	BoardLayout = Meta("view_style=board")
)
