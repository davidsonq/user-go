package validations

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetCustomErrorMessageUser(errs validator.ValidationErrors) string {
	for _, err := range errs {
		switch err.Tag() {
		case "email":
			return fmt.Sprintf("'%s' isn't invalid.", err.Field())
		case "required":
			return fmt.Sprintf("The field '%s' is required.", err.Field())

		case "min":
			if err.Field() == "Password" {
				return fmt.Sprintf("The field '%s' must have at least 6 characters.", err.Field())
			}
			return fmt.Sprintf("The field '%s' must have at least 3 characters.", err.Field())
		case "max":
			if err.Field() == "Password" {
				return fmt.Sprintf("The field '%s' must have at most 16 characters.", err.Field())
			}
			return fmt.Sprintf("The field '%s' must have at most 50 characters.", err.Field())
		}
	}
	return "Invalid input data."
}

func DuplicateErrorUser(err error) map[string]string {

	switch {
	case strings.Contains(err.Error(), "email"):
		return map[string]string{"email": "email already exists!"}
	case strings.Contains(err.Error(), "nickname"):
		return map[string]string{"nickname": "nickname already exists"}
	}
	return map[string]string{"error": err.Error()}
}
