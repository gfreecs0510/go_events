package clients

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data/api.db")

	if err != nil {
		panic("Cannot connect to db")
	}

	fmt.Println("Db connected")

	_, err = DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic("Failed to enable foreign key constraints")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	q := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, err := DB.Exec(q)

	if err != nil {
		fmt.Println(err)
		panic("users table not created")
	}

	q = `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        created_at DATETIME NOT NULL,
        user_id INTEGER,
    	FOREIGN KEY (user_id) REFERENCES user(id)
    );`

	_, err = DB.Exec(q)

	if err != nil {
		panic("events table not created")
	}

	q = `
	CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
	created_at DATETIME NOT NULL,
	UNIQUE(event_id, user_id),
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(q)

	if err != nil {
		panic("registrations table not created")
	}
}
