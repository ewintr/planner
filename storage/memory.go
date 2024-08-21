package storage

import "code.ewintr.nl/planner/planner"

type Memory struct {
	items map[string]planner.Syncable
}

func NewMemory() *Memory {
	return &Memory{
		items: make(map[string]planner.Syncable),
	}
}

func (m *Memory) StoreProject(item planner.Syncable) error {
	m.items[item.ID()] = item

	return nil
}

/*
func (m *Memory) RemoveProject(id string) error {
	if _, ok := m.items[id]; !ok {
		return ErrNotFound
	}
	delete(m.items, id)

	return nil
}

func (m *Memory) FindProject(id string) (Project, error) {
	project, ok := m.items[id]
	if !ok {
		return Project{}, ErrNotFound
	}
	return project, nil
}

func (m *Memory) FindAllProjects() ([]Project, error) {
	items := make([]Project, 0, len(m.items))
	for _, p := range m.items {
		items = append(items, p)
	}
	return items, nil
}
*/
