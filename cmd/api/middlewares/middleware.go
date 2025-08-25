package middlewares

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppMiddleware struct {
	Logger echo.Logger
	Db     *gorm.DB
}
