package repositories

import (
	"gorm.io/gorm"

	"go-learning/internal/models"
)

type ProductRepository interface {
	Create(db *gorm.DB, p *models.Product) error
	GetByID(db *gorm.DB, id uint) (*models.Product, error)
	Update(db *gorm.DB, p *models.Product) error
	Delete(db *gorm.DB, id uint) error
	List(db *gorm.DB) ([]models.Product, error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(db *gorm.DB, p *models.Product) error {
	return db.Create(p).Error
}

func (r *productRepository) GetByID(db *gorm.DB, id uint) (*models.Product, error) {
	var p models.Product
	if err := db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepository) Update(db *gorm.DB, p *models.Product) error {
	return db.Save(p).Error
}

func (r *productRepository) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) List(db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
