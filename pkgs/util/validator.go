package util

import (
	"regexp"
	"unicode"
)

func CheckEmailRegex(email string) (result bool) {
	result, _ = regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return
}

func CheckPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower := false
	for _, c := range password {
		if unicode.IsLower(c) {
			hasLower = true
			break
		}
	}
	if !hasLower {
		return false
	}

	hasNumber := false
	for _, c := range password {
		if unicode.IsNumber(c) {
			hasNumber = true
			break
		}
	}
	if !hasNumber {
		return false
	}

	hasSpecial := false
	for _, c := range password {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			hasSpecial = true
			break
		}
	}
	if !hasSpecial {
		return false
	}
	
	return true
}

// Deprecated
func CheckPasswordRegex(password string) (result bool) {
	//TODO: fix regex
	result, _ = regexp.MatchString("^(([a-zA-Z])|([0-9])|([.!@#$%^&*()-+/=?^_{|}~-])){8,}$", password)
	return
}
