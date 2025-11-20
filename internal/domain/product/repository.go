package product

import "gorm.io/gorm"

type Repository interface {
	Create(db *gorm.DB, p *Product) error
	GetByID(db *gorm.DB, id uint) (*Product, error)
	Update(db *gorm.DB, p *Product) error
	Delete(db *gorm.DB, id uint) error
	List(db *gorm.DB) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Create(db *gorm.DB, p *Product) error {
	return db.Create(p).Error
}

func (r *repository) GetByID(db *gorm.DB, id uint) (*Product, error) {
	var p Product
	if err := db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) Update(db *gorm.DB, p *Product) error {
	return db.Save(p).Error
}

func (r *repository) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&Product{}, id).Error
}

func (r *repository) List(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
