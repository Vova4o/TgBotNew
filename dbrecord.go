package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// DbRecord проверяет если юзер есть в базе, если нет, то добавляет
// сделано это для того чтобы потом рассылать через бота сообщения
// тем юзерам которые есть в базе
func DbRecord(update *tgbotapi.Update, db *sql.DB) {
    if update.Message.Chat.ID != 0 {
        userID := update.Message.From.ID
		userName := update.Message.Chat.UserName

        // Check if user ID already exists
        var exists bool
        err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM tgusers WHERE UserIDSQL=?)", userID).Scan(&exists)
        if err != nil {
            log.Fatal(err)
        }

        // If user ID does not exist, insert it
        if !exists {
            stmt, err := db.Prepare("INSERT INTO tgusers(UserIDSQL, UserName) VALUES(?, ?)")
            if err != nil {
                log.Fatal(err)
            }
            defer stmt.Close()

            _, err = stmt.Exec(userID, userName)
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
