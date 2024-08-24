package storage_test

import (
	"errors"
	"testing"
	"time"

	"code.ewintr.nl/planner/planner"
	"code.ewintr.nl/planner/storage"
)

func TestMemoryItem(t *testing.T) {
	t.Parallel()

	mem := storage.NewMemory()

	t.Log("start empty")
	actItems, actErr := mem.Updated(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 0 {
		t.Errorf("exp 0, got %d", len(actItems))
	}

	t.Log("add one")
	t1 := planner.NewTask("test")
	if actErr := mem.Update(t1); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actItems, actErr = mem.Updated(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 1 {
		t.Errorf("exp 1, gor %d", len(actItems))
	}
	if actItems[0].ID() != t1.ID() {
		t.Errorf("exp %v, got %v", actItems[0].ID(), t1.ID())
	}

	before := time.Now()

	t.Log("add second")
	t2 := planner.NewTask("test 2")
	if actErr := mem.Update(t2); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actItems, actErr = mem.Updated(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 2 {
		t.Errorf("exp 2, gor %d", len(actItems))
	}
	if actItems[0].ID() != t1.ID() {
		t.Errorf("exp %v, got %v", actItems[0].ID(), t1.ID())
	}
	if actItems[1].ID() != t2.ID() {
		t.Errorf("exp %v, got %v", actItems[1].ID(), t2.ID())
	}
	actDeleted, actErr := mem.Deleted(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actDeleted) != 0 {
		t.Errorf("exp 0, got %d", len(actDeleted))
	}

	actItems, actErr = mem.Updated(before)
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 1 {
		t.Errorf("exp 1, gor %d", len(actItems))
	}
	if actItems[0].ID() != t2.ID() {
		t.Errorf("exp %v, got %v", actItems[0].ID(), t2.ID())
	}

	t.Log("remove first")
	if actErr := mem.Delete(t1.ID()); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actItems, actErr = mem.Updated(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 1 {
		t.Errorf("exp 2, gor %d", len(actItems))
	}
	if actItems[0].ID() != t2.ID() {
		t.Errorf("exp %v, got %v", actItems[0].ID(), t1.ID())
	}
	actDeleted, actErr = mem.Deleted(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actDeleted) != 1 {
		t.Errorf("exp 1, got %d", len(actDeleted))
	}
	if actDeleted[0] != t1.ID() {
		t.Errorf("exp %v, got %v", actDeleted[0], t1.ID())
	}
	actDeleted, actErr = mem.Deleted(time.Now())
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actDeleted) != 0 {
		t.Errorf("exp 0, got %d", len(actDeleted))
	}

	t.Log("remove non-existing")
	if actErr := mem.Delete("test"); !errors.Is(actErr, storage.ErrNotFound) {
		t.Errorf("exp %v, got %v", storage.ErrNotFound, actErr)
	}
}
