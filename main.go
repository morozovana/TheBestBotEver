package main

import (
	"log"
	"strings"

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
			ans := handlemsg(update.Message.Text)
			for i := 0; i < len(ans); i++ {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans[i])
				msg.ReplyToMessageID = update.Message.MessageID
				_, err := bot.Send(msg)
				if err != nil {
					log.Panic(err)
				}
			}

		}
	}
}
func handlemsg(msg string) []string {
	var answers []string
	switch strings.ToLower(msg) {
	case "привет":
		answers = append(answers, "сейчас напишу abc", "abc")
	case "добрый день":
		answers = append(answers, "def")
	case "добрый вечер":
		answers = append(answers, "ghi")
	default:
		answers = append(answers, "jkl")
	}
	return answers
}
