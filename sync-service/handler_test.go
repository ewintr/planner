package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"testing"
	"time"
)

func TestServerServeHTTP(t *testing.T) {
	t.Parallel()

	apiKey := "test"
	srv := NewServer(NewMemory(), apiKey, slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	for _, tc := range []struct {
		name      string
		key       string
		url       string
		method    string
		expStatus int
	}{
		{
			name:      "index always visible",
			url:       "/",
			method:    http.MethodGet,
			expStatus: http.StatusOK,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, tc.url, nil)
			if err != nil {
				t.Errorf("exp nil, got %v", err)
			}
			res := httptest.NewRecorder()
			srv.ServeHTTP(res, req)
			if res.Result().StatusCode != tc.expStatus {
				t.Errorf("exp %v, got %v", tc.expStatus, res.Result().StatusCode)
			}
		})
	}
}

func TestSyncGet(t *testing.T) {
	t.Parallel()

	now := time.Now()
	mem := NewMemory()

	items := []Syncable{
		{ID: "id-0", Updated: now.Add(-10 * time.Minute)},
		{ID: "id-1", Updated: now.Add(-5 * time.Minute)},
		{ID: "id-2", Updated: now.Add(time.Minute)},
	}

	for _, item := range items {
		if err := mem.Update(item); err != nil {
			t.Errorf("exp nil, got %v", err)
		}
	}

	apiKey := "test"
	srv := NewServer(mem, apiKey, slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	for _, tc := range []struct {
		name      string
		ts        time.Time
		expStatus int
		expItems  []Syncable
	}{
		{
			name:      "full",
			expStatus: http.StatusOK,
			expItems:  items,
		},
		{
			name:      "normal",
			ts:        now.Add(-6 * time.Minute),
			expStatus: http.StatusOK,
			expItems:  []Syncable{items[1], items[2]},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("/sync?ts=%s", url.QueryEscape(tc.ts.Format(time.RFC3339)))
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				t.Errorf("exp nil, got %v", err)
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			res := httptest.NewRecorder()
			srv.ServeHTTP(res, req)

			if res.Result().StatusCode != tc.expStatus {
				t.Errorf("exp %v, got %v", tc.expStatus, res.Result().StatusCode)
			}
			var actItems []Syncable
			actBody, err := io.ReadAll(res.Result().Body)
			if err != nil {
				t.Errorf("exp nil, got %v", err)
			}
			defer res.Result().Body.Close()

			if err := json.Unmarshal(actBody, &actItems); err != nil {
				t.Errorf("exp nil, got %v", err)
			}

			if len(actItems) != len(tc.expItems) {
				t.Errorf("exp %d, got %d", len(tc.expItems), len(actItems))
			}
			sort.Slice(actItems, func(i, j int) bool {
				return actItems[i].ID < actItems[j].ID
			})
			for i := range actItems {
				if actItems[i].ID != tc.expItems[i].ID {
					t.Errorf("exp %v, got %v", tc.expItems[i].ID, actItems[i].ID)
				}
			}
		})
	}
}

func TestSyncPost(t *testing.T) {
	t.Parallel()

	apiKey := "test"
	for _, tc := range []struct {
		name      string
		reqBody   []byte
		expStatus int
		expItems  []Syncable
	}{
		{
			name:      "empty",
			expStatus: http.StatusBadRequest,
		},
		{
			name:      "invalid",
			reqBody:   []byte(`{"fail}`),
			expStatus: http.StatusBadRequest,
		},
		{
			name: "normal",
			reqBody: []byte(`[
  {"ID":"id-1","Updated":"2024-09-06T08:00:00Z","Deleted":false,"Item":""},
  {"ID":"id-2","Updated":"2024-09-06T08:12:00Z","Deleted":false,"Item":""}
]`),
			expStatus: http.StatusNoContent,
			expItems: []Syncable{
				{ID: "id-1", Updated: time.Date(2024, 9, 6, 8, 0, 0, 0, time.UTC)},
				{ID: "id-2", Updated: time.Date(2024, 9, 6, 12, 0, 0, 0, time.UTC)},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mem := NewMemory()
			srv := NewServer(mem, apiKey, slog.New(slog.NewJSONHandler(os.Stdout, nil)))
			req, err := http.NewRequest(http.MethodPost, "/sync", bytes.NewBuffer(tc.reqBody))
			if err != nil {
				t.Errorf("exp nil, got %v", err)
			}
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
			res := httptest.NewRecorder()
			srv.ServeHTTP(res, req)

			if res.Result().StatusCode != tc.expStatus {
				t.Errorf("exp %v, got %v", tc.expStatus, res.Result().StatusCode)
			}

			actItems, err := mem.Updated(time.Time{})
			if err != nil {
				t.Errorf("exp nil, git %v", err)
			}
			if len(actItems) != len(tc.expItems) {
				t.Errorf("exp %d, got %d", len(tc.expItems), len(actItems))
			}
			sort.Slice(actItems, func(i, j int) bool {
				return actItems[i].ID < actItems[j].ID
			})
			for i := range actItems {
				if actItems[i].ID != tc.expItems[i].ID {
					t.Errorf("exp %v, got %v", tc.expItems[i].ID, actItems[i].ID)
				}
			}
		})
	}
}
