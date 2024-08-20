package storage

import "code.ewintr.nl/planner/planner"

type Memory struct {
	projects map[string]planner.Task
}

func NewMemory() *Memory {
	return &Memory{
		projects: make(map[string]planner.Task),
	}
}

func (m *Memory) StoreProject(project planner.Task) error {
	m.projects[project.ID] = project

	return nil
}

func (m *Memory) RemoveProject(id string) error {
	if _, ok := m.projects[id]; !ok {
		return ErrNotFound
	}
	delete(m.projects, id)

	return nil
}

func (m *Memory) FindProject(id string) (Project, error) {
	project, ok := m.projects[id]
	if !ok {
		return Project{}, ErrNotFound
	}
	return project, nil
}

func (m *Memory) FindAllProjects() ([]Project, error) {
	projects := make([]Project, 0, len(m.projects))
	for _, p := range m.projects {
		projects = append(projects, p)
	}
	return projects, nil
}
