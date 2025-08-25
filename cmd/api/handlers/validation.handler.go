package handlers

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
)

func (h *Handler) ValidateBodyRequest(payload interface{}) []*common.ApiValidationError {
	validate := validator.New(validator.WithRequiredStructEnabled())
	var validationError []*common.ApiValidationError

	err := validate.Struct(payload)
	if err != nil {
		var validationErrors validator.ValidationErrors
		ok := errors.As(err, &validationErrors)
		if ok {
			reflected := reflect.ValueOf(payload)
			if reflected.Kind() == reflect.Ptr {
				reflected = reflected.Elem()
			}

			for _, validationErr := range validationErrors {
				field, _ := reflected.Type().FieldByName(validationErr.StructField())
				key := field.Tag.Get("json")
				param := validationErr.Param()

				if key == "" {
					key = strings.ToLower(validationErr.StructField())
				}
				condition := validationErr.Tag()
				keyToTitleCase := strings.ReplaceAll(key, "_", " ")

				errMsg := fmt.Sprintf("%s field is %s", keyToTitleCase, condition)

				switch condition {
				case "required":
					errMsg = fmt.Sprintf("%s is required", keyToTitleCase)
				case "email":
					errMsg = fmt.Sprintf("%s must be a valid email address", keyToTitleCase)
				case "min":
					errMsg = fmt.Sprintf("%s must be at least %s characters", keyToTitleCase, param)
				case "max":
					errMsg = fmt.Sprintf("%s cannot be greater than %s characters", keyToTitleCase, param)
				case "eqfield":
					errMsg = fmt.Sprintf("%s must be equal to %s", keyToTitleCase, strings.ToLower(param))
				}

				validationError = append(validationError, &common.ApiValidationError{
					Error:     errMsg,
					Key:       key,
					Condition: condition,
				})
			}
		}
	}
	return validationError
}
