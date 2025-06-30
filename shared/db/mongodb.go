package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	TripsCollection     = "trips"
	RideFaresCollection = "ride_fares"
)

// MongoConfig holds MongoDB connection configuration
type MongoConfig struct {
	URI      string
	Database string
}

// NewMongoConfig creates a new MongoDB configuration from environment variables
func NewMongoDefaultConfig() *MongoConfig {
	return &MongoConfig{
		URI:      os.Getenv("MONGODB_URI"),
		Database: "ride-sharing",
	}
}

// NewMongoClient creates a new MongoDB client
func NewMongoClient(ctx context.Context, cfg *MongoConfig) (*mongo.Client, error) {
	if cfg.URI == "" {
		return nil, fmt.Errorf("mongodb URI is required")
	}
	if cfg.Database == "" {
		return nil, fmt.Errorf("mongodb database is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to MongoDB at %s", cfg.URI)
	return client, nil
}

// GetDatabase returns a database instance
func GetDatabase(client *mongo.Client, cfg *MongoConfig) *mongo.Database {
	return client.Database(cfg.Database)
}
