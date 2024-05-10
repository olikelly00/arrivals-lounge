package main

import (
	"arrivals_lounge/flights"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("Welcome to the Arrivals Lounge!\n")
	fmt.Printf("Here are the upcoming arrivals at your chosen airport: %s\n", os.Args[1])
	airports := os.Args[1:]
	for _, airport := range airports {
		flights.Display(flights.GetArrivals(airport))
	}
}
