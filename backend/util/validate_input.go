package util

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

/*
ValidateInput checks for the validity of input values provided via the form.
*/
func ValidateInput(name, email, password, confirmedPassword string) error {
	if len(strings.TrimSpace(name)) == 0 || len(strings.TrimSpace(email)) == 0 || len(strings.TrimSpace(password)) == 0 || len(strings.TrimSpace(confirmedPassword)) == 0 {
		return fmt.Errorf("required fields cannot be empty")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email address format")
	}

	if len(strings.TrimSpace(password)) < 6 {
		return fmt.Errorf("password must contain atleast 6 characters")
	}

	if !reflect.DeepEqual(password, confirmedPassword) {
		return fmt.Errorf("passwords do not match")
	}
	return nil
}
