package main

import (
	"time"
)

type Memory struct {
	items map[string]Item
}

func NewMemory() *Memory {
	return &Memory{
		items: make(map[string]Item),
	}
}

func (m *Memory) Update(item Item) error {
	m.items[item.ID] = item

	return nil
}

func (m *Memory) Updated(timestamp time.Time) ([]Item, error) {
	result := make([]Item, 0)

	for _, i := range m.items {
		if timestamp.IsZero() || i.Updated.Equal(timestamp) || i.Updated.After(timestamp) {
			result = append(result, i)
		}
	}

	return result, nil
}
