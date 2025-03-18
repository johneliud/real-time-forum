package routes

import (
	"net/http"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/backend/middleware"
	"github.com/johneliud/real-time-forum/backend/wbsocket"
)

func Routes() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && !(len(r.URL.Path) >= 5 && r.URL.Path[:5] == "/api/") {
			if !fileExists(r.URL.Path) {
				http.ServeFile(w, r, "./frontend/templates/index.html")
				return
			}
		}
		http.ServeFile(w, r, "./frontend/templates/index.html")
	})

	mux.HandleFunc("/api/sign-up", controller.SignupHandler)
	mux.HandleFunc("/api/validate", controller.ValidateInputHandler)
	mux.HandleFunc("/api/sign-in", controller.SigninHandler)
	mux.HandleFunc("/api/logout", controller.LogoutHandler)

	mux.HandleFunc("/api/auth-status", controller.AuthStatusHandler)

	mux.HandleFunc("/api/messages", controller.GetMessagesHandler)

	mux.Handle("/api/user/profile", middleware.AuthenticateMiddleware(http.HandlerFunc(controller.GetUserProfileHandler)))

	// WebSocket route
	mux.HandleFunc("/ws", wbsocket.HandleWebSocket)

	handler := middleware.AuthenticateMiddleware(mux)
	http.Handle("/", handler)
}

// fileExists checks if the request is for a static file that exists.
func fileExists(path string) bool {
	cleanPath := path
	if len(path) > 0 && path[0] == '/' {
		cleanPath = path[1:]
	}

	// Check if file exists in frontend directory
	_, err := http.Dir("./frontend").Open(cleanPath)
	return err == nil
}
