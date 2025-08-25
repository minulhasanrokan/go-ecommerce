package common

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse map[string]interface{}

type ApiJasonSuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ApiValidationError struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

type ApiJsonFailedValidationResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Errors  []*ApiValidationError `json:"errors"`
}

type ApiJsonErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ApiSendSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, ApiJasonSuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ApiSendFailedValidationResponse(c echo.Context, errors []*ApiValidationError) error {
	return c.JSON(http.StatusUnprocessableEntity, ApiJsonFailedValidationResponse{
		Success: false,
		Errors:  errors,
		Message: "Validation Failed",
	})
}

func ApiSendErrorResponse(c echo.Context, message string, statusCode int) error {
	return c.JSON(statusCode, ApiJsonErrorResponse{
		Success: false,
		Message: message,
	})
}

func ApiSendBadRequestResponse(c echo.Context, message string) error {
	return ApiSendErrorResponse(c, message, http.StatusBadRequest)
}

func ApiSendNotFoundResponse(c echo.Context, message string) error {
	return ApiSendErrorResponse(c, message, http.StatusNotFound)
}

func ApiSendInternalServerErrorResponse(c echo.Context, message string) error {
	return ApiSendErrorResponse(c, message, http.StatusInternalServerError)
}

func ApiSendUnauthorizedResponse(c echo.Context, message string) error {

	if message == "" {
		message = "Unauthorized"
	}
	return ApiSendErrorResponse(c, message, http.StatusUnauthorized)
}
