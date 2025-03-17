package middleware

import (
	"database/sql"
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/backend/logger"
	"github.com/johneliud/real-time-forum/database"
)

// PublicRoutes is a list of routes that don't require authentication
var publicRoutes = map[string]bool{
	"/":                true,
	"/api/sign-in":     true,
	"/api/sign-up":     true,
	"/api/validate":    true,
	"/api/auth-status": true,
	"/frontend/":       true, // Prefix match for static files
}

// IsPublicRoute checks if a route is public (doesn't require authentication).
func IsPublicRoute(path string) bool {
	if publicRoutes[path] {
		return true
	}

	// Special case for static files with the /frontend/ prefix
	if len(path) >= 10 && path[:10] == "/frontend/" {
		return true
	}

	return false
}

// IsAPIRoute checks if a route is an API route.
func IsAPIRoute(path string) bool {
	return len(path) >= 5 && path[:5] == "/api/"
}

// AuthenticateMiddleware ensures that a user is authenticated before accessing certain routes.
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the requested path is public
		if IsPublicRoute(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("auth_token")
		if err != nil {
			// Handle authentication failures differently based on request type
			isXHR := r.Header.Get("X-Requested-With") == "XMLHttpRequest"

			if IsAPIRoute(r.URL.Path) || isXHR {
				// Return 401 Unauthorized for API routes or AJAX requests
				logger.Info("No auth_token cookie found for API request, returning 401")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"authenticated": false, "message": "Not authenticated"}`))
			} else {
				// For regular browser requests, handle within the SPA by serving index.html
				logger.Info("No auth_token cookie found, serving SPA")
				http.ServeFile(w, r, "./frontend/templates/index.html")
			}
			return
		}

		// Verify the session token in the database
		var userID int
		err = database.DB.QueryRow("SELECT id FROM users WHERE session_token = ?", cookie.Value).Scan(&userID)
		if err == sql.ErrNoRows {
			// Clear the invalid cookie
			http.SetCookie(w, &http.Cookie{
				Name:     "auth_token",
				Value:    "",
				Path:     "/",
				MaxAge:   -1,
				HttpOnly: true,
				Secure:   r.TLS != nil,
				SameSite: http.SameSiteLaxMode,
			})

			// Handle invalid session differently based on request type
			isXHR := r.Header.Get("X-Requested-With") == "XMLHttpRequest"

			if IsAPIRoute(r.URL.Path) || isXHR {
				// For API routes or AJAX requests, return 401 Unauthorized
				logger.Warn("Invalid session token for API request, returning 401")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(`{"authenticated": false, "message": "Session expired"}`))
			} else {
				// For regular browser requests, handle within the SPA
				logger.Warn("Invalid session token, serving SPA")
				http.ServeFile(w, r, "./frontend/templates/index.html")
			}
			return
		} else if err != nil {
			logger.Error("Error checking session: %v", err)
			controller.ErrorHandler(w, "Something Unexpected Happened. Try Again Later", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
