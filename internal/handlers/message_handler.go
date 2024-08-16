package handlers

import (
	"TelegramBot/internal/services"
	"TelegramBot/internal/utils"
	"TelegramBot/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var service *services.Service

func SetupMessageHandlers(bot *telegram.Bot) {
    // Инициализация сервиса с подключением к БД
    service = services.NewService(utils.DB)

    bot.Handle("start", func(message *tgbotapi.Message) {
        handleStartCommand(bot, message)
    })

    bot.Handle("consultation", func(message *tgbotapi.Message) {
        handleConsultationCommand(bot, message)
    })

    bot.Handle("bankruptcy_individual", func(message *tgbotapi.Message) {
        handleBankruptcyIndividualCommand(bot, message)
    })

    bot.Handle("bankruptcy_corporate", func(message *tgbotapi.Message) {
        handleBankruptcyCorporateCommand(bot, message)
    })

    bot.Handle("legal_disputes", func(message *tgbotapi.Message) {
        handleLegalDisputesCommand(bot, message)
    })

    bot.Handle("contact_info", func(message *tgbotapi.Message) {
        handleContactInfoCommand(bot, message)
    })


    bot.Handle("", func(message *tgbotapi.Message) {
        handleUnknownCommand(bot, message)
    })
}

func handleStartCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetHelpMessage()
	bot.SendMessage(response, message.Chat.ID)
}

func handleConsultationCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetConsultationInfo()
    bot.SendMessage(response, message.Chat.ID)
}



func handleBankruptcyIndividualCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetBankruptcyIndividualInfo()
    bot.SendMessage(response, message.Chat.ID)
}

func handleBankruptcyCorporateCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetBankruptcyCorporateInfo()
    bot.SendMessage(response, message.Chat.ID)
}

func handleLegalDisputesCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetLegalDisputesInfo()
    bot.SendMessage(response, message.Chat.ID)
}


func handleContactInfoCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := service.GetContactInfo()
    bot.SendMessage(response, message.Chat.ID)
}



func handleUnknownCommand(bot *telegram.Bot, message *tgbotapi.Message) {
    response := "Извините, я не понимаю эту команду. Пожалуйста, используйте /start для получения списка доступных команд."
    bot.SendMessage(response, message.Chat.ID)
}