package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"go-learning/internal/config"
)

type MinioClient struct {
	Client        *minio.Client
	Bucket        string
	PublicBaseURL string
}

func NewMinioClient() *MinioClient {
	cfg := config.LoadConfig()

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if err != nil {
		log.Fatalf("failed to init MinIO: %v", err)
	}

	// ensure bucket exists
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, cfg.MinioBucket)
	if err != nil {
		log.Fatalf("failed to check bucket: %v", err)
	}
	if !exists {
		if err := minioClient.MakeBucket(ctx, cfg.MinioBucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatalf("failed to create bucket: %v", err)
		}
	}

	return &MinioClient{
		Client:        minioClient,
		Bucket:        cfg.MinioBucket,
		PublicBaseURL: cfg.MinioPublicBaseURL,
	}
}

func (m *MinioClient) UploadCategoryImage(fileHeader *multipart.FileHeader) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	objectName := fmt.Sprintf("category-%s%s", uuid.New().String(), ext)

	_, err = m.Client.PutObject(
		ctx,
		m.Bucket,
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: fileHeader.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return "", err
	}

	if m.PublicBaseURL != "" {
		return fmt.Sprintf("%s/%s", m.PublicBaseURL, objectName), nil
	}

	// fallback (kalau public base URL gak di set)
	return objectName, nil
}
