package product

import "gorm.io/gorm"

type Service interface {
	Create(db *gorm.DB, p *Product) (*Product, error)
	Get(db *gorm.DB, id uint) (*Product, error)
	Update(db *gorm.DB, id uint, input *Product) (*Product, error)
	Delete(db *gorm.DB, id uint) error
	List(db *gorm.DB) ([]Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(db *gorm.DB, p *Product) (*Product, error) {
	if err := s.repo.Create(db, p); err != nil {
		return nil, err
	}
	return p, nil
}

// sisanya simple wrapper ke repo

func (s *service) Get(db *gorm.DB, id uint) (*Product, error) {
	return s.repo.GetByID(db, id)
}

func (s *service) Update(db *gorm.DB, id uint, input *Product) (*Product, error) {
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

func (s *service) Delete(db *gorm.DB, id uint) error {
	return s.repo.Delete(db, id)
}

func (s *service) List(db *gorm.DB) ([]Product, error) {
	return s.repo.List(db)
}
