package auth

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
)

const REDIRECT_URL = "http://localhost:8080"

/*
Generates a random 16-byte state cookie, encodes it in base64, sets the cookie in the HTTP response, and returns the encoded state.
*/
func generateStateCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Printf("Error generating random state: %v", err)
		return ""
	}

	state := base64.URLEncoding.EncodeToString(b)

	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Path:     "/",
		Domain:   "",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
	})
	return state
}
