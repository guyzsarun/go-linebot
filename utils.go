package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku" + port)
	}
	return ":" + port
}

func random_text(text string) string {
	rand.Seed(time.Now().Unix())
	s := strings.Split(strings.TrimSpace(text), " ")
	return s[rand.Intn(len(s)-1)+1]
}
