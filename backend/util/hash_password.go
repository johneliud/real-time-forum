package util

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword uses bcrypt library to hash sensitive user data before storage.
*/
func HashPassword(password string) (string, error) {
	if len(strings.TrimSpace(password)) == 0 {
		return "", fmt.Errorf("password must contain atleast 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed generating password hash: %v", err)
	}
	return string(hashedPassword), nil
}
