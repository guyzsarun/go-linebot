package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		if message.Text == "choose" {
			return_msg = random_text(message.Text)
		} else if strings.Contains(message.Text, "profile") {
			res, _ := bot.GetProfile(userID).Do()
			user_info := User_info{
				res.UserID,
				res.DisplayName,
				res.Language,
				res.PictureURL,
				res.StatusMessage}
			return_msg = user_info.print_info()
		} else if strings.ToLower(message.Text) == "status" {
			wallet := get_wallet()
			resp, err := http.Get("https://api.ethermine.org/miner/:" + wallet + "/currentStats")
			if err != nil {
				return_msg = "Cannot get miner status"
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			miner_status := Miner{}
			json.Unmarshal(body, &miner_status)
			returnStruct := miner_status.postProcess()
			return_msg = "Miner :" + wallet + "\n" + returnStruct.print_info()

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
	case *linebot.LocationMessage:
		return_msg := fmt.Sprintf("Latitude: %f\nLongtitude: %f", message.Latitude, message.Longitude)
		if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(return_msg)).Do(); err != nil {
			log.Print(err)
		}

	}
}
