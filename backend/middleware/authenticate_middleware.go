package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/database"
)

type contextKey string

const userIDKey contextKey = "userID"

var publicRoutes = map[string]bool{
	"/api/sign-in": true,
	"/api/sign-up": true,
}

// IsPublicRoute checks if a route is public (doesn't require authentication).
func IsPublicRoute(path string) bool {
	return publicRoutes[path]
}

// IsAPIRoute checks if a route is an API route.
func IsAPIRoute(path string) bool {
	uiRoutes := map[string]bool{
		"/":        true,
		"/sign-up": true,
		"/sign-in": true,
	}

	return strings.Contains(path, "/api/") && !uiRoutes[path]
}

// sendUnauthorizedResponse writes a JSON response for unauthorized access.
func sendUnauthorizedResponse(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(`{"authenticated": false, "message": "` + message + `"}`))
}

// AuthenticateMiddleware ensures that a user is authenticated before accessing certain routes.
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions || IsPublicRoute(r.URL.Path) || strings.HasPrefix(r.URL.Path, "/ws") {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the request is an AJAX request
		isXHR := r.Header.Get("X-Requested-With") == "XMLHttpRequest"

		cookie, err := r.Cookie("session_token")
		if err != nil {
			if IsAPIRoute(r.URL.Path) || isXHR {
				logger.Warn("Unauthorized to access route", "path", r.URL.Path, "err", err)
				sendUnauthorizedResponse(w, "Not authenticated")
				return
			}
			http.ServeFile(w, r, "./frontend/templates/index.html")
			return
		}

		var userID int
		err = database.DB.QueryRow("SELECT id FROM users WHERE session_token = ?", cookie.Value).Scan(&userID)
		if err == sql.ErrNoRows {
			// Clear invalid cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "session_token",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   r.TLS != nil,
				SameSite: http.SameSiteLaxMode,
			})

			if IsAPIRoute(r.URL.Path) || isXHR {
				logger.Warn("Unauthorized to access route", "path", r.URL.Path, "err", err)
				sendUnauthorizedResponse(w, "Not authenticated")
				return
			}
			http.ServeFile(w, r, "./frontend/templates/index.html")
			return
		} else if err != nil {
			logger.Error("Error checking session", "err", err)
			controller.ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}

		// Add userID to the context as a string
		ctx := context.WithValue(r.Context(), userIDKey, strconv.Itoa(userID))
		fmt.Println("ctx value:", ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
