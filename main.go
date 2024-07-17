package main

import (
	"TheBestBotEver/httpServer"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Authorization interface {
}
type TodoList interface {
}
type TodoIten interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoIten
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}

func main() {
	go func() {
		httpServer.Run()
	}()
	bot, err := tgbotapi.NewBotAPI("6843460530:AAHYsOhi4wdRXFGn40M3nmSrfPlbEZohFSY")
	if err != nil {
		log.Panic(err)
	}

	db, err := NewPostgresDB(Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
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

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
