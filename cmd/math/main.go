package main

import (
	"log"
	"net/http"
	"os"

	"exercise1/adapters/http"
	calculator "exercise1/domain/calculator"
)

func main() {
	startServer := os.Args[1]

	calculator := calculator.New()

	router := temphttp.NewRouter(calculator)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

	if startServer == "--web-server" {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		panic("ops, seems like the flag is wrong. try --web-server")
	}
}
