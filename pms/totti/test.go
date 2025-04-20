package totti

import (
	"net/mail"
	"regexp"
	"strings"
	"fmt"
)

// IsValidEmail checks if the provided string is a valid email address.
func IsValidEmail(email string) bool {
	// 1. Basic Length Check (Optional but good for quick filtering)
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	// 2. Using the net/mail package for more robust parsing
	_, err := mail.ParseAddress(email)
	if err == nil {
		return true
	}

	// 3. (Optional but Recommended) Additional Regular Expression Validation
	// The net/mail package is good but might accept some technically valid
	// but unusual or potentially problematic formats. A well-crafted
	// regex can provide an extra layer of validation for common patterns.
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return true
	}

	return false
}

// EnforceValidEmail attempts to ensure an email is valid.
// It can either return the valid email or an error if invalid.
func EnforceValidEmail(email string) (string, error) {
	trimmedEmail := strings.TrimSpace(email)
	if !IsValidEmail(trimmedEmail) {
		return "", fmt.Errorf("invalid email format: %s", email)
	}
	return trimmedEmail, nil
}

// Example Usage:
func Use() {
	emails := []string{
		"test@example.com",
		"invalid-email",
		"another.test@sub.example.co.uk",
		"test@localhost",
		".test@example.com",
		"test.@example.com",
		"test..test@example.com",
		"test@example..com",
		"test@123.123.123.123",
	}

	for _, email := range emails {
		if IsValidEmail(email) {
			fmt.Printf("'%s' is a valid email.\n", email)
		} else {
			fmt.Printf("'%s' is NOT a valid email.\n", email)
		}

		validEmail, err := EnforceValidEmail(email)
		if err != nil {
			fmt.Println("  Error:", err)
		} else {
			fmt.Println("  Enforced:", validEmail)
		}
	}
}