package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"code.ewintr.nl/planner/storage"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status":"ok"}`)
}

func NewSyncHandler(mem storage.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		timestamp := time.Time{}
		tsStr := r.URL.Query().Get("ts")
		if tsStr != "" {
			var err error
			if timestamp, err = time.Parse(time.RFC3339, tsStr); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		items, err := mem.NewSince(timestamp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(body))
	}

}
