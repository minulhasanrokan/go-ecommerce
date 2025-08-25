package requests

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=250"`
	Description string  `json:"description" validate:"required,min=3,max=250"`
	CategoryId  uint    `json:"category_id" validate:"required"`
	ImageUrl    string  `json:"image_url" validate:"required,url"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock"`
}
