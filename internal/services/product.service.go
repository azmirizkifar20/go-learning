package services

import (
	"context"
	"go-learning/internal/models"
	"go-learning/internal/repositories"
)

type Productservice struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *Productservice {
	return &Productservice{repo}
}

func (s *Productservice) Create(ctx context.Context, p *models.Product) (*models.Product, error) {
	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Productservice) Get(ctx context.Context, id uint) (*models.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Productservice) Update(ctx context.Context, id uint, input *models.Product) (*models.Product, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	p.ProductName = input.ProductName
	p.CategoryID = input.CategoryID
	p.Price = input.Price

	if err := s.repo.Update(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Productservice) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *Productservice) List(ctx context.Context) ([]models.Product, error) {
	return s.repo.List(ctx)
}
