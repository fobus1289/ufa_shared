package validator

import (
	"github.com/go-playground/validator/v10"
)

type ValidationErrors = validator.ValidationErrors

var validate = validator.New()

func Validate(val any) ValidationErrors {

	validate.RegisterValidation("iso639_2_alpha3", Iso6392Alpha3)

	if err := validate.Struct(val); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		return validationErrors
	}

	return nil
}
