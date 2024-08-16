package main

import (
	"fmt"
	"net/http"
	"time"
)

type Todoist struct {
	apiKey string
	client http.Client
	done   chan bool
}

func NewTodoist(apiKey string) *Todoist {
	td := &Todoist{
		apiKey: apiKey,
		done:   make(chan bool),
	}

	return td
}

func (td *Todoist) Run() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-td.done:
			return
		case <-ticker.C:
			fmt.Println("hoi")
		}
	}
}
