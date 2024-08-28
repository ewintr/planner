package storage

import (
	"errors"
	"time"

	"code.ewintr.nl/planner/planner"
)

var (
	ErrNotFound = errors.New("not found")
)

type Syncer interface {
	Update(item planner.Syncable) error
	Updated(t time.Time) ([]planner.Syncable, error)
}
