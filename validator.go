package validator

import "regexp"

func PasswordIsValid(password string) bool {
	if len(password) < 9 {
		return false
	}

	containsDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	containsLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	containsUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	containsSpecialChars := regexp.MustCompile(`[!@#$%^&*()\-+]`).MatchString(password)

	if !containsDigit || !containsLower || !containsUpper || !containsSpecialChars {
		return false
	}
	seen := make(map[rune]bool)
	for _, char := range password {
		if _, ok := seen[char]; ok {
			return false
		}
		seen[char] = true
	}

	return true
}
