package services

import (
	"gorm.io/gorm"

	"go-learning/internal/models"
	"go-learning/internal/repositories"
)

type ProductService interface {
	Create(db *gorm.DB, p *models.Product) (*models.Product, error)
	Get(db *gorm.DB, id uint) (*models.Product, error)
	Update(db *gorm.DB, id uint, input *models.Product) (*models.Product, error)
	Delete(db *gorm.DB, id uint) error
	List(db *gorm.DB) ([]models.Product, error)
}

type productservice struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productservice{repo}
}

func (s *productservice) Create(db *gorm.DB, p *models.Product) (*models.Product, error) {
	if err := s.repo.Create(db, p); err != nil {
		return nil, err
	}
	return p, nil
}

// sisanya simple wrapper ke repo

func (s *productservice) Get(db *gorm.DB, id uint) (*models.Product, error) {
	return s.repo.GetByID(db, id)
}

func (s *productservice) Update(db *gorm.DB, id uint, input *models.Product) (*models.Product, error) {
	p, err := s.repo.GetByID(db, id)
	if err != nil {
		return nil, err
	}

	p.ProductName = input.ProductName
	p.CategoryID = input.CategoryID
	p.Price = input.Price

	if err := s.repo.Update(db, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *productservice) Delete(db *gorm.DB, id uint) error {
	return s.repo.Delete(db, id)
}

func (s *productservice) List(db *gorm.DB) ([]models.Product, error) {
	return s.repo.List(db)
}
