package services

import (
	"go-learning/internal/models"
	"go-learning/internal/repositories"
)

type Productservice struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *Productservice {
	return &Productservice{repo}
}

func (s *Productservice) Create(p *models.Product) (*models.Product, error) {
	if err := s.repo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

// sisanya simple wrapper ke repo

func (s *Productservice) Get(id uint) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *Productservice) Update(id uint, input *models.Product) (*models.Product, error) {
	p, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	p.ProductName = input.ProductName
	p.CategoryID = input.CategoryID
	p.Price = input.Price

	if err := s.repo.Update(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Productservice) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *Productservice) List() ([]models.Product, error) {
	return s.repo.List()
}
