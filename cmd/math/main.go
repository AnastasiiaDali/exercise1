package main

import (
	"fmt"
	"log"
	"net/http"

	"exercise1/adapters/http"
	calculator "exercise1/domain/calculator"
)

func main() {

	calculator := calculator.New()

	router := temphttp.NewRouter(calculator)

	fmt.Printf("Starting server on port 8081...\n")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}

}
