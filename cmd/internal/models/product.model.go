package models

type ProductModel struct {
	BaseModel
	Name        string  `json:"name" gorm:"index;column:name;type:varchar(250);not null"`
	Description string  `json:"description" gorm:"index;column:description;"`
	CategoryId  uint    `json:"category_id" gorm:"index;column:category_id;"`
	ImageUrl    string  `json:"image_url" gorm:"index;column:image_url;varchar(250);"`
	Price       float64 `json:"price" gorm:"index;column:price;"`
	UserId      int     `json:"user_id" gorm:"index;column:user_id;"`
	Stock       uint    `json:"stock" gorm:"index;column:stock;"`
}

func (u ProductModel) TableName() string {
	return "products"
}
