package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(val interface{}) error {
	err := validate.Struct(val)
	if err == nil {
		return nil
	}

	return err.(validator.ValidationErrors)
}
