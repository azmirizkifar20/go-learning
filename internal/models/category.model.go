package models

import "time"

type Category struct {
	CategoryID   uint      `gorm:"primaryKey" json:"category_id"`
	CategoryName string    `gorm:"column:category_name" json:"category_name"`
	ImageURL     string    `gorm:"column:image_url" json:"image_url"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
