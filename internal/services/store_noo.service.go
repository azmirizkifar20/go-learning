package services

import (
	"context"

	"go-learning/internal/dtos"
	"go-learning/internal/repositories"
)

type StoreService struct {
	repo *repositories.StoreRepository
}

func NewStoreService(repo *repositories.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) GetStoreNoo(
	ctx context.Context,
	userID int64,
	page, limit int,
	filter dtos.StoreNooFilter,
) (*dtos.StoreNooResult, error) {

	return s.repo.GetStoreNoo(ctx, userID, page, limit, filter)
}
