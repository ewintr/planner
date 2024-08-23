package planner

import (
	"time"

	"github.com/google/uuid"
)

type Syncable interface {
	ID() string
	Updated() time.Time
}

type Task struct {
	id          string
	description string
	updated     time.Time
}

func NewTask(description string) Task {
	return Task{
		id:          uuid.New().String(),
		description: description,
		updated:     time.Now(),
	}
}

func (t Task) ID() string {
	return t.id
}

func (t Task) Updated() time.Time {
	return t.updated
}
