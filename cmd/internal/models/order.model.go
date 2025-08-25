package models

type OrderModel struct {
	BaseModel
	UserId         uint             `json:"user_id" gorm:"column:user_id"`
	Status         string           `json:"status" gorm:"column:status;varchar(250)"`
	Amount         float64          `json:"amount" gorm:"column:amount"`
	TransactionId  string           `json:"transaction_id" gorm:"column:transaction_id;varchar(250)"`
	OrderRefNumber string           `json:"order_ref_number" gorm:"column:order_ref_number;varchar(250)"`
	PaymentId      string           `json:"payment_id" gorm:"column:payment_id;varchar(250)"`
	Items          []OrderItemModel `json:"items" gorm:"foreignkey:OrderId"`
}

func (u OrderModel) TableName() string {
	return "orders"
}
