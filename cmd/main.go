package main

import (
	"TelegramBot/internal/config"
	"TelegramBot/internal/handlers"
	"TelegramBot/internal/models"
	"TelegramBot/internal/utils"
	"TelegramBot/pkg/telegram"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        logrus.Fatalf("Error loading .env file")
    }
    utils.SetupLogger()

    cfg, err := config.LoadConfig()
    if err != nil {
        logrus.Fatalf("Error loading config: %v", err)
    }

    db, err := utils.ConnectDatabase(cfg)
    if err != nil {
        logrus.Fatalf("Error connecting to database: %v", err)
    }

    if err := models.Migrate(db); err != nil {
        logrus.Fatalf("Error migrating the database: %v", err)
    }

    utils.DB = db

    bot, err := telegram.NewBot(cfg.TelegramToken)
    if err != nil {
        logrus.Fatalf("Error initializing Telegram Bot: %v", err)
    }
    handlers.SetupMessageHandlers(bot)

    logrus.Info("Bot started...")
    bot.Start()
}
