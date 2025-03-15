package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/johneliud/real-time-forum/backend/logger"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func executeSchema(DB *sql.DB) error {
	content, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema: %w", err)
	}

	// Split SQL statements and execute each
	statements := string(content)
	_, err = DB.Exec(statements)
	if err != nil {
		return fmt.Errorf("failed to execute SQL statements: %w", err)
	}
	return nil
}

/*
InitDB installs the neccessary drivers and environment required to run the database.
*/
func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "data/forum.db")
	if err != nil {
		logger.Error("Failed to open database:", err)
		return
	}

	if err = DB.Ping(); err != nil {
		logger.Error("Connection to database failed:", err)
		return
	}

	if err = executeSchema(DB); err != nil {
		logger.Error("failed to execute SQL file:", err)
		return
	}
	logger.Info("Database initialized successfully")
}
