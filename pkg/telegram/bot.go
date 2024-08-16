package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	API      *tgbotapi.BotAPI
	commands map[string]func(msg *tgbotapi.Message)
}

func NewBot(token string) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Fatal("Error to create bot", err)
	}
	return &Bot{API: botAPI, commands: make(map[string]func(msg *tgbotapi.Message))}, nil
}

func (bot *Bot) Handle(cmd string, handler func(msg *tgbotapi.Message)) {
	bot.commands[cmd] = handler
}

func (bot *Bot) SendMessage(text string, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.API.Send(msg); err != nil {
		logrus.Error("Failed to send message", err)
	}
}

func (bot *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.API.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil || update.Message.IsCommand() {
			handler, exists := bot.commands[update.Message.Command()]
			if exists {
				handler(update.Message)
			} else {
				if defaultHandler, exists := bot.commands[""]; exists{

					defaultHandler(update.Message)
				}
			}
		}
	}
	logrus.Info("Bot started")
}