package conn

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoDBConnection creates a new mongodb connection
func NewMongoDBConnection(ctx context.Context, uri string) (*mongo.Client, error) {

	if len(uri) <= 0 {
		return nil, errors.New("Cannot load MONGODB_URI from env, please provide one")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}
