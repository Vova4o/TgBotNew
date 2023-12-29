package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var commonKeyboard = tgbotapi.NewInlineKeyboardRow(
	tgbotapi.NewInlineKeyboardButtonData("Кнопка", "button"),
	tgbotapi.NewInlineKeyboardButtonData("Contact", "contact"),
)

// Тута надо нафигачить кучу клавиатур
var startKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("С начала", "start"),
	),
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Первый урок", "firstLesson"),
	),
	commonKeyboard,
)

var numericKeyboard2 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Второй урок", "secondLesson"),
	),
	commonKeyboard,
)

var numericKeyboard3 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Третий урок", "thirdLesson"),
	),
	commonKeyboard,
)

var mainKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
	),
)
