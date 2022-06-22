package main

import (
	"os"

	"exercise1/adapters/http"
)

func main() {
	startServer := os.Args[1]

	if startServer == "web-server" {
		http.MathServer()
	}
}
