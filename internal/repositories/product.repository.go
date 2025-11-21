package repositories

import (
	"gorm.io/gorm"

	"go-learning/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(p *models.Product) error {
	return r.db.Create(p).Error
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var p models.Product
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *models.Product) error {
	return r.db.Save(p).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) List() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
