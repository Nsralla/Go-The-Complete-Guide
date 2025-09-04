package db

import (
	"database/sql" // we will use this package to interact with the database
	"fmt"

	_ "github.com/mattn/go-sqlite3" // it must be imported, but will not be used directly, go will use it under the hood
)

var DB *sql.DB
func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")  // Openining a database connection
	// it does not create a connection, it just prepares the database connection
	if err != nil {
		fmt.Println("Error while connecting to the database")
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5) 

	CreateTables()// Create the tables if they do not exist
}

func CreateTables(){
	createEventsTable :=
	 `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
	//Execute the query every time the application starts
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Error while creating events table")
		panic(err)
	}


	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err = DB.Exec(createUsersTable)
	if err != nil{
		fmt.Println("Error while creating users table")
		panic(err)
	}


	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			event_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (event_id) REFERENCES events(id)
		)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		fmt.Println("Error while creating registrations table")
		panic(err)
	}
}