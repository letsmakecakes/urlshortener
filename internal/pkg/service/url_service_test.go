package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"urlshortener/internal/domain/models"
)

// MockURLRepository is a mock implementation of the URLRepository interface
type MockURLRepository struct {
	mock.Mock
}

// CreateURL creates a new URL in the repository
func (m *MockURLRepository) CreateURL(ctx context.Context, url *models.URL) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

// GetURLByShortCode retrieves a URL by its short code from the repository
func (m *MockURLRepository) GetURLByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	args := m.Called(ctx, shortCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

// UpdateURL updates an existing URL in the repository
func (m *MockURLRepository) UpdateURL(ctx context.Context, url *models.URL) error {
	args := m.Called(ctx, url)
	return args.Error(0)
}

// DeleteURL deletes a URL by its short code from the repository
func (m *MockURLRepository) DeleteURL(ctx context.Context, shortCode string) error {
	args := m.Called(ctx, shortCode)
	return args.Error(0)
}

// IncrementURLAccessCount increments the access count of a URL in the repository
func (m *MockURLRepository) IncrementURLAccessCount(ctx context.Context, shortCode string) error {
	args := m.Called(ctx, shortCode)
	return args.Error(0)
}

// TestURLService_CreateShortURL tests the CreateShortURL method of the URLService
func TestURLService_CreateShortURL(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLService(mockRepo)
	ctx := context.Background()

	testURL := "https://example.com"

	// Set up the mock to expect a CreateURL call and return no error
	mockRepo.On("CreateURL", ctx, mock.AnythingOfType("*models.URL")).Return(nil)

	// Call the CreateShortURL method
	result, err := service.CreateShortURL(ctx, testURL)

	// Assert no error and valid result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testURL, result.OriginalURL)
	assert.NotEmpty(t, result.ShortCode)

	// Ensure expectations were met
	mockRepo.AssertExpectations(t)
}
