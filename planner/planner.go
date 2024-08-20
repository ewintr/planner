package planner

import (
	"time"

	"github.com/google/uuid"
)

type Syncable interface {
	LastUpdated() time.Time
}

type Task struct {
	ID          string
	description string
	updated     time.Time
}

func NewTask(description string) Task {
	return Task{
		ID:          uuid.New(),
		description: description,
		updated:     time.Now(),
	}
}
