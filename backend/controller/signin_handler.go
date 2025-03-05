package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/johneliud/real-time-forum/database"
	"golang.org/x/crypto/bcrypt"
)

/*
SigninHandler function handles the sign in logic by validating if a user exists in the database.
*/
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-in" {
		log.Printf("Failed to find path %q\n", r.URL.Path)
		ErrorHandler(w, "Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("frontend/templates/sign-in.html")
		if err != nil {
			log.Printf("Failed parsing template: %v\n", err)
			ErrorHandler(w, "Page Not Found", http.StatusNotFound)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("Failed to execute template: %v\n", err)
			ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			log.Printf("Failed parsing form: %v\n", err)
			ErrorHandler(w, "Bad Request", http.StatusBadRequest)
			return
		}

		email, password := r.FormValue("email"), r.FormValue("password")

		if len(strings.TrimSpace(email)) == 0 {
			fmt.Println("Email value cannot be an empty string")
			return
		}

		if len(strings.TrimSpace(password)) == 0 {
			fmt.Println("Password value cannot be an empty string")
			return
		}

		var (
			userID       int
			storedHash   string
			sessionToken sql.NullString
		)

		// Check user existance in the db
		err := database.DB.QueryRow("SELECT id, password, session_token FROM users WHERE email = ?", email).Scan(&userID, &storedHash, &sessionToken)
		if err == sql.ErrNoRows {
			log.Printf("Invalid credentials. No user found: %v\n", err)
			ErrorHandler(w, "Unauthorized User", http.StatusUnauthorized)
			return
		} else if err != nil {
			log.Printf("Failed quering database row: %v\n", err)
			ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}

		// Compare hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
			log.Printf("Invalid credentials: %v\n", err)
			ErrorHandler(w, "Unauthorized User", http.StatusUnauthorized)
			return
		}

		newSessionToken := uuid.New().String()

		_, err = database.DB.Exec("UPDATE users SET session_token = ? WHERE id = ?", newSessionToken, userID)
		if err != nil {
			log.Printf("Failed to update session token: %v\n", err)
			ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    newSessionToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	default:
		log.Println("Invalid request method")
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
