package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/applicationData"
)

func UserRoutes(api *echo.Group, data *applicationData.Application) {

	prefix := api.Group("/users")

	//public endpoint
	prefix.POST("/register", data.Handler.UserRegister)
	prefix.POST("/login", data.Handler.UserLogin)

	//private endpoint

	privatePrefix := api.Group("/users", data.AppMiddleware.AuthenticationMiddleware)

	privatePrefix.GET("/verify", data.Handler.GetUserVerificationCode)
	privatePrefix.POST("/verify", data.Handler.VerifyUser)

	privatePrefix.GET("/profile", data.Handler.GetUserProfile)
	privatePrefix.POST("/profile", data.Handler.CreateUserProfile)
	privatePrefix.PATCH("/profile", data.Handler.UpdateUserProfile)

	privatePrefix.POST("/cart", data.Handler.AddToCart)
	privatePrefix.GET("/cart", data.Handler.GetCartItem)

	privatePrefix.GET("/order", data.Handler.GetOrders)
	privatePrefix.GET("/order/:orderId", data.Handler.GetSingleOrder)

	privatePrefix.POST("/become-seller", data.Handler.BecomeSeller)
}
