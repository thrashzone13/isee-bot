package main

import (
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	chatRepo *ChatRepo
}

func NewHandler(chatRepo *ChatRepo) *Handler {
	return &Handler{
		chatRepo,
	}
}

func (h *Handler) Start(c tele.Context) error {
	chatID := c.Chat().ID
	// if
	// h.chatRepo.Find(chatID)
	h.chatRepo.Create(chatID)

	return c.Send("Done")
}
