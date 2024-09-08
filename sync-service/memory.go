package main

import (
	"time"
)

type Memory struct {
	items map[string]Syncable
}

func NewMemory() *Memory {
	return &Memory{
		items: make(map[string]Syncable),
	}
}

func (m *Memory) Update(item Syncable) error {
	m.items[item.ID] = item

	return nil
}

func (m *Memory) Updated(timestamp time.Time) ([]Syncable, error) {
	result := make([]Syncable, 0)

	for _, i := range m.items {
		if timestamp.IsZero() || i.Updated.Equal(timestamp) || i.Updated.After(timestamp) {
			result = append(result, i)
		}
	}

	return result, nil
}
