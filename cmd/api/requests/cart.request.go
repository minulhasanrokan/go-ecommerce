package requests

type CreateCartRequest struct {
	ProductId uint `json:"product_id" validate:"required"`
	Qty       uint `json:"qty" validate:"required,min=1"`
}

type CreatePaymentRequest struct {
	OrderId      string  `json:"order_id" validate:"required"`
	PaymentId    string  `json:"payment_id" validate:"required"`
	ClientSecret string  `json:"client" validate:"required"`
	Amount       float64 `json:"amount" validate:"required,min=1"`
	UserId       uint    `json:"user_id" validate:"required,min=1"`
}
