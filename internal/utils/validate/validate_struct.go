package validate

import (
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(request any) error {
	validateData := validator.New()
	return validateData.Struct(request)
}
