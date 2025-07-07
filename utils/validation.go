package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error, message *string) gin.H {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationErrors {
			errors["path"] = e.Field()
			errors["code"] = "custom"
			errors["message"] = *message
		}

		return gin.H{"errors": errors}

	}
	return gin.H{"message": "yeu cau khong hop le"}
}
