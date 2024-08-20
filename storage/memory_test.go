package storage_test

import (
	"testing"

	"code.ewintr.nl/planner/service"
	"github.com/google/go-cmp/cmp"
)

func TestMemoryProjects(t *testing.T) {
	t.Parallel()

	mem := service.NewMemory()

	t.Log("start empty")
	actProjects, actErr := mem.FindAllProjects()
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actProjects) != 0 {
		t.Errorf("exp 0, got %d", len(actProjects))
	}

	t.Log("add one")
	p1 := service.Project{
		ID:   "p1",
		Name: "project 1",
	}
	p2 := service.Project{
		ID:   "p2",
		Name: "project 2",
	}
	if actErr := mem.StoreProject(p1); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actProjects, actErr = mem.FindAllProjects()
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	expProjects := []service.Project{p1}
	if diff := cmp.Diff(expProjects, actProjects); diff != "" {
		t.Errorf("(-exp, +got):\n%s", diff)
	}

	t.Log("add second")
	if actErr := mem.StoreProject(p2); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actProjects, actErr = mem.FindAllProjects()
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	expProjects = []service.Project{p1, p2}
	if diff := cmp.Diff(expProjects, actProjects); diff != "" {
		t.Errorf("(-exp, +act):\n%s", diff)
	}

	t.Log("remove first")
	if actErr := mem.RemoveProject(p1.ID); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actProjects, actErr = mem.FindAllProjects()
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	expProjects = []service.Project{p2}
	if diff := cmp.Diff(expProjects, actProjects); diff != "" {
		t.Errorf("-exp, +act:\b%s", diff)
	}
}
