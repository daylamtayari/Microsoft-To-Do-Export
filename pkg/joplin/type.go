package joplin

import (
	"time"

	"github.com/google/uuid"
)

// Joplin note type
type Note struct {
	Id            uuid.UUID
	ParentId      *uuid.UUID
	Type          int
	Title         string
	Body          string
	IsTodo        bool
	TodoDue       *time.Time
	TodoCompleted *time.Time
	NoteId        *uuid.UUID
	TagId         *uuid.UUID
	CreatedTime   time.Time
}
