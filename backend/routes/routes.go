package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/johneliud/real-time-forum/backend/controller"
	"github.com/johneliud/real-time-forum/backend/middleware"
	"github.com/johneliud/real-time-forum/backend/util"
	"github.com/johneliud/real-time-forum/backend/wbsocket"
)

// Routes sets up and returns an http.ServeMux that maps URLs to their corresponding handlers.
func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./frontend"))
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	mux.HandleFunc("/", serveFrontend)
	mux.HandleFunc("/api/sign-up", controller.SignupHandler)
	mux.HandleFunc("/api/sign-in", controller.SigninHandler)
	mux.HandleFunc("/api/validate", controller.ValidateInputHandler)
	mux.HandleFunc("/api/logout", controller.LogoutHandler)
	mux.HandleFunc("/api/auth-status", controller.AuthStatusHandler)
	mux.HandleFunc("/api/messages", util.GetMessagesHandler)

	mux.Handle("/api/profile", middleware.AuthenticateMiddleware(http.HandlerFunc(util.GetUserProfileHandler)))

	mux.HandleFunc("/ws", wbsocket.HandleWebSocket)
	return mux
}

// serveFrontend serves the frontend.
func serveFrontend(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && !(strings.Contains(r.URL.Path, "/api/")) {
		if !fileExists(r.URL.Path) {
			http.ServeFile(w, r, "./frontend/templates/index.html")
			return
		}
	}
	http.ServeFile(w, r, "./frontend/templates/index.html")
}

// fileExists checks if the request is for a static file that exists.
func fileExists(path string) bool {
	cleanPath := strings.TrimPrefix(path, "/")

	_, err := os.Stat("./frontend/" + cleanPath)
	return err == nil
}
