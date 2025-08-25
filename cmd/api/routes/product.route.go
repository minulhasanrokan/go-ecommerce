package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/applicationData"
)

func ProductRoutes(api *echo.Group, data *applicationData.Application) {

	prefix := api.Group("/products")

	//public endpoint
	prefix.GET("/products", data.Handler.GetProducts)
	prefix.GET("/products/:productId", data.Handler.GetProduct)

	prefix.GET("/categories", data.Handler.GetCategories)
	prefix.GET("/categories/:categoryId", data.Handler.GetCategoryById)

	//private endpoint
	privatePrefix := api.Group("/seller/products", data.AppMiddleware.AuthenticationMiddleware)

	privatePrefix.POST("/categories", data.Handler.CreateCategories)
	privatePrefix.PATCH("/categories/:id", data.Handler.EditCategory)
	privatePrefix.DELETE("/categories/:id", data.Handler.DeleteCategory)

	privatePrefix.POST("/products", data.Handler.CreateProducts)
	privatePrefix.GET("/products", data.Handler.GetProducts)
	privatePrefix.GET("/products/:id", data.Handler.GetProduct)
	privatePrefix.PUT("/products/:id", data.Handler.EditProduct)
	privatePrefix.PATCH("/products/:id", data.Handler.UpdateStock) // update stock
	privatePrefix.DELETE("/products/:id", data.Handler.DeleteProduct)
}
