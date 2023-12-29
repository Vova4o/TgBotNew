package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CallBackAnswer обрабатывает нажатие кнопок
// по идее должна писать нужный текст и подсавлять
// нужную клаву
func CallBackAnswer(data string) (string, tgbotapi.InlineKeyboardMarkup) {
	switch data {
	case "start":
		text := "Вы выбрали кнопку 1"
		return text, numericKeyboard
	case "letter":
		text := "Вы выбрали кнопку 2" 
		return text, numericKeyboard
	case "number":
		text := "Вы выбрали кнопку 3" 
		return text, numericKeyboard
	case "enterLetter":
		text := "Вы выбрали кнопку 4" 
		return text, numericKeyboard
	}
	return "Пожалуйста, нажмите *start* для начала работы с ботом", startKeyboard
}
