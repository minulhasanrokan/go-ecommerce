package models

type CategoryModel struct {
	BaseModel
	Name         string         `json:"name" gorm:"index;column:name;type:varchar(255);not null"`
	ParentId     uint           `json:"parent_id" gorm:"index;column:parent_id"`
	ImageUrl     string         `json:"image_url" gorm:"column:image_url;type:varchar(255);"`
	Products     []ProductModel `json:"products" gorm:"foreignkey:CategoryId"`
	DisplayOrder int            `json:"display_order" gorm:"column:display_order"`
}

func (u CategoryModel) TableName() string {
	return "categories"
}
