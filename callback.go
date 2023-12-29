package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CallBackAnswer обрабатывает нажатие кнопок
// по идее должна писать нужный текст и подсавлять
// нужную клаву
func CallBackAnswer(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	switch data {
	case "firstLesson":
		text := "Добро пожаловать на первый урок, ниже ссылка на PDF файл с уроком.\n" +
			"https://drive.google.com/file/d/1v43vjHqzVqfrN4lT7OOLKJB2VALwCUnV/view"
		return text, numericKeyboard2
	case "secondLesson":
		text := "Вы выбрали кнопку 2"
		return text, numericKeyboard3
	case "button":
		text := "Вы выбрали кнопку Кнопка"
		return text, numericKeyboard
	case "contact":
		text := "Вы выбрали кнопку Contact"
		return text, numericKeyboard
	}
	return "Пожалуйста, нажмите *start* для начала работы с ботом", startKeyboard
}
