package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDbClient() (*mongo.Client, error) {
	dbAddress := os.Getenv("MONGO_URI")
	if dbAddress == "" {
		return nil, fmt.Errorf("MONGO_URI not set")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(dbAddress).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err = client.Ping(pingCtx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %w", err)
	}

	return client, nil
}
