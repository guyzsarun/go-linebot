package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func load_dev_env() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	LINE_API := os.Getenv("LINE_API")
	LINE_SECRET := os.Getenv("LINE_SECRET")
	return LINE_API, LINE_SECRET
}

func getPort() string {
	var port = os.Getenv("PORT") // ----> (A)
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port // ----> (B)
}
