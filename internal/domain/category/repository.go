package category

import "gorm.io/gorm"

type Repository interface {
	Create(db *gorm.DB, category *Category) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(db *gorm.DB, category *Category) error {
	return db.Create(category).Error
}
