package internalErrors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}
	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]
	field := strings.ToLower(validationError.Field())
	switch validationError.Tag() {
	case "required":
		return errors.New(field + " is required")
	case "max":
		return errors.New(field + " must be less than " + validationError.Param())
	case "min":
		return errors.New(field + " must be greater than " + validationError.Param())
	case "email":
		return errors.New(field + " is not a valid email")
	}

	return nil
}
