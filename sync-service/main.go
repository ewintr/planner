package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	dbPath := os.Getenv("PLANNER_DB_PATH")
	if dbPath == "" {
		fmt.Println("PLANNER_DB_PATH is empty")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Getenv("PLANNER_PORT"))
	if err != nil {
		fmt.Println("PLANNER_PORT env is not an integer")
		os.Exit(1)
	}
	apiKey := os.Getenv("PLANNER_API_KEY")
	if apiKey == "" {
		fmt.Println("PLANNER_API_KEY is empty")
		os.Exit(1)
	}
	crtPath := os.Getenv("PLANNER_CRT_PATH")
	keyPath := os.Getenv("PLANNER_KEY_PATH")

	repo, err := NewSqlite(dbPath)
	if err != nil {
		fmt.Printf("could not open sqlite db: %s", err.Error())
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	address := fmt.Sprintf(":%d", port)
	srv := NewServer(repo, apiKey, logger)
	if crtPath != "" || keyPath != "" {
		go http.ListenAndServeTLS(address, crtPath, keyPath, srv)
	} else {
		go http.ListenAndServe(address, srv)
	}

	logger.Info("service started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Info("service stopped")
}
