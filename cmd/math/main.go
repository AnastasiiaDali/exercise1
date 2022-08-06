package main

import (
	"fmt"
	"log"
	"net/http"

	"exercise1/adapters/http"
	calculator "exercise1/domain/calculator"
	fibonacci "exercise1/domain/fibonacci"
)

func main() {
	calculator := calculator.New()
	fibonacci := fibonacci.NewFibonacci()

	router := temphttp.NewRouter(calculator, fibonacci)

	fmt.Printf("Starting server on port 8081...\n")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}

}
