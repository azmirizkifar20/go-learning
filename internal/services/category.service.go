package services

import (
	"mime/multipart"

	"go-learning/internal/models"
	"go-learning/internal/repositories"
	"go-learning/internal/storage"

	"gorm.io/gorm"
)

type CateogyService interface {
	CreateCategory(db *gorm.DB, name string, imageFile *multipart.FileHeader) (*models.Category, error)
}

type categoryService struct {
	repo  repositories.CategoryRepository
	minio *storage.MinioClient
}

func NewCategoryService(repo repositories.CategoryRepository, minio *storage.MinioClient) CateogyService {
	return &categoryService{
		repo,
		minio,
	}
}

func (s *categoryService) CreateCategory(db *gorm.DB, name string, imageFile *multipart.FileHeader) (*models.Category, error) {
	imageURL := ""
	var err error

	if imageFile != nil {
		imageURL, err = s.minio.UploadCategoryImage(imageFile)
		if err != nil {
			return nil, err
		}
	}

	category := &models.Category{
		CategoryName: name,
		ImageURL:     imageURL,
	}

	if err := s.repo.Create(db, category); err != nil {
		return nil, err
	}

	return category, nil
}
