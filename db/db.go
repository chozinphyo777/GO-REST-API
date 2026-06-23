package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // import the SQLite3 driver not used directly, but required for database/sql
)

var DB *sql.DB // global variable to hold the database connection

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // open a connection to the SQLite database
	if err != nil {
		panic("Could not connect to Database.") // panic if there is an error opening the database
	}
	DB.SetMaxOpenConns(10) // set the maximum number of open connections to 10
	DB.SetMaxIdleConns(5)  // keep a maximum of 5 idle connections if no one is using these connections

	createTables()

}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL
	)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}
}
