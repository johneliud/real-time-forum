package controller

import (
	"log"
	"net/http"
	"text/template"

	"github.com/johneliud/forum/backend/util"
	"github.com/johneliud/forum/database"
)

/*
SignupHandler handles the main sign up logic by retrieving form data, performing validations and hashing sensitive data for secure storage in the database.
*/
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sign-up" {
		log.Printf("Failed to find path %q\n", r.URL.Path)
		ErrorHandler(w, "Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := template.ParseFiles("frontend/templates/sign-up.html")
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

		name, email, password, confirmedPassword := r.FormValue("name"), r.FormValue("email"), r.FormValue("password"), r.FormValue("confirmed-password")

		if err := util.ValidateInput(name, email, password, confirmedPassword); err != nil {
			log.Printf("Failed validating form values: %v\n", err)
			ErrorHandler(w, "Bad Request", http.StatusBadRequest)
			return
		}

		hashedPassword, err := util.HashPassword(password)
		if err != nil {
			log.Printf("Failed hashing passord: %v\n", err)
			ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}

		if database.DB == nil {
			log.Println("Database not initialized")
			return
		}

		stmt, err := database.DB.Prepare("INSERT INTO users(name, email, password) VALUES (?, ?, ?)")
		if err != nil {
			log.Printf("Failed preparing statement to database: %v\n", err)
			ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, email, hashedPassword)
		if err != nil {
			log.Printf("Failed executing statement: %v\n", err)
			ErrorHandler(w, "Bad Request", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
	default:
		log.Println("Invalid request method")
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
