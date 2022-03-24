package utils

import (
	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field string
	Error string
}

func Validate(form interface{}) []FieldError {

	var validate = validator.New()
	err := validate.Struct(form)
	var listError []FieldError
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errorMessage FieldError
			errorMessage.Field = err.Field()
			errorMessage.Error = err.Error()
			listError = append(listError, errorMessage)
		}
	}
	return listError
}
