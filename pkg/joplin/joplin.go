package joplin

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const timeFormat = "2006-01-02T15:04:05.000000-0700"

// Outputs a Joplin UUID (no hyphens) from a provided UUID
func OutputId(uuid uuid.UUID) string {
	return strings.ReplaceAll(uuid.String(), "-", "")
}

// Output a note contents in the Joplin format
func OutputNote(note Note) string {
	metadata := ""
	metadata += "\nid: " + OutputId(note.Id)
	metadata += "\nparent_id: "
	if note.ParentId != nil {
		metadata += OutputId(*note.ParentId)
	}
	if note.IsTodo {
		metadata += "\nis_todo: 1"
		if note.TodoDue != nil {
			metadata += "\ntodo_due: " + strconv.FormatInt(note.TodoDue.UnixMilli(), 10)
		}
		if note.TodoCompleted != nil {
			metadata += "\ntodo_completed: " + strconv.FormatInt(note.TodoCompleted.UnixMilli(), 10)
		}
	}
	if note.TagId != nil {
		metadata += "\ntag_id: " + OutputId(*note.TagId)
	}
	if note.NoteId != nil {
		metadata += "\nnote_id: " + OutputId(*note.NoteId)
	}
	createdTimeStr := note.CreatedTime.Format(timeFormat)
	metadata += "\ncreated_time: " + createdTimeStr + "\nuser_created_time: " + createdTimeStr
	metadata += "\ntype_: " + strconv.Itoa(note.Type)

	return note.Title + "\n\n" + note.Body + "\n" + metadata
}

// Creates a folder note
func CreateFolder(title string, createdAt *time.Time, parent *uuid.UUID) Note {
	if createdAt == nil {
		currentTime := time.Now()
		createdAt = &currentTime
	}
	return Note{
		Id:          uuid.New(),
		Type:        2,
		Title:       title,
		CreatedTime: *createdAt,
		ParentId:    parent,
	}
}

// Creates a generic note
func CreateNote(title, body string, parent *uuid.UUID, createdAt *time.Time) Note {
	if createdAt == nil {
		currentTime := time.Now()
		createdAt = &currentTime
	}
	return Note{
		Id:          uuid.New(),
		Type:        1,
		ParentId:    parent,
		Title:       title,
		Body:        body,
		CreatedTime: *createdAt,
	}
}

// Creates a to do note
func CreateToDo(title, body string, parent *uuid.UUID, due *time.Time, completed *time.Time, createdAt *time.Time) Note {
	if createdAt == nil {
		currentTime := time.Now()
		createdAt = &currentTime
	}
	return Note{
		Id:            uuid.New(),
		Type:          1,
		ParentId:      parent,
		IsTodo:        true,
		Title:         title,
		Body:          body,
		TodoDue:       due,
		TodoCompleted: completed,
		CreatedTime:   *createdAt,
	}
}

// Creates a tag note
func CreateTag(name string, createdAt *time.Time) Note {
	if createdAt == nil {
		currentTime := time.Now()
		createdAt = &currentTime
	}
	return Note{
		Id:          uuid.New(),
		Type:        5,
		Title:       name,
		CreatedTime: *createdAt,
	}
}

// Create a tag <-> note reference
func CreateNoteTag(tagId uuid.UUID, noteId uuid.UUID) Note {
	id := uuid.New()
	return Note{
		Id:          id,
		Title:       "",
		Type:        6,
		CreatedTime: time.Now(),
		TagId:       &tagId,
		NoteId:      &noteId,
	}
}
