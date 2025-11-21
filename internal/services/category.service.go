package services

import (
	"mime/multipart"

	"go-learning/internal/models"
	"go-learning/internal/repositories"
	"go-learning/internal/storage"
)

type CategoryService struct {
	repo  *repositories.CategoryRepository
	minio *storage.MinioClient
}

func NewCategoryService(repo *repositories.CategoryRepository, minio *storage.MinioClient) *CategoryService {
	return &CategoryService{
		repo,
		minio,
	}
}

func (s *CategoryService) CreateCategory(name string, imageFile *multipart.FileHeader) (*models.Category, error) {
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

	if err := s.repo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}
