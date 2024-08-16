package service_test

import (
	"testing"

	"code.ewintr.nl/planner/service"
)

func TestMemoryProjects(t *testing.T) {
	t.Parallel()

	mem := service.NewMemory()
	actProjects, actErr := mem.FindAllProjects()
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actProjects) != 0 {
		t.Errorf("exp 0, got %d", len(actProjects))
	}
}
