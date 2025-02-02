package controller

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/johneliud/forum/database"
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

		if len(strings.TrimSpace(email)) == 0 || len(strings.TrimSpace(password)) == 0 {
			log.Println("Failed validating form values")
			ErrorHandler(w, "Bad Request", http.StatusBadRequest)
			return
		}

		var hashedPassword string

		// Check user existance in the db
		err := database.DB.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&hashedPassword)
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
		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
			log.Printf("Invalid credentials: %v\n", err)
			ErrorHandler(w, "Unauthorized User", http.StatusUnauthorized)
			return
		}

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		log.Println("Invalid request method")
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
