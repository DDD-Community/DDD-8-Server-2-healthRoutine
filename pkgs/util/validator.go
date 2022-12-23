package util

import (
	"regexp"
	"unicode"
)

func CheckEmail(email string) (result bool) {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	result = emailRegex.MatchString(email)
	//_, err := mail.ParseAddress(email)
	//return err == nil
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
