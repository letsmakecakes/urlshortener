package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"urlshortener/internal/domain/models"
	"urlshortener/internal/domain/repositories"
)

// MongoURLRepository implements the URLRepository interface using MongoDB as the storage.
type MongoURLRepository struct {
	db         *MongoDB
	collection *mongo.Collection
}

// NewMongoURLRepository creates a new instance of MongoURLRepository
func NewMongoURLRepository(db *MongoDB) repositories.URLRepository {
	return &MongoURLRepository{
		db:         db,
		collection: db.Collection("urls"),
	}
}

// CreateURL inserts a new URL document into the MongoDB collection.
func (r *MongoURLRepository) CreateURL(ctx context.Context, url *models.URL) error {
	_, err := r.collection.InsertOne(ctx, url)
	return err
}

// GetURLByShortCode retrieves a URL document by its short code.
func (r *MongoURLRepository) GetURLByShortCode(ctx context.Context, shortCode string) (*models.URL, error) {
	var url models.URL
	err := r.collection.FindOne(ctx, bson.M{"short_code": shortCode}).Decode(&url)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("url not found")
		}
		return nil, err
	}
	return &url, err
}

// UpdateURL modifies an existing URL document in the MongoDB collection.
func (r *MongoURLRepository) UpdateURL(ctx context.Context, url *models.URL) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"short_code": url.ShortCode},
		bson.M{
			"$set": bson.M{
				"original_url": url.OriginalURL,
				"updated_at":   url.UpdatedAt,
			},
		},
	)
	return err
}

// DeleteURL removes a URL document from the MongoDB collection by its short code.
func (r *MongoURLRepository) DeleteURL(ctx context.Context, shortCode string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"short_code": shortCode})
	return err
}

// IncrementURLAccessCount increments the access count of a URL document by its short code.
func (r *MongoURLRepository) IncrementURLAccessCount(ctx context.Context, shortCode string) error {
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"short_code": shortCode},
		bson.M{"$inc": bson.M{"access_count": 1}},
	)
	return err
}
