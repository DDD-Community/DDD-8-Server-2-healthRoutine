package util

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"unicode"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(any interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(any)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

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
