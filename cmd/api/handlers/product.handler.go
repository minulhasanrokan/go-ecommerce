package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/requests"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/services"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
)

func (h *Handler) GetProducts(c echo.Context) error {

	service := services.NewAppService(h.Db)

	products, err := service.GetProducts()

	if err != nil {

		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "products", products)
}

func (h *Handler) GetProduct(c echo.Context) error {

	var productId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &productId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	service := services.NewAppService(h.Db)

	product, err := service.GetProductById(productId.Id)

	if err != nil {
		return common.ApiSendNotFoundResponse(c, err.Error())
	}
	return common.ApiSendSuccessResponse(c, "product", product)
}

func (h *Handler) GetCategories(c echo.Context) error {

	service := services.NewAppService(h.Db)

	categories, err := service.GetCategories()

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "categories", categories)
}

func (h *Handler) GetCategoryById(c echo.Context) error {

	var categoryId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &categoryId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	service := services.NewAppService(h.Db)

	category, err := service.GetCategory(categoryId.Id)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "category", category)
}

func (h *Handler) CreateCategories(c echo.Context) error {

	payload := new(requests.CreateCategoryRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	category, err := service.CreateCategory(payload)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "category", category)
}

func (h *Handler) EditCategory(c echo.Context) error {

	var categoryId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &categoryId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.CreateCategoryRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	updateCategory, err := service.EditCategory(categoryId.Id, payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "category", updateCategory)
}

func (h *Handler) DeleteCategory(c echo.Context) error {

	var categoryId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &categoryId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	service := services.NewAppService(h.Db)

	err = service.DeleteCategory(categoryId.Id)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "category deleted successfully", nil)
}

func (h *Handler) CreateProducts(c echo.Context) error {

	payload := new(requests.CreateProductRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	product, err := service.CreateProduct(payload, user)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "product", product)
}

func (h *Handler) EditProduct(c echo.Context) error {

	var productId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &productId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.CreateProductRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	product, err := service.EditProduct(productId.Id, payload, user)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "product", product)
}

func (h *Handler) UpdateStock(c echo.Context) error {

	var productId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &productId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	payload := new(requests.CreateProductRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	product := models.ProductModel{
		Stock:  uint(payload.Stock),
		UserId: int(user.Id),
	}

	updatedProduct, err := service.UpdateProductStock(product, productId.Id)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "product", updatedProduct)
}

func (h *Handler) DeleteProduct(c echo.Context) error {

	var productId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &productId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	err = service.DeleteProduct(productId.Id, user)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "product deleted successfully", nil)
}
