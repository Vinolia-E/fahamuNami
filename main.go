package main

// import "fmt"

import (
	"log"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Welcome to Fahamu Nami!")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

}
