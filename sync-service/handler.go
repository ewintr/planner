package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"path"
	"strings"
	"time"
)

type Server struct {
	syncer Syncer
	apiKey string
	logger *slog.Logger
}

func NewServer(syncer Syncer, apiKey string, logger *slog.Logger) *Server {
	return &Server{
		syncer: syncer,
		apiKey: apiKey,
		logger: logger,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/" {
		Index(w, r)
		return
	}

	if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", s.apiKey) {
		http.Error(w, `{"error":"not authorized"}`, http.StatusUnauthorized)
		return
	}

	head, tail := ShiftPath(r.URL.Path)
	switch {
	case head == "sync" && tail != "/":
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
	case head == "sync" && r.Method == http.MethodGet:
		s.SyncGet(w, r)
	case head == "sync" && r.Method == http.MethodPost:
		s.SyncPost(w, r)
	default:
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)
	}
}

func (s *Server) SyncGet(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Time{}
	tsStr := r.URL.Query().Get("ts")
	if tsStr != "" {
		var err error
		if timestamp, err = time.Parse(time.RFC3339, tsStr); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	items, err := s.syncer.Updated(timestamp)
	if err != nil {
		http.Error(w, fmtError(err), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(items)
	if err != nil {
		http.Error(w, fmtError(err), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(body))
}

func (s *Server) SyncPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmtError(err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var items []Syncable
	if err := json.Unmarshal(body, &items); err != nil {
		http.Error(w, fmtError(err), http.StatusBadRequest)
		return
	}

	for _, item := range items {
		item.Updated = time.Now()
		if err := s.syncer.Update(item); err != nil {
			http.Error(w, fmtError(err), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)

}

// ShiftPath splits off the first component of p, which will be cleaned of
// relative components before processing. head will never contain a slash and
// tail will always be a rooted path without trailing slash.
// See https://blog.merovius.de/posts/2017-06-18-how-not-to-use-an-http-router/
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"status":"ok"}`)
}

func fmtError(err error) string {
	return fmt.Sprintf(`{"error":%q}`, err.Error())
}
