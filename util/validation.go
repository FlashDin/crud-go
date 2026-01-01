package util

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				return e.Field() + " is required"
			case "min":
				return e.Field() + " is too small"
			}
		}
	}
	return errors.New("invalid request").Error()
}
