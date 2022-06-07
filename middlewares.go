package main

import (
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
