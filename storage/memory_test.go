package storage_test

import (
	"testing"
	"time"

	"code.ewintr.nl/planner/planner"
	"code.ewintr.nl/planner/storage"
)

func TestMemoryItem(t *testing.T) {
	t.Parallel()

	mem := storage.NewMemory()

	t.Log("start empty")
	actItems, actErr := mem.NewSince(time.Time{})
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 0 {
		t.Errorf("exp 0, got %d", len(actItems))
	}

	t.Log("add one")
	t1 := planner.NewTask("test")
	if actErr := mem.Store(t1); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actItems, actErr = mem.NewSince(time.Time{})
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
	if actErr := mem.Store(t2); actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	actItems, actErr = mem.NewSince(time.Time{})
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

	actItems, actErr = mem.NewSince(before)
	if actErr != nil {
		t.Errorf("exp nil, got %v", actErr)
	}
	if len(actItems) != 1 {
		t.Errorf("exp 1, gor %d", len(actItems))
	}
	if actItems[0].ID() != t2.ID() {
		t.Errorf("exp %v, got %v", actItems[0].ID(), t2.ID())
	}

	/*
			t.Log("remove first")
			if actErr := mem.RemoveProject(p1.ID); actErr != nil {
				t.Errorf("exp nil, got %v", actErr)
			}
			actItems , actErr = mem.FindAllItems ()
			if actErr != nil {
				t.Errorf("exp nil, got %v", actErr)
			}
			expItems = []service.Project{p2}
			if diff := cmp.Diff(expItems
		  , actItems ); diff != "" {
				t.Errorf("-exp, +act:\b%s", diff)
			}
	*/
}
