package main

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func text_handler(event linebot.Event) {
	var err error
	userID := event.Source.UserID
	replyToken := event.ReplyToken

	log.Print("Message from  ", userID)

	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
			log.Print(err)
		}
	case *linebot.StickerMessage:
		replyMessage := fmt.Sprintf(
			"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
		if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
			log.Print(err)
		}
	}

}
