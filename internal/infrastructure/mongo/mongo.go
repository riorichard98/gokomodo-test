package mongo

import (
	"context"
	"gdk/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// NewMongoDB creates a new MongoDB client and returns it.
func NewMongoDB(config *config.MongoConfig) (mongoClient *mongo.Client, err error) {
	// Create a context with a timeout based on config.Timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.Timeout)*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.URI)

	// Set max and min pool size options
	clientOptions.SetMaxPoolSize(config.MaxPoolSize)
	clientOptions.SetMinPoolSize(config.MinPoolSize)

	// Connect to MongoDB
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		// Panic if connection credentials are invalid
		panic("connection credential is invalid")
	}

	// Ping the MongoDB server to check if the connection is established
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		// Panic if MongoDB can't connect
		panic("mongodb can't connect")
	}

	return
}
