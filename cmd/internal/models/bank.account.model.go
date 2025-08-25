package models

type BankAccountModel struct {
	BaseModel
	UserId      uint   `json:"user_id" gorm:"column:user_id"`
	BankAccount string `json:"bank_account" gorm:"index;unique;not null;column:bank_account; varchar(250)"`
	SwiftCode   string `json:"swift_code" gorm:"column:swift_code;varchar(50);not null;"`
	PaymentType string `json:"payment_type" gorm:"column:payment_type;varchar(50);not null;"`
}

func (u BankAccountModel) TableName() string {
	return "bank_accounts"
}
