package category

import (
	"mime/multipart"

	"go-learning/internal/storage"

	"gorm.io/gorm"
)

type Service interface {
	CreateCategory(db *gorm.DB, name string, imageFile *multipart.FileHeader) (*Category, error)
}

type service struct {
	repo  Repository
	minio *storage.MinioClient
}

func NewService(repo Repository, minio *storage.MinioClient) Service {
	return &service{
		repo,
		minio,
	}
}

func (s *service) CreateCategory(db *gorm.DB, name string, imageFile *multipart.FileHeader) (*Category, error) {
	imageURL := ""
	var err error

	if imageFile != nil {
		imageURL, err = s.minio.UploadCategoryImage(imageFile)
		if err != nil {
			return nil, err
		}
	}

	category := &Category{
		CategoryName: name,
		ImageURL:     imageURL,
	}

	if err := s.repo.Create(db, category); err != nil {
		return nil, err
	}

	return category, nil
}
