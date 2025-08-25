package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	Health bool `json:"health"`
}

func (h *Handler) HealthCheck(c echo.Context) error {

	data := HealthResponse{
		Health: true,
	}

	return c.JSON(http.StatusOK, data)
}
