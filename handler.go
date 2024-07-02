package main

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(tgBot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	switch parseCMD(update.Message.Text) {
	case "/hello":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
		tgBot.Send(msg)
	case "/chat_id":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Chat ID: %d", update.Message.Chat.ID))
		tgBot.Send(msg)
	default:
	}
}

func parseCMD(text string) string {
	textBlocks := strings.Split(text, " ")
	if len(textBlocks) > 0 && strings.HasPrefix(textBlocks[0], "/") {
		return textBlocks[0]
	}
	return ""
}
