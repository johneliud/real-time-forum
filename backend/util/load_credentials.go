package util

import (
	"os"
)

/*
Retrieves the values of two environment variables, GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET, and returns them as strings.
*/
func LoadCredentials() (string, string) {
	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	return googleClientID, googleClientSecret
}
