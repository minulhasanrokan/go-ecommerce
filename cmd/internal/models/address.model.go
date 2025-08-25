package models

type AddressModel struct {
	BaseModel
	AddressLine1 string `json:"address_line1" gorm:"column:address_line_1;type:varchar(250)"`
	AddressLine2 string `json:"address_line2" gorm:"column:address_line_2;type:varchar(250)"`
	City         string `json:"city" gorm:"column:city;type:varchar(250)"`
	PostCode     uint   `json:"postCode" gorm:"column:post_code;type:int(11)"`
	Country      string `json:"country" gorm:"column:country;type:varchar(250)"`
	UserId       uint   `json:"user_id" gorm:"column:user_id;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u AddressModel) TableName() string {
	return "addresses"
}
