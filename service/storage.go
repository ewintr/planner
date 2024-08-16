package service

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	FindProject(id string) (Project, error)
	FindAllProjects() ([]Project, error)
	StoreProject(project Project) error
}
