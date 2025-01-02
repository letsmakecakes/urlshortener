package service

import (
	"context"
	"fmt"
	log "go.uber.org/zap"
	"time"
	"urlshortener/internal/domain/models"
	"urlshortener/internal/domain/repositories"
	"urlshortener/internal/pkg/generator"
)

// URLService provides methods to manage URLs
type URLService struct {
	repo repositories.URLRepository
}

// NewURLService creates a new instance of URLService
func NewURLService(repo repositories.URLRepository) *URLService {
	return &URLService{repo: repo}
}

// CreateShortURL creates a new shortened URL
func (s *URLService) CreateShortURL(ctx context.Context, originalURL string) (*models.URL, error) {
	shortCode := generator.GenerateShortCode()
	url := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		AccessCount: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateURL(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}

// GetURL retrieves a URL by its short code and increments the access count
func (s *URLService) GetURL(ctx context.Context, shortCode string) (*models.URL, error) {
	url, err := s.repo.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	if err := s.repo.IncrementURLAccessCount(ctx, shortCode); err != nil {
		log.Error(fmt.Errorf("error incrementing URL access acount: %v", err))
	}

	return url, nil
}

// UpdateURL updates the original URL of an existing short code.
func (s *URLService) UpdateURL(ctx context.Context, shortCode string, newURL string) (*models.URL, error) {
	url, err := s.repo.GetURLByShortCode(ctx, shortCode)
	if err != nil {
		return nil, err
	}

	url.OriginalURL = newURL
	url.UpdatedAt = time.Now()

	if err := s.repo.UpdateURL(ctx, url); err != nil {
		return nil, err
	}

	return url, nil
}

// DeleteURL defines a URL by its short code
func (s *URLService) DeleteURL(ctx context.Context, shortCode string) error {
	return s.repo.DeleteURL(ctx, shortCode)
}

// GetStats retrieves the statistics of a URL by its short code.
func (s *URLService) GetStats(ctx context.Context, shortCode string) (*models.URL, error) {
	return s.repo.GetURLByShortCode(ctx, shortCode)
}
