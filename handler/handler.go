package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"code.ewintr.nl/planner/planner"
	"code.ewintr.nl/planner/storage"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status":"ok"}`)
}

type ChangeSummary struct {
	Updated []planner.Syncable
	Deleted []string
}

func NewSyncHandler(mem storage.Syncer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			timestamp := time.Time{}
			tsStr := r.URL.Query().Get("ts")
			if tsStr != "" {
				var err error
				if timestamp, err = time.Parse(time.RFC3339, tsStr); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			items, err := mem.Updated(timestamp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			deleted, err := mem.Deleted(timestamp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			result := ChangeSummary{
				Updated: items,
				Deleted: deleted,
			}

			body, err := json.Marshal(result)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprint(w, string(body))

		case http.MethodPost:
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			var changes ChangeSummary
			if err := json.Unmarshal(body, changes); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			for _, updated := range changes.Updated {
				if err := mem.Update(updated); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			for _, deleted := range changes.Deleted {
				if err := mem.Delete(deleted); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
