package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CallBackAnswer(data tgbotapi.CallbackQuery) string {
	switch data.Data {
	case "1":
		return "Вы выбрали кнопку 1"
	case "2":
		return "Вы выбрали кнопку 2"
	case "3":
		return "Вы выбрали кнопку 3"
	}
	return ""
}
