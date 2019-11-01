package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func runExampleBot() {
	bot, err := newBot(os.Getenv("EXAMPLE_BOT_TOKEN"))
	if err != nil {
		log.Println("Fail on example bot authorization:", err)
	}

	updates, err := botUpdatesChan(bot)
	if err != nil {
		log.Println("Fail on getting example bot updates:", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
