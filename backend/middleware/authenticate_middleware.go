package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/database"
)

/*
AuthenticateMiddleware ensures that a user is authenticated before accessing certain routes.
*/
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			// log.Println("Redirecting to '/sign-in'")
			// http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}

		var userID int
		err = database.DB.QueryRow("SELECT id FROM users WHERE session_token = ?", cookie.Value).Scan(&userID)
		if err == sql.ErrNoRows {
			// log.Println("Redirecting to '/sign-in'")
			// http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		} else if err != nil {
			log.Printf("Error checking session: %v\n", err)
			controller.ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
