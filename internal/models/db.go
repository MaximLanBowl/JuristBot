package models

import (
	"gorm.io/gorm"
)

type ContactInfo struct {
    gorm.Model
    Address string
    City    string
    Phone   string
}

func Migrate(db *gorm.DB) error {
    err := db.AutoMigrate(&ContactInfo{})
    if err != nil {
        return err
    }

    if db.Migrator().HasTable(&ContactInfo{}) {
        var count int64
        db.Model(&ContactInfo{}).Count(&count)
        if count == 0 {
            // Создаем записи, если таблица пуста
            db.Create(&ContactInfo{
                Address: "Гатчина, ул.Урицкого, д.9А",
                City:    "Гатчина",
                Phone:   "+7 (931) 215-13-00",
            })
            db.Create(&ContactInfo{
                Address: "Санкт-Петербург, Гагаринская 6А, коллегия адвокатов, этаж 3",
                City:    "Санкт-Петербург",
                Phone:   "+7 (931) 215-13-00",
            })
        } else {
            // Обновляем номер телефона для существующих записей
            newPhone := "+7 (931) 215-13-00" // Ваш новый номер телефона
            db.Model(&ContactInfo{}).Where("Phone = ?", "+7 (931) 215-13-00").Update("Phone", newPhone)
        }
    }
    return nil
}