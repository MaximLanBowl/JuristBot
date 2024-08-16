package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
    TelegramToken string
    DBHost        string
    DBPort        string
    DBUser        string
    DBPassword    string
    DBName        string
}

func LoadConfig() (*Config, error) {
    token := os.Getenv("TELEGRAM_BOT_TOKEN")
    if token == "" {
        logrus.Fatal("TELEGRAM_BOT_TOKEN is missing")
    }
    
    return &Config{
        TelegramToken: token,
        DBHost:        os.Getenv("DB_HOST"),
        DBPort:        os.Getenv("DB_PORT"),
        DBUser:        os.Getenv("DB_USER"),
        DBPassword:    os.Getenv("DB_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
    }, nil
}