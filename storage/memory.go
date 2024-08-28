package storage

import (
	"time"

	"code.ewintr.nl/planner/planner"
)

type Memory struct {
	items map[string]planner.Syncable
}

func NewMemory() *Memory {
	return &Memory{
		items: make(map[string]planner.Syncable),
	}
}

func (m *Memory) Update(item planner.Syncable) error {
	m.items[item.ID] = item

	return nil
}

func (m *Memory) Updated(timestamp time.Time) ([]planner.Syncable, error) {
	result := make([]planner.Syncable, 0)

	for _, i := range m.items {
		if timestamp.IsZero() || i.Updated.Equal(timestamp) || i.Updated.After(timestamp) {
			result = append(result, i)
		}
	}

	return result, nil
}
