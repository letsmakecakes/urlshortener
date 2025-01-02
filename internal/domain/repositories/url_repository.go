package repositories

import (
	"context"
	"urlshortener/internal/domain/models"
)

type URLRepository interface {
	CreateURL(ctx context.Context, url *models.URL) error
	GetURLByShortCode(ctx context.Context, shortCode string) (*models.URL, error)
	UpdateURL(ctx context.Context, url *models.URL) error
	DeleteURL(ctx context.Context, shortCode string) error
	IncrementURLAccessCount(ctx context.Context, shortCode string) error
}
