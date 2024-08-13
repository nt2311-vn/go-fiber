package validation

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if !regexp.MustCompile(regex).MatchString(email) {
		return errors.New("invalid email address")
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func ValidatePasswordConfirm(password, confirm string) error {
	if confirm == "" {
		return errors.New("confirm password is required")
	}
	if password != confirm {
		return errors.New("passwords do not match")
	}
	return nil
}
