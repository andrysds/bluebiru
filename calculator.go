package main

import (
	"log"
	"os"

	"github.com/andrysds/calculator"
)

func runCalculatorBot() {
	bot, err := newBot(os.Getenv("CALCULATOR_BOT_TOKEN"))
	if err != nil {
		log.Println("Fail on calculator bot authorization:", err)
	}

	updates, err := botUpdatesChan(bot)
	if err != nil {
		log.Println("Fail on getting calculator bot updates:", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		var msgText string
		if update.Message.IsCommand() && update.Message.Command() == "calc" {
			msgText, err = calculator.Calculate(update.Message.CommandArguments())
			if err != nil {
				msgText = err.Error()
			}
		}
		bot.Send(newTextMessage(&update, msgText))
	}
}
