package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

// В данной функции мы подключаемся к боту
// выдаем на экран сообщение, что подключение произведено
// Если что-то пошло не так, выдаем сообщение об ошибке
func tgConnect() (*tgbotapi.BotAPI, error) {

	//bot, err := tgbotapi.NewBotAPI(os.Getenv())
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot, nil
}
