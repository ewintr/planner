package main

import "os"

func main() {
	td := NewTodoist(os.Getenv("TODOIS_API_TOKEN"))
	td.Run()
}
