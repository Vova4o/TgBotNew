package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// dbConnect function will connect to the database
// check if connection is successful return connection and nil
// also check if table tgusers created and create it if not
func dbConnect() (*sql.DB, error) {
	var user = os.Getenv("LOGIN")
	var password = os.Getenv("PASSWORD")
	var ip = os.Getenv("IP")
	var port = os.Getenv("PORT")
	var dbName = os.Getenv("DBNAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, ip, port, dbName)
	fmt.Println(connectionString)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Printf("Error connecting to MySQL: %s\n", err)
		return nil, err
	}

	// Ping the database to ensure connection is successful
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging database: %s\n", err)
		return nil, err
	}

	// Check if table exists and create it if it doesn't
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tgusers (
        id INT AUTO_INCREMENT PRIMARY KEY,
        UserIDSQL BIGINT NOT NULL,
        UserName VARCHAR(255)
    )`)
	if err != nil {
		fmt.Printf("Error creating table: %s\n", err)
		return nil, err
	}

	fmt.Printf("Connected to database %s\n", dbName)

	return db, nil
}
