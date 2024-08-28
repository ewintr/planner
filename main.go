package main

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"code.ewintr.nl/planner/handler"
	"code.ewintr.nl/planner/storage"
)

func main() {
	mem := storage.NewMemory()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	go http.ListenAndServe(":8092", handler.NewServer(mem, logger))

	logger.Info("service started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Info("service stopped")
}
