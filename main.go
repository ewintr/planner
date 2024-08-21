package main

import (
	"net/http"

	"code.ewintr.nl/planner/handler"
	"code.ewintr.nl/planner/storage"
)

func main() {
	mem := storage.NewMemory()

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/sync", handler.NewSyncHandler(mem))

	http.ListenAndServer(":8092", nil)
}
