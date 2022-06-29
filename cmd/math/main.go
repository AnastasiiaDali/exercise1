package main

import (
	"os"

	"exercise1/adapters/http"
)

func main() {
	startServer := os.Args[1]

	if startServer == "--web-server" {
		http.MathServer()
	} else {
		panic("ops, seems like the flag is wrong. try --web-server")
	}
}
