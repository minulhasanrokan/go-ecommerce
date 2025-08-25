package models

type PaymentModel struct {
	BaseModel
	UserId        uint          `json:"user_id" gorm:"column:user_id"`
	CaptureMethod string        `json:"capture_method" gorm:"column:capture_method;type:varchar(250)"`
	Amount        float64       `json:"amount" gorm:"column:amount"`
	OrderId       string        `json:"order_id" gorm:"column:order_id;varchar(250)"`
	CustomerId    string        `json:"customer_id" gorm:"column:customer_id;varchar(250)"`
	PaymentId     string        `json:"payment_id" gorm:"column:payment_id;varchar(250)"`
	ClientSecret  string        `json:"client_secret" gorm:"column:client_secret;varchar(250)"`
	Status        PaymentStatus `json:"status"`
	Response      string        `json:"response" gorm:"column:response;varchar(250)"`
}

func (u PaymentModel) TableName() string {
	return "payments"
}

type PaymentStatus string

const (
	PaymentStatusInitial PaymentStatus = "initial"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)
