package services

import (
	"TelegramBot/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type Service struct {
    DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{DB: db}
}

func (s *Service) GetConsultationInfo() string {
    return "Консультации проводятся с понедельника по пятницу с 10:00 до 19:00. Пожалуйста, запишитесь заранее."
}

func (s *Service) GetBankruptcyIndividualInfo() string {
    return "Банкротство физических лиц возможно при соблюдении определенных условий. Свяжитесь с нами для получения подробной информации."
}

func (s *Service) GetBankruptcyCorporateInfo() string {
    return "Мы предлагаем услуги по банкротству юридических лиц. Пожалуйста, обратитесь к нам для получения консультации."
}

func (s *Service) GetLegalDisputesInfo() string {
    return "Мы предоставляем правовую помощь в юридических спорах, включая представительство в суде. Свяжитесь с нами для более подробной информации."
}

func (s *Service) GetContactInfo() string {
    var contactInfos []models.ContactInfo
    s.DB.Find(&contactInfos)

    var response string
    for _, contact := range contactInfos {
        response += fmt.Sprintf("Адрес: %s, %s. Телефон: %s\n", contact.City, contact.Address, contact.Phone)
    }
    return response
}

func (s *Service) GetHelpMessage() string {
    return `
    "Добро пожаловать в юридический бот! Я здесь, чтобы помочь вам с юридическими вопросами. Используйте команды для получения различных услуг."

    
    Доступные команды:
    /start - Начать использовать бота
    /consultation - Получить информацию о консультациях
    /bankruptcy_individual - Информация о банкротстве физических лиц
    /bankruptcy_corporate - Информация о банкротстве юридических лиц
    /legal_disputes - Информация о юридических спорах
    /contact_info - Контактная информация, в случае обращения(назовите промокод ЮРИСТ-47, вам будет предоставлена скидка в размере 5%)

    Узнать больше о нас можно на нашем сайте: [jurist47.ru](https://jurist47.ru/) также если обращаетесь через телеграмм, называйте промокод - ЮРИСТ-47
    `
}