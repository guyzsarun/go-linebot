package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

func handler(c *gin.Context) {
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		log.Print(err)
	}
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			log.Println("EventTypeMessage")
			text_handler(*event)
		case linebot.EventTypeFollow:
			log.Println("EventTypeFollow")
		case linebot.EventTypeUnfollow:
			log.Println("EventTypeUnfollow")
		case linebot.EventTypeJoin:
			log.Println("EventTypeJoin")
		case linebot.EventTypeLeave:
			log.Println("EventTypeLeave")
		case linebot.EventTypePostback:
			log.Println("EventTypePostback")
		case linebot.EventTypeBeacon:
			log.Println("EventTypeBeacon")
		default:
			log.Println("Event type not supported")
		}

	}
}
