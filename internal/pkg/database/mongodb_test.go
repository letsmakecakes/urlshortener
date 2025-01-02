package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"urlshortener/internal/domain/models"
)

// setupTestDB sets up a MongoDB test database and returns a repository and a cleanup function.
func setupTestDB(t *testing.T) (*MongoURLRepository, func()) {
	ctx := context.Background()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}

	// Get a handle to the test database
	db := client.Database("urlshortener_test")

	// Create a repository for testing
	repo := &MongoURLRepository{
		collection: db.Collection("urls_test"),
	}

	// Return the repository and a cleanup function to drop the collection and disconnect the client
	return repo, func() {
		repo.collection.Drop(ctx)
		client.Disconnect(ctx)
	}
}

// TestMongoURLRepository_CreateAndGet tests creating and retrieving a URL in the repository.
func TestMongoURLRepository_CreateAndGet(t *testing.T) {
	// Set up the test database and ensure cleanup is called after the test
	repo, cleanup := setupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test URL
	url := &models.URL{
		OriginalURL: "https://example.com",
		ShortCode:   "test123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Test creating the URL in the repository
	err := repo.CreateURL(ctx, url)
	assert.NoError(t, err)

	// Test retrieving the URL by its short code
	retrieved, err := repo.GetURLByShortCode(ctx, url.ShortCode)
	assert.NoError(t, err)
	assert.Equal(t, url.OriginalURL, retrieved.OriginalURL)
}
