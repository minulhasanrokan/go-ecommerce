package applicationData

import (
	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/handlers"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/middlewares"
)

type Application struct {
	Logger        echo.Logger
	Server        *echo.Echo
	Handler       handlers.Handler
	AppMiddleware middlewares.AppMiddleware
}
