package controller

import (
	"log"
	"net/http"

	"github.com/johneliud/forum/database"
)

/*
LogoutHandler deletes the session cookie from the databse when a user logs out.
*/
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println("Redirecting to '/sign-in'")
		http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
		return
	}

	_, err = database.DB.Exec("UPDATE users SET session_token = NULL WHERE session_token = ?", cookie.Value)
	if err != nil {
		log.Printf("Failed to clear session cookie: %v\n", err)
		ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
}
