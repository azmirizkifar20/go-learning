package models

import "time"

type Product struct {
	ProductID   uint      `gorm:"primaryKey" json:"product_id"`
	ProductName string    `gorm:"column:product_name" json:"product_name"`
	CategoryID  uint      `gorm:"column:category_id" json:"category_id"`
	Price       float64   `gorm:"column:price" json:"price"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Product) TableName() string {
	return "products"
}
