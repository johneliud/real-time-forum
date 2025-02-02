package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./forum.db")
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
		name TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Printf("Failed executing query: %v\n", err)
		return
	}
	log.Println("Database initialized successfully")
}
