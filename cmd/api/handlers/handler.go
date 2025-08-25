package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/payment"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/mailer"
	"gorm.io/gorm"
)

type Handler struct {
	Db     *gorm.DB
	Logger echo.Logger
	Mailer mailer.Mailer
}

func (h *Handler) BindBodyRequest(c echo.Context, payload interface{}) error {

	if err := (&echo.DefaultBinder{}).BindBody(c, &payload); err != nil {

		return errors.New("failed to bind body, make sure you are sending valid payload")
	}

	return nil
}
