package middleware

import (
	"fmt"
	"mini-wallet-exercise/constant"
	"mini-wallet-exercise/pkg"
	"mini-wallet-exercise/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequestJSON[T any]() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		obj := new(T)

		if err := httpContext.ShouldBind(obj); err != nil {
			if validationErrors, isNotValid := err.(validator.ValidationErrors); isNotValid {
				// Convert validation errors to ErrorValidation struct
				errors := make([]pkg.ErrorValidation, 0)

				for _, validationErr := range validationErrors {
					errors = append(errors, pkg.ErrorValidation{
						Key:     utils.StringToSnakeCase(validationErr.Field()), // Fully qualified field name
						Message: fmt.Sprintf("Error %s", validationErr.Error()),
					})
				}

				// Print errors as JSON
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(errors))
			} else {
				errors := make([]pkg.ErrorValidation, 0)
				errors = append(errors, pkg.ErrorValidation{
					Key:     "error",
					Message: err.Error(),
				})
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(errors))
			}

			httpContext.Abort()
			return
		}

		httpContext.Set(constant.RequestBodyJSONKey, obj)
	}
}

func ValidateRequestFormData[T any]() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		obj := new(T)

		if err := httpContext.ShouldBind(obj); err != nil {
			if validationErrors, isNotValid := err.(validator.ValidationErrors); isNotValid {
				errors := make([]pkg.ErrorValidation, len(validationErrors))

				for i, validationErr := range validationErrors {
					errors[i] = pkg.ErrorValidation{
						Key:     utils.StringToSnakeCase(validationErr.Field()),
						Message: fmt.Sprintf("Error %s", validationErr.Error()),
					}
				}

				fmt.Println("error validation", errors)

				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(errors))
			} else {
				fmt.Println("error else", err)
				httpContext.JSON(http.StatusBadRequest, utils.ErrorValidationResponse(err.Error()))
			}

			httpContext.Abort()
			return
		}

		httpContext.Set(constant.RequestBodyJSONKey, obj)
		httpContext.Next()
	}
}
