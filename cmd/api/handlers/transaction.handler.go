package handlers

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/payment"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/services"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/helpers"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
)

func (h *Handler) MakePayment(c echo.Context) error {

	user, _ := c.Get("user").(models.UserModel)

	pubKey := os.Getenv("PUBLIC_KEY")

	service := services.NewAppService(h.Db)

	activePayment, err := service.GetActivePayment(user.Id)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	if activePayment.Id > 0 {

		return common.ApiSendSuccessResponse(c, "payment", map[string]string{
			"pubKey": pubKey,
			"secret": activePayment.ClientSecret,
		})
	}

	_, amount, err := service.FindAllCart(user.Id)

	orderId, err := helpers.GenerateCode()

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return nil
}

func (h *Handler) VerifyPayment(c echo.Context) error {

	return nil
}

func (h *Handler) GetOrderDetails(c echo.Context) error {

	return nil
}
