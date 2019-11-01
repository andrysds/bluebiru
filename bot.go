package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func newBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err == nil {
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}
	return bot, err
}

func botUpdatesChan(bot *tgbotapi.BotAPI) (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}
