package validations

import (
	"fmt"
	"strings"
	"user-go/errors"
	"user-go/models"

	"github.com/go-playground/validator/v10"
)

func ValidateUser(u *models.User) map[string]string {
	err := validator.New().Struct(u)

	if err != nil {
		validationErrors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[fmt.Sprint(err.Field())] = err.Error()
		}
		return validationErrors
	}
	return nil
}

func DuplicateError(err error) *errors.AppError {
	switch {
	case strings.Contains(err.Error(), "email"):
		return errors.NewAppError(map[string]string{"email": "email already exists!"})
	case strings.Contains(err.Error(), "nickname"):
		return errors.NewAppError(map[string]string{"nickname": "nickname already exists"})
	}
	return errors.NewAppError(map[string]string{"Error": err.Error()})
}
