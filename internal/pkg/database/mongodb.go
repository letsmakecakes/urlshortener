package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

// NewMongoDB initializes a new MongoDB client and returns a MongoDB struct.
// It connects to the MongoDB server using the provided URI and database name.
func NewMongoDB(uri, dbName string) (*MongoDB, error) {
	// Create a context with a timeout for the MongoDB connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Verify the MongoDB connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoDB{
		client: client,
		db:     client.Database(dbName),
	}, nil
}

// Collection returns a reference to the specified collection in the MongoDB database.
func (m *MongoDB) Collection(name string) *mongo.Collection {
	return m.db.Collection(name)
}

// Disconnect safely closes the connection to the MongoDB server.
func (m *MongoDB) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
