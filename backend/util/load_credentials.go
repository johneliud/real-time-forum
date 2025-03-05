package util

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split key-value pair
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set environment variable
		os.Setenv(key, value)
	}
	LoadCredentials()
	return scanner.Err()
}
