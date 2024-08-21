package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.ewintr.nl/planner/storage"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status":"ok"}`)
}

func NewSyncHandler(mem storage.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		items, err := mem.NewSince(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(items)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, body)
	}

}
