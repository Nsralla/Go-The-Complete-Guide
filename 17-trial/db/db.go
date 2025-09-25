package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		fmt.Println("Error while connecting to a db.", err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}


func createTables() {
	eventsQuery := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCReMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER NOT NULL
		) 
	`
	_, err := DB.Exec(eventsQuery)
	if err != nil {
		fmt.Println("Error creating events table.", err)
	}
}
