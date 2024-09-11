package main

import "fmt"

func main() {
	fmt.Println("cal")

	c := sync.NewClient("https://localhost:8092", "testKey", "server.crt")
}
