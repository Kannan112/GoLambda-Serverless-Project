package validater

import "regexp"

func IsEmailValid(email string) bool {
	// Regular expression pattern for email validation
	rxEmail := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

	// Check if the email matches the pattern
	if len(email) < 3 || len(email) > 254 || rxEmail.MatchString(email) {
		return false
	}
	return true
}
