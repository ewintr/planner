package main

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Syncer interface {
	Update(item Syncable) error
	Updated(t time.Time) ([]Syncable, error)
}
