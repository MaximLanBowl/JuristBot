package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() {
    // Настройки формата логов
    logrus.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    
    // Запись логов в файл
    file, err := os.OpenFile("bot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        logrus.SetOutput(file)
    } else {
        logrus.Info("Failed to log to file, using default stderr")
    }
    
    // Установка уровня логирования
    logrus.SetLevel(logrus.InfoLevel)
}