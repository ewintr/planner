package main

import (
	"time"

	"github.com/google/uuid"
)

type Kind string

const (
	KindTask Kind = "task"
)

type Syncable struct {
	ID      string
	Kind    Kind
	Updated time.Time
	Deleted bool
	Item    string
}

func NewSyncable(item string) Syncable {
	return Syncable{
		ID:      uuid.New().String(),
		Updated: time.Now(),
		Item:    item,
	}
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
