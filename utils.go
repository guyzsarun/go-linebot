package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func load_dev_env() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	LINE_API := os.Getenv("LINE_API")
	LINE_SECRET := os.Getenv("LINE_SECRET")
	return LINE_API, LINE_SECRET
}
