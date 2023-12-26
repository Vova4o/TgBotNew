package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("Error loading .env file")
		os.Exit(1)
	}
}

// var tableName = os.Getenv("TABLE_NAME")
var ChatIdSQL = os.Getenv("ROW_NAME")

func main() {

	db, err := dbConnect()
	if err != nil {
		fmt.Printf("Error connecting to MySQL: %s", err)
		os.Exit(1)
	}

	defer db.Close()

	bot, err := tgConnect()
	if err != nil {
		fmt.Printf("Error connecting to Telegram: %s", err)
		os.Exit(1)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)
	// Получаем обновления из канала updates
	// и обрабатываем каждое по очереди
	for update := range updates {

		if update.CallbackQuery != nil {
			QueryData := update.CallbackQuery.Data
			if QueryData != "" {
				// Создаем новое сообщение, которое отправим боту
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

				msg.Text, msg.ReplyMarkup = CallBackAnswer(QueryData)
				msg.ParseMode = "markdown"
				// Отправляем
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}
		}
		// Проверяем, что сообщение не пустое
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		DbRecord(&update, db)
		// конструируем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, нажмите *start* для начала работы с ботом")
		// делаем шрифт жирным
		msg.ParseMode = "markdown"
		msg.ReplyMarkup = mainKeyboard

		switch update.Message.Text {
		case "/start":
			msg.Text = "Привет, добро пожаловать в наш бот!\n Ниже представлены кнопки для навигации по боту."
			msg.ReplyMarkup = numericKeyboard
			// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

		}
		// Отправляем сообщение пользователю
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

//**************************************************************************
// Create a new UpdateConfig struct with an offset of 0. Offsets are used
// to make sure Telegram knows we've handled previous values and we don't
// need them repeated.
//updateConfig := tgbotapi.NewUpdate(0)

// Tell Telegram we should wait up to 30 seconds on each request for an
// update. This way we can get information just as quickly as making many
// frequent requests without having to send nearly as many.
//updateConfig.Timeout = 30

// Start polling Telegram for updates.
//updates := bot.GetUpdatesChan(updateConfig)

// Let's go through each update that we're getting from Telegram.
//for update := range updates {
// Telegram can send many types of updates depending on what your Bot
// is up to. We only want to look at messages for now, so we can
// discard any other updates.
//if update.Message == nil {
//	continue
//}

// Now that we know we've gotten a new message, we can construct a
// reply! We'll take the Chat ID and Text from the incoming message
// and use it to create a new message.
//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// We'll also say that this message is a reply to the previous message.
// For any other specifications than Chat ID or Text, you'll need to
// set fields on the `MessageConfig`.
// Ответить на сообщение как реплай
//msg.ReplyToMessageID = update.Message.MessageID

// Okay, we're sending our message off! We don't care about the message
// we just sent, so we'll discard it.
//	if _, err := bot.Send(msg); err != nil {
//		// Note that panics are a bad way to handle errors. Telegram can
//		// have service outages or network errors, you should retry sending
//		// messages or more gracefully handle failures.
//		panic(err)
//	}
//}
