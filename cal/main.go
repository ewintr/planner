package main

import (
	"fmt"
	"time"

	"code.ewintr.nl/planner/sync"
)

func main() {
	fmt.Println("cal")

	c := sync.NewClient("https://localhost:8092", "testKey", "server.crt")
	items, err := c.Updates([]sync.Kind{sync.KindEvent}, time.Time{})
}
