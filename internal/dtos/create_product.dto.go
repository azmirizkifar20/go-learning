package dtos

type CreateProductRequest struct {
	ProductName string  `json:"product_name" form:"product_name" validate:"required,min=3"`
	CategoryID  uint    `json:"category_id" form:"category_id" validate:"required,gt=0"`
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"`
}
