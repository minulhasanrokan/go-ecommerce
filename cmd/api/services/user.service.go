package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/minulhasanrokan/go-ecommerce/cmd/api/requests"
	"github.com/minulhasanrokan/go-ecommerce/cmd/common"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/helpers"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
	"gorm.io/gorm/clause"
)

func (service *AppService) RegisterUser(userRequest *requests.RegisterUserRequest) (*models.UserModel, error) {

	hashedPassword, err := helpers.HashPassword(userRequest.Password)

	if err != nil {

		fmt.Println(err)
		return nil, errors.New("registration failed")
	}

	createdUser := models.UserModel{
		Email:     userRequest.Email,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Password:  hashedPassword,
	}

	result := service.Db.Create(&createdUser)

	if result.Error != nil {

		fmt.Println(result.Error)

		return nil, errors.New("registration failed")
	}

	return &createdUser, nil
}

func (service *AppService) IsVerifiedUser(id uint) bool {

	var user models.UserModel

	result := service.Db.Where("id = ?", id).First(&user)

	if result.Error != nil {

		return false
	}

	return user.Verified
}

func (service *AppService) GetVerificationCode(model *models.UserModel) (int, error) {

	if service.IsVerifiedUser(model.Id) {

		return 0, errors.New("user already verified")
	}

	code, err := helpers.GenerateCode()

	if err != nil {
		return 0, err
	}
	codeData, err := strconv.Atoi(code)

	if err != nil {
		return 0, err
	}

	user := models.UserModel{
		ExpiredAt: time.Now().Add(30 * time.Minute),
		Code:      codeData,
	}

	_, err = service.UpdateUser(model.Id, user)

	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	return codeData, nil
}

func (service *AppService) UpdateUser(id uint, data models.UserModel) (*models.UserModel, error) {

	var user models.UserModel

	err := service.Db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(data).Error

	if err != nil {
		return &models.UserModel{}, errors.New("failed update user")
	}

	return &user, nil
}

func (service *AppService) VerifyCode(userId uint, payload *requests.VerificationCodeRequest) error {

	if service.IsVerifiedUser(userId) {
		return errors.New("user already verified")
	}

	userData, err := service.GetUserById(userId)

	if err != nil {
		return err
	}

	code, err := strconv.Atoi(payload.Code)

	if err != nil {
		return err
	}

	if userData.Code != code {

		return errors.New("code does not match")
	}

	if !time.Now().Before(userData.ExpiredAt) {

		return errors.New("verification code expired")
	}

	updatedUser := models.UserModel{
		Verified: true,
	}

	_, err = service.UpdateUser(userId, updatedUser)

	if err != nil {
		return err
	}

	return nil
}

func (service *AppService) CreateUserProfile(userId uint, payload *requests.ProfileInputRequest) error {

	_, err := service.GetUserById(userId)

	if err != nil {
		return err
	}

	user := models.UserModel{}

	if payload.FirstName != "" {

		user.FirstName = payload.FirstName
	}

	if payload.LastName != "" {
		user.LastName = payload.LastName
	}

	_, err = service.UpdateUser(userId, user)

	if err != nil {
		return err
	}

	address := models.AddressModel{
		AddressLine1: payload.AddressInput.AddressLine1,
		AddressLine2: payload.AddressInput.AddressLine2,
		City:         payload.AddressInput.City,
		Country:      payload.AddressInput.Country,
		PostCode:     payload.AddressInput.PostCode,
		UserId:       userId,
	}

	err = service.Db.Create(&address).Error

	if err != nil {

		return err
	}

	return nil
}

func (service *AppService) GetUserProfile(userId uint) (*models.UserModel, error) {

	userData, err := service.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (service *AppService) UpdateUserProfile(userId uint, payload *requests.ProfileUpdateRequest) (*models.UserModel, error) {

	userData, err := service.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	user := models.UserModel{}

	if payload.FirstName != "" {

		user.FirstName = payload.FirstName
	}

	if payload.LastName != "" {
		user.LastName = payload.LastName
	}

	_, err = service.UpdateUser(userId, user)

	if err != nil {

		return nil, err
	}

	var address models.AddressModel

	err = service.Db.Where("user_id=?", userId).First(&address, payload.AddressInput.AddressId).Error

	if err != nil {

		return nil, err
	}

	address.AddressLine1 = payload.AddressInput.AddressLine1
	address.AddressLine2 = payload.AddressInput.AddressLine2
	address.City = payload.AddressInput.City
	address.PostCode = payload.AddressInput.PostCode
	address.Country = payload.AddressInput.Country
	address.UserId = userId

	err = service.Db.Save(&address).Error

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (service *AppService) CreateCart(user *models.UserModel, payload *requests.CreateCartRequest) ([]models.CartModel, error) {

	cart, _ := service.FindCart(user.Id, payload.ProductId)

	if cart.Id > 0 {

		if payload.ProductId == 0 {

			return nil, errors.New("please provide a product")
		}

		if payload.Qty < 0 {

			err := service.DeleteCartById(cart.Id)

			if err != nil {
				return nil, err
			}
		} else {
			cart.Qty = payload.Qty

			err := service.UpdateCart(*cart)

			if err != nil {
				return nil, err
			}
		}
	} else {

		product, err := service.FindProductById(payload.ProductId)

		if err != nil {
			return nil, err
		}

		if product.Id < 1 {

			return nil, errors.New("product not found")
		}

		err = service.CreateCartData(models.CartModel{
			UserId:    user.Id,
			ProductId: product.Id,
			Name:      product.Name,
			ImageUrl:  product.ImageUrl,
			Qty:       payload.Qty,
			Price:     product.Price,
			SellerId:  uint(product.UserId),
		})

		if err != nil {
			return nil, err
		}
	}

	cartItems, err := service.FindCartItems(user.Id)

	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (service *AppService) FindCartItems(uId uint) ([]models.CartModel, error) {

	var carts []models.CartModel

	err := service.Db.Where("user_id=?", uId).Find(&carts).Error

	return carts, err
}

func (service *AppService) CreateCartData(c models.CartModel) error {

	return service.Db.Create(&c).Error
}

func (service *AppService) UpdateCart(c models.CartModel) error {

	var cart models.CartModel

	err := service.Db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.Id).Updates(c).Error

	return err
}

func (service *AppService) DeleteCartById(id uint) error {

	err := service.Db.Delete(&models.CartModel{}, id).Error

	return err
}

func (service *AppService) FindCart(userId uint, productId uint) (*models.CartModel, error) {

	var cartItem models.CartModel

	err := service.Db.Where("user_id=? AND product_id=?", userId, productId).First(&cartItem).Error

	return &cartItem, err
}

func (service *AppService) FindAllCart(userId uint) ([]models.CartModel, float64, error) {

	cartItems, err := service.FindCartItems(userId)

	if err != nil {
		return nil, 0, err
	}

	var totalAmounts float64

	for _, cartItem := range cartItems {

		totalAmounts += cartItem.Price * float64(cartItem.Qty)
	}

	return cartItems, totalAmounts, nil
}

func (service *AppService) GetOrders(model *models.UserModel) ([]models.OrderModel, error) {

	orders, err := service.FindOrders(model.Id)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (service *AppService) FindOrders(uId uint) ([]models.OrderModel, error) {

	var orders []models.OrderModel

	err := service.Db.Where("user_id=?", uId).Find(&orders).Error

	if err != nil {
		return nil, errors.New("failed to fetch orders")
	}

	return orders, nil
}

func (service *AppService) GetOrderById(userId uint, orderId uint) (models.OrderModel, error) {

	order, err := service.FindOrderById(userId, orderId)

	if err != nil {

		return order, err
	}

	return order, nil
}

func (service *AppService) FindOrderById(id uint, uId uint) (models.OrderModel, error) {

	var order models.OrderModel

	err := service.Db.Preload("Items").Where("id=? AND user_id=?", id, uId).First(&order).Error

	if err != nil {

		return order, errors.New("failed to fetch order")
	}

	return order, nil
}

func (service *AppService) BecomeSeller(userId uint, payload *requests.SellerInputRequest) (string, error) {

	user, err := service.GetUserById(userId)
	if err != nil {
		return "", err
	}

	if user.UserType == models.SELLER {

		return "", errors.New("you have already joined seller program")
	}

	seller, err := service.UpdateUser(user.Id, models.UserModel{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Mobile:    payload.PhoneNumber,
		UserType:  models.SELLER,
	})

	if err != nil {
		return "", err
	}

	token, _, err := common.GenerateJWT(*seller)

	if err != nil {
		return "", err
	}

	err = service.CreateBankAccount(models.BankAccountModel{
		BankAccount: payload.BankAccountNumber,
		SwiftCode:   payload.SwiftCode,
		PaymentType: payload.PaymentType,
		UserId:      user.Id,
	})

	if err != nil {
		return "", err
	}

	return *token, nil
}

func (service *AppService) CreateBankAccount(model models.BankAccountModel) error {

	return service.Db.Create(&model).Error
}

func (service *AppService) GetUserByEmail(email string) (*models.UserModel, error) {

	var user models.UserModel

	result := service.Db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (service *AppService) GetUserById(userId uint) (*models.UserModel, error) {

	var user models.UserModel

	result := service.Db.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
