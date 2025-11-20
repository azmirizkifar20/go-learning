package repositories

import (
	"gorm.io/gorm"

	"go-learning/internal/models"
)

type CategoryRepository interface {
	Create(db *gorm.DB, category *models.Category) error
}

type repository struct{}

func NewCategoryRepository() CategoryRepository {
	return &repository{}
}

func (r *repository) Create(db *gorm.DB, category *models.Category) error {
	return db.Create(category).Error
}
