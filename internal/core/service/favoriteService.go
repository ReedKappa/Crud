package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
)

type _favoriteService struct {
	repo repository.FavoriteRepository
}

func (favoriteService _favoriteService) AddFavorite(ctx context.Context, login string, bookId int) error {
	return favoriteService.repo.AddFavorite(ctx, login, bookId)
}

func (favoriteService _favoriteService) GetFavorite(ctx context.Context, login string) ([]model.Book, error) {
	return favoriteService.repo.GetFavorite(ctx, login)
}

func NewFavoriteService(repo repository.FavoriteRepository) service.FavoriteService {
	return _favoriteService{repo: repo}
}
