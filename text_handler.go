package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func text_handler(event linebot.Event) {
	var err error
	userID := event.Source.UserID
	replyToken := event.ReplyToken

	log.Print("Message from  ", userID)

	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		var return_msg string
		if strings.Contains(message.Text, "choose") {
			return_msg = random_text(message.Text)
		} else {
			return_msg = message.Text
		}
		if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(return_msg)).Do(); err != nil {
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
