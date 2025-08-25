package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/applicationData"
)

func TransactionRoutes(api *echo.Group, data *applicationData.Application) {

	prefix := api.Group("/buyer", data.AppMiddleware.AuthenticationMiddleware)

	prefix.GET("/payment", data.Handler.MakePayment)
	prefix.GET("/verify", data.Handler.VerifyPayment)

	sellerPrefix := api.Group("/seller", data.AppMiddleware.AuthenticationMiddleware)

	sellerPrefix.GET("/orders", data.Handler.GetOrders)
	sellerPrefix.GET("/orders/:id", data.Handler.GetOrderDetails)
}
