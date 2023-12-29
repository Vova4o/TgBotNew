package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// DbRecord проверяет если юзер есть в базе, если нет, то добавляет
// сделано это для того чтобы потом рассылать через бота сообщения
// тем юзерам которые есть в базе
func DbRecord(update *tgbotapi.Update, db *sql.DB) {
	tableName := os.Getenv("TABLE_NAME")
	row1 := os.Getenv("ROW_1")
	row2 := os.Getenv("ROW_2")
	row3 := os.Getenv("ROW_3")
	if update.Message.From.ID != 0 {
		userID := update.Message.From.ID
		userName := update.Message.Chat.UserName
		time := update.Message.Time()

		// Check if user ID already exists
		var exists bool
		query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s=?)", tableName, row1)
		err := db.QueryRow(query, userID).Scan(&exists)
		if err != nil {
			log.Fatal(err)
		}

		// If user ID does not exist, insert it
		if !exists {

			connect := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES(?, ?, ?)", tableName, row1, row2, row3)
			stmt, err := db.Prepare(connect)
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			// Передаем в запрос переменные из update ЮзерАйДи, Имя Юзера и Время
			_, err = stmt.Exec(userID, userName, time)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// func DbRecord(update *tgbotapi.Update, db *sql.DB) {
// 	if update.Message.Chat.ID != 0 {
// 		// Prepare statement for inserting data
// 		stmt, err := db.Prepare("INSERT INTO tgusers(UserIDSQL, UserName) VALUES(?, ?)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer stmt.Close()

// 		// Execute SQL statement
// 		_, err = stmt.Exec(update.Message.From.ID, update.Message.Chat.UserName)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
