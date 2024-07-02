package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	println("TG Bot starting...")
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	if tgBotToken == "" {
		panic("TG_BOT_TOKEN not set")
	}
	tgBot, err := tgbotapi.NewBotAPI(tgBotToken)
	if err != nil {
		panic(err)
	}
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 10
	updateChan := tgBot.GetUpdatesChan(updateConfig)
	for update := range updateChan {
		if update.Message != nil {
			HandleMessage(tgBot, update)
		}
	}
}
