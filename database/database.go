package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/backend/model"
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

// InitDB installs the neccessary drivers and environment required to run the database.
func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "files/real_time_forum.db")
	if err != nil {
		logger.Error("Failed to open database", "err", err)
		return
	}

	if err = DB.Ping(); err != nil {
		logger.Error("Connection to database failed:", "err", err)
		return
	}

	if err = executeSchema(DB); err != nil {
		logger.Error("failed to execute SQL file:", "err", err)
		return
	}
	logger.Info("Database initialized successfully")
}

// InsertMessage inserts a new message into the database
func InsertMessage(content string, sender string) error {
	stmt, err := DB.Prepare("INSERT INTO messages(content, sender) VALUES(?, ?)")
	if err != nil {
		logger.Error("Failed to prepare statement", "err", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, sender)
	return err
}

// GetMessages retrieves all messages from the database
func GetMessages() ([]map[string]interface{}, error) {
	rows, err := DB.Query("SELECT id, content, sender, timestamp FROM messages ORDER BY timestamp DESC")
	if err != nil {
		logger.Error("Failed to query messages", "err", err)
		return nil, err
	}
	defer rows.Close()

	var messages []map[string]interface{}
	for rows.Next() {
		var id int
		var content, sender string
		var timestamp string
		if err := rows.Scan(&id, &content, &sender, &timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, map[string]interface{}{
			"id":        id,
			"content":   content,
			"sender":    sender,
			"timestamp": timestamp,
		})
	}
	return messages, nil
}

// GetUserProfile retrieves the user's profile data based on user ID
func GetUserProfile(userID string) (*model.User, error) {
	var user model.User

	row := DB.QueryRow("SELECT id, first_name, last_name, nick_name, gender, age, email FROM users WHERE id = ?", userID)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.NickName, &user.Gender, &user.Age, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("User not found in database", "err", err)
			return nil, nil
		}
		logger.Error("Error fetching user profile", "err", err)
		return nil, err
	}
	return &user, nil
}
