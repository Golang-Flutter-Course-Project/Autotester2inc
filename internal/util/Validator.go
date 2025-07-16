package util

import (
	"Autotester/internal/domain"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

// ValidateUrl validates the given URL string pointer.
func ValidateUrl(payload *domain.UrlRequest) error {
	validate := validator.New()
	err := validate.Struct(*payload)
	if err != nil {
		log.Println("Validation error:", err)
		return errors.New("Validator error: " + err.Error())
	}
	return nil
}
