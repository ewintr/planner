package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("cal")

	c := NewClient("http://localhost:8092", "testKey")
	items, err := c.Updated([]Kind{KindEvent}, time.Time{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", items)

	i := Item{
		ID:      "id-1",
		Kind:    KindEvent,
		Updated: time.Now(),
		Body:    "body",
	}
	if err := c.Update([]Item{i}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items, err = c.Updated([]Kind{KindEvent}, time.Time{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", items)
}
