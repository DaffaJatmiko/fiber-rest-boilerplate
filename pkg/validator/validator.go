package validator

import (
	"github.com/DaffaJatmiko/fiber-rest-boilerplate/pkg/response"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateStruct memvalidasi struct dan return validation errors
func ValidateStruct(s interface{}) []response.ValidationError {
	var errors []response.ValidationError

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ValidationError
			element.Field = err.Field()
			element.Message = getErrorMessage(err)
			errors = append(errors, element)
		}
	}

	return errors
}

// getErrorMessage convert validation tag ke user-friendly message
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Minimum length is " + err.Param()
	case "max":
		return "Maximum length is " + err.Param()
	case "oneof":
		return "Must be one of: " + err.Param()
	default:
		return "Invalid value"
	}
}
