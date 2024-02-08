package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6843460530:AAHYsOhi4wdRXFGn40M3nmSrfPlbEZohFSY")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, handlemsg(update.Message.Text))
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
func handlemsg(msg string) string {
	if msg == "привет" {
		return "abc"
	}
	if msg == "добрый день" {
		return "def"
	}
	if msg == "добрый вечер" {
		return "ghi"
	}
	return "jkl"
}
