package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FieldErrors(err error) string {
	var errorMsg string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, ve := range validationErrors {
			fieldName := ve.Field()
			tagName := ve.Tag()
			paramValue := ve.Param()

			if tagName == "required" {
				errorMsg = fmt.Sprintf("%s is required", fieldName)
				break
			}

			if tagName == "min" {
				errorMsg = fmt.Sprintf("%s must be at least %s characters", fieldName, paramValue)
				break
			}

			if tagName == "max" {
				errorMsg = fmt.Sprintf("%s must not exceed %s characters", fieldName, paramValue)
				break
			}

			if tagName == "numeric" {
				errorMsg = fmt.Sprintf("%s must be numeric", fieldName)
				break
			}

			if tagName == "alpha" {
				errorMsg = fmt.Sprintf("%s must contain only alphabetic characters", fieldName)
				break
			}

			if tagName == "alphanum" {
				errorMsg = fmt.Sprintf("%s must be alphanumeric", fieldName)
				break
			}

			if tagName == "len" {
				errorMsg = fmt.Sprintf("%s must be exactly %s characters", fieldName, paramValue)
				break
			}

			errorMsg = fmt.Sprintf("%s is invalid", fieldName)
		}
	}

	return errorMsg
}
