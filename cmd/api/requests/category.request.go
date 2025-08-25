package requests

type CreateCategoryRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=250"`
	ParentId     uint   `json:"parent_id" validate:"required"`
	ImageUrl     string `json:"image_url" validate:"required"`
	DisplayOrder int    `json:"display_order" validate:"required"`
}
