package storage

import (
	"time"

	"code.ewintr.nl/planner/planner"
)

type deletedItem struct {
	ID        string
	Timestamp time.Time
}

type Memory struct {
	items   map[string]planner.Syncable
	deleted []deletedItem
}

func NewMemory() *Memory {
	return &Memory{
		items:   make(map[string]planner.Syncable),
		deleted: make([]deletedItem, 0),
	}
}

func (m *Memory) Update(item planner.Syncable) error {
	m.items[item.ID()] = item

	return nil
}

func (m *Memory) Updated(timestamp time.Time) ([]planner.Syncable, error) {
	result := make([]planner.Syncable, 0)

	for _, i := range m.items {
		if timestamp.IsZero() || i.Updated().Equal(timestamp) || i.Updated().After(timestamp) {
			result = append(result, i)
		}
	}

	return result, nil
}

func (m *Memory) Delete(id string) error {
	if _, exists := m.items[id]; !exists {
		return ErrNotFound
	}

	delete(m.items, id)
	m.deleted = append(m.deleted, deletedItem{
		ID:        id,
		Timestamp: time.Now(),
	})

	return nil
}

func (m *Memory) Deleted(t time.Time) ([]string, error) {
	result := make([]string, 0)
	for _, di := range m.deleted {
		if di.Timestamp.Equal(t) || di.Timestamp.After(t) {
			result = append(result, di.ID)
		}
	}
	return result, nil
}
