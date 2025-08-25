package models

type OrderItemModel struct {
	BaseModel
	OrderId   uint    `json:"order_id" gorm:"column:order_id;"`
	ProductId uint    `json:"product_id" gorm:"column:product_id;"`
	Name      string  `json:"name" gorm:"column:name;varchar(250)"`
	ImageUrl  string  `json:"image_url" gorm:"column:image_url;varchar(250)"`
	SellerId  uint    `json:"seller_id" gorm:"column:seller_id;"`
	Price     float64 `json:"price" gorm:"column:price;"`
	Qty       uint    `json:"qty" gorm:"column:qty;"`
}

func (u OrderItemModel) TableName() string {
	return "orders_items"
}
