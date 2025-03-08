package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

/*
InitDB installs the neccessary drivers and environment required to run the database.
*/
func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "data/forum.db")
	if err != nil {
		log.Printf("Failed to open database: %v\n", err)
		return
	}

	if err = DB.Ping(); err != nil {
		log.Printf("Database connection failed: %v", err)
		return
	}

	query := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT UNIQUE NOT NULL,
    last_name TEXT UNIQUE NOT NULL,
    nick_name TEXT UNIQUE NOT NULL,
    gender TEXT UNIQUE NOT NULL,
    age INTEGER NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    session_token TEXT UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Printf("Failed executing query: %v\n", err)
		return
	}
	log.Println("Database initialized successfully")
}
