package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

var LINE_API string
var LINE_SECRET string

var bot *linebot.Client

func greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func main() {
	var err error
	LINE_API, LINE_SECRET = load_dev_env()
	bot, err = linebot.New(LINE_SECRET, LINE_API)

	log.Println("Bot:", bot, " err:", err)

	r := setupRouter()
	r.Run()
}
