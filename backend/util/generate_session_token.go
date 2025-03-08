package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateSessionToken creates a secure random session token.
func GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("failed to read b: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
