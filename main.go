package main

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

// func init() {
// 	err := godotenv.Load()
// 	LogIfError(err, "Error loading .env file")
// }

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	LogIfError(err)

	db := ConnectMongoDB()
	chatRepo := NewChatRepo(db)
	handler := NewHandler(chatRepo)

	b.Handle("/start", handler.Start)

	b.Start()
}
