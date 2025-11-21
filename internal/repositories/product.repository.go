package repositories

import (
	"context"

	"gorm.io/gorm"

	"go-learning/internal/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p *models.Product) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *ProductRepository) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	var p models.Product
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(ctx context.Context, p *models.Product) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Product{}, id).Error
}

func (r *ProductRepository) List(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
