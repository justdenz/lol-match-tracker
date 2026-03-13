package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }
}

var (
	apiKey = os.Getenv("API_KEY")
)