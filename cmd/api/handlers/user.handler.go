package handlers

import (
	"errors"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/requests"
	"github.com/minulhasanrokan/go-ecommerce/cmd/api/services"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/helpers"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/mailer"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
	"gorm.io/gorm"
)

func (h *Handler) UserRegister(c echo.Context) error {

	payload := new(requests.RegisterUserRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	_, err := service.GetUserByEmail(payload.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) == false {

		return common.ApiSendBadRequestResponse(c, "E-mail already has been registered.")
	}

	registeredUser, err := service.RegisterUser(payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	mailData := mailer.EmailData{
		Subject: "Welcome To " + os.Getenv("APP_NAME"),
		Meta: struct {
			FirstName string
			LoginLink string
		}{
			FirstName: registeredUser.FirstName,
			LoginLink: "#",
		},
	}

	err = h.Mailer.Send(payload.Email, "welcome.html", mailData)

	token, refreshToken, err := common.GenerateJWT(*registeredUser)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "User registration successful", map[string]interface{}{
		"user":          registeredUser,
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) UserLogin(c echo.Context) error {

	payload := new(requests.LoginUserRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	userData, err := service.GetUserByEmail(payload.Email)

	if errors.Is(err, gorm.ErrRecordNotFound) || helpers.CheckPasswordHash(payload.Password, userData.Password) == false {
		return common.ApiSendBadRequestResponse(c, "Invalid email address or password.")
	}

	accessToken, refreshToken, err := common.GenerateJWT(*userData)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "User logged in", map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user":          userData,
	})
}

func (h *Handler) GetUserVerificationCode(c echo.Context) error {

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	userData, err := service.GetUserById(user.Id)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, "unable to generate verification code")
	}

	code, err := service.GetVerificationCode(userData)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	mailData := mailer.EmailData{
		Subject: "Verify Code",
		Meta: struct {
			FirstName string
			Code      int
		}{
			FirstName: userData.FirstName,
			Code:      code,
		},
	}

	err = h.Mailer.Send(userData.Email, "VerifyCode.html", mailData)

	token, refreshToken, err := common.GenerateJWT(*userData)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "send verification cade successfully", map[string]interface{}{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) VerifyUser(c echo.Context) error {

	payload := new(requests.VerificationCodeRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	err := service.VerifyCode(user.Id, payload)

	if err != nil {

		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "verified successfully", nil)
}

func (h *Handler) CreateUserProfile(c echo.Context) error {

	payload := new(requests.ProfileInputRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	err := service.CreateUserProfile(user.Id, payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "user profile created successfully", nil)
}

func (h *Handler) GetUserProfile(c echo.Context) error {

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	userProfile, err := service.GetUserProfile(user.Id)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "user profile retrieved successfully", userProfile)
}

func (h *Handler) UpdateUserProfile(c echo.Context) error {

	payload := new(requests.ProfileUpdateRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	userProfile, err := service.UpdateUserProfile(user.Id, payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "user profile updated successfully", userProfile)
}

func (h *Handler) GetCartItem(c echo.Context) error {

	payload := new(requests.CreateCartRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)

	user, _ := c.Get("user").(models.UserModel)

	cartItems, err := service.CreateCart(&user, payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "cartItems created successfully", cartItems)
}

func (h *Handler) AddToCart(c echo.Context) error {

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	cart, _, er := service.FindAllCart(user.Id)

	if er != nil {

		return common.ApiSendInternalServerErrorResponse(c, er.Error())
	}

	return common.ApiSendSuccessResponse(c, "get cart", cart)
}

func (h *Handler) GetOrders(c echo.Context) error {

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	orders, err := service.GetOrders(&user)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "get orders", orders)
}

func (h *Handler) GetSingleOrder(c echo.Context) error {

	var orderId requests.IdParamRequest

	err := (&echo.DefaultBinder{}).BindPathParams(c, &orderId)

	if err != nil {
		return common.ApiSendBadRequestResponse(c, err.Error())
	}

	user, _ := c.Get("user").(models.UserModel)

	service := services.NewAppService(h.Db)

	order, err := service.GetOrderById(user.Id, orderId.Id)

	if err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "get order successfully", order)
}

func (h *Handler) BecomeSeller(c echo.Context) error {

	payload := new(requests.SellerInputRequest)

	if err := h.BindBodyRequest(c, payload); err != nil {

		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	validationErrors := h.ValidateBodyRequest(*payload)

	if validationErrors != nil {

		return common.ApiSendFailedValidationResponse(c, validationErrors)
	}

	service := services.NewAppService(h.Db)
	user, _ := c.Get("user").(models.UserModel)

	token, err := service.BecomeSeller(user.Id, payload)

	if err != nil {
		return common.ApiSendInternalServerErrorResponse(c, err.Error())
	}

	return common.ApiSendSuccessResponse(c, "become seller successfully", token)
}
