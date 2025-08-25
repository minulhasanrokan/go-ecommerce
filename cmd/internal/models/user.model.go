package models

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type UserModel struct {
	BaseModel
	FirstName string         `gorm:"column:first_name;type:varchar(250)" json:"first_name"`
	LastName  string         `gorm:"column:last_name;type:varchar(250)" json:"last_name"`
	Email     string         `gorm:"column:email;type:varchar(100)" json:"email"`
	Mobile    string         `gorm:"column:mobile;type:varchar(50)" json:"mobile"`
	Password  string         `gorm:"column:password;type:varchar(100)" json:"-"`
	Code      int            `gorm:"column:code" json:"code"`
	ExpiredAt time.Time      `gorm:"column:expired_at" json:"expired_at"`
	Verified  bool           `gorm:"column:verified;default:false" json:"verified"`
	UserType  string         `gorm:"column:user_type;type:varchar(250);default:buyer" json:"user_type"`
	Address   AddressModel   `gorm:"foreignKey:UserId" json:"address"`
	Cart      CartModel      `json:"cart" gorm:"foreignKey:UserId"`
	Orders    []OrderModel   `json:"orders" gorm:"foreignKey:UserId"`
	Payments  []PaymentModel `json:"payment" gorm:"foreignKey:UserId"`
}

func (u UserModel) TableName() string {
	return "users"
}
