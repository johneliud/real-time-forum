package util

import (
	"log"
	"os"
)

/*
Retrieves the values of two environment variables, GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET, and returns them as strings.
*/
func LoadCredentials() (string, string) {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if len(googleClientID) == 0 || len(googleClientSecret) == 0 {
		log.Printf("Invalid credentials from .env")
		return "", ""
	}

	return googleClientID, googleClientSecret
}
