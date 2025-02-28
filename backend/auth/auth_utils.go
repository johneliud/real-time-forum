package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
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
		log.Printf("Error generating random state: %v\n", err)
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

/*
Checks if the state query parameter in an HTTP request matches the value of the oauth_state cookie.
*/
func validateState(r *http.Request) error {
	state := r.URL.Query().Get("state")
	cookie, err := r.Cookie("oauth_state")
	if err != nil {
		log.Printf("Cookie error: %v\n", err)
		return err
	}

	if cookie.Value != state {
		log.Printf("State mismatch. Cookie: %s, State: %s\n", cookie.Value, state)
		return errors.New("invalid state")
	}
	return nil
}
