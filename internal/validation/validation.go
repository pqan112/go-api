package validation

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"pqan.com/go-api/internal/utils"
)

func InitValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	RegisterCustomValidation(v)
	return nil
}

func HandleValidationErrors(err error) utils.ValidationErrorResponse {
	var errors []utils.ErrorDetail

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var msg string
			log.Print(validationErrors)
			switch e.Tag() {
			case "required":
				msg = "error.required"
			case "email":
				msg = "error.invalid_email"
			case "userStatus":
				msg = "error.invalid_user_status"
			default:
				msg = "error.input_invalid"
			}

			errors = append(errors, utils.ErrorDetail{
				Path:    strings.ToLower(e.Field()),
				Message: msg,
			})
		}

	}

	if len(errors) == 0 {
		errors = append(errors, utils.ErrorDetail{
			Path:    "",
			Message: "error.invalid_body_format",
		})
	}

	return utils.ValidationErrorResponse{
		Message:    errors,
		Error:      "error.unprocessable_entity",
		StatusCode: http.StatusUnprocessableEntity,
	}

}

func RegisterCustomValidation(v *validator.Validate) {
	v.RegisterValidation("userStatus", func(fl validator.FieldLevel) bool {
		var validStatus = map[string]bool{
			"ACTIVE":   true,
			"INACTIVE": true,
			"BLOCKED":  true,
		}

		return validStatus[fl.Field().String()]
	})
}

func IsOneOf(value string, allowed ...string) bool {
	for _, v := range allowed {
		if value == v {
			return true
		}
	}
	return false
}
