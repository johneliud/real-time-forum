package routes

import (
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/backend/middleware"
)

func Routes() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	// All routes will serve the same index.html for SPA
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && !(len(r.URL.Path) >= 5 && r.URL.Path[:5] == "/api/") {
			// Check if requesting a file - if not, serve index.html
			if !fileExists(r.URL.Path) {
				http.ServeFile(w, r, "./frontend/templates/index.html")
				return
			}
		}
		http.ServeFile(w, r, "./frontend/templates/index.html")
	})

	// API routes for authentication
	mux.HandleFunc("/api/sign-up", controller.SignupHandler)
	mux.HandleFunc("/api/validate", controller.ValidateInputHandler)
	mux.HandleFunc("/api/sign-in", controller.SigninHandler)
	mux.HandleFunc("/api/logout", controller.LogoutHandler)

	// API endpoint to check authentication status
	mux.HandleFunc("/api/auth-status", controller.AuthStatusHandler)

	// Apply the global authentication middleware
	handler := middleware.AuthenticateMiddleware(mux)
	http.Handle("/", handler)
}

// fileExists checks if the request is for a static file that exists
func fileExists(path string) bool {
	// Remove leading slash
	cleanPath := path
	if len(path) > 0 && path[0] == '/' {
		cleanPath = path[1:]
	}
	
	// Check if file exists in frontend directory
	_, err := http.Dir("./frontend").Open(cleanPath)
	return err == nil
}
