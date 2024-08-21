package storage

import (
	"errors"
	"time"

	"code.ewintr.nl/planner/planner"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	NewSince(t time.Time) ([]planner.Syncable, error)
	Store(item planner.Syncable) error
	// FindTask(id string) (planner.Task, error)
	// FindAllTasks() ([]planner.Task, error)
	// StoreTask(project planner.Task) error
}
