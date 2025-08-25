package services

import "github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"

func (service *AppService) GetActivePayment(uId uint) (*models.PaymentModel, error) {

	return service.FindInitialPayment(uId)
}

func (service *AppService) FindInitialPayment(uId uint) (*models.PaymentModel, error) {

	var payment *models.PaymentModel

	err := service.Db.First(&payment, "user_id=? AND status=?", uId, "initial").Order("created_at desc").Error

	return payment, err
}
