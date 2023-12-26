package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CallBackAnswer(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	switch data {
	case "1":
		return "Вы выбрали кнопку 1", numericKeyboard
	case "2":
		return "Вы выбрали кнопку 2", numericKeyboard
	case "3":
		return "Вы выбрали кнопку 3", numericKeyboard
	}
	return "Пожалуйста, нажмите *start* для начала работы с ботом", startKeyboard
}
