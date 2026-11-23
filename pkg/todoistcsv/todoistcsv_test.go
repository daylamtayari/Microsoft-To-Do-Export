package todoistcsv

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestCSVEscapesSpecialCharacters(t *testing.T) {
	parent := CreateTask(`Pay bills, then "report"`, "Chase paid, do cash report\r\nbefore paying Discover\n", Priority1, Indent1, nil)
	parent.AddChildTask(CreateTask("Step, one", "", Priority1, Indent2, nil))
	parent.AddNote(Note{Content: "note, with\nnewline"})

	section := CreateSection("Home, Stuff", nil)
	section.AddTask(parent)
	section.AddTask(CreateTask("Plain task", "", Priority1, Indent1, nil))

	export := CreateExport(ListLayout, []Section{*section})

	records, err := csv.NewReader(strings.NewReader(export.CSV())).ReadAll()
	if err != nil {
		t.Fatalf("output is not valid CSV: %v", err)
	}

	for i, r := range records {
		if len(r) != columnCount {
			t.Errorf("record %d has %d fields, want %d: %q", i, len(r), columnCount, r)
		}
	}

	want := [][2]string{
		{"TYPE", "CONTENT"},
		{"meta", string(ListLayout)},
		{"", ""},
		{"section", "Home, Stuff"},
		{"task", `Pay bills, then "report"`},
		{"note", "note, with\nnewline"},
		{"task", "Step, one"},
		{"task", "Plain task"},
	}
	if len(records) != len(want) {
		t.Fatalf("got %d records, want %d", len(records), len(want))
	}
	for i, w := range want {
		if records[i][0] != w[0] || records[i][1] != w[1] {
			t.Errorf("record %d = %q, %q; want %q, %q", i, records[i][0], records[i][1], w[0], w[1])
		}
	}

	if got, want := records[4][2], "Chase paid, do cash report\nbefore paying Discover\n"; got != want {
		t.Errorf("description = %q, want %q", got, want)
	}
	if records[6][4] != "2" {
		t.Errorf("child task indent = %q, want 2", records[6][4])
	}
}
