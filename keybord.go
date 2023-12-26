package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Start"),
		tgbotapi.NewKeyboardButton("Hello"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Pay now"),
		tgbotapi.NewKeyboardButton("Contact"),
	),
)
