package main

import (
	"encoding/json"
	"log"

	"github.com/mavihq/persian"
	tele "gopkg.in/telebot.v3"
)

func SanitizePersianDigits(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		msg := c.Message()
		msg.Text = persian.ToEnglishDigits(msg.Text)
		return next(c)
	}
}

func logUserMessage(next tele.HandlerFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		b, err := json.MarshalIndent(c.Message(), "", "\t")
		if err != nil {
			log.Println(err)
		}else{
			log.Println(string(b))
		}
		
		return next(c)
	}
}
