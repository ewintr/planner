package main

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Syncer interface {
	Update(item Item) error
	Updated(t time.Time) ([]Item, error)
}
