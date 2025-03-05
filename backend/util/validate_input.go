package util

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/johneliud/real-time-forum/backend/model"
)

// ValidateSignupRequest checks for the validity of all signup input values
func ValidateSignupRequest(req model.SignupRequest) error {
	// Validate first name
	if len(strings.TrimSpace(req.FirstName)) < 2 {
		return fmt.Errorf("first name must be at least 2 characters long")
	}

	// Validate last name
	if len(strings.TrimSpace(req.LastName)) < 2 {
		return fmt.Errorf("last name must be at least 2 characters long")
	}

	// Validate nickname
	if len(strings.TrimSpace(req.NickName)) < 3 {
		return fmt.Errorf("nickname must be at least 3 characters long")
	}

	// Validate gender
	validGenders := map[string]bool{
		"male":   true,
		"female": true,
	}
	if !validGenders[strings.ToLower(req.Gender)] {
		return fmt.Errorf("invalid gender selection")
	}

	// Validate age
	if req.Age < 13 || req.Age > 120 {
		return fmt.Errorf("age must be between 13 and 120")
	}

	// Validate email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return fmt.Errorf("invalid email address format")
	}

	// Validate password
	if len(req.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	// Check password complexity
	if !regexp.MustCompile(`[A-Z]`).MatchString(req.Password) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(req.Password) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(req.Password) {
		return fmt.Errorf("password must contain at least one number")
	}
	if !regexp.MustCompile(`[!,.:;(){}?_@#$%^&*]`).MatchString(req.Password) {
		return fmt.Errorf("password must contain at least one special character")
	}

	// Validate password confirmation
	if req.Password != req.ConfirmedPassword {
		return fmt.Errorf("passwords do not match")
	}

	return nil
}
