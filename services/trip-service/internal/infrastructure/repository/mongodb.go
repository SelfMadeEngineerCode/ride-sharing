package repository

import (
	"context"
	"fmt"

	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/db"
	pbd "ride-sharing/shared/proto/driver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *mongoRepository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	result, err := r.db.Collection(db.TripsCollection).InsertOne(ctx, trip)
	if err != nil {
		return nil, err
	}

	trip.ID = result.InsertedID.(primitive.ObjectID)

	return trip, nil
}

func (r *mongoRepository) GetTripByID(ctx context.Context, id string) (*domain.TripModel, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Collection(db.TripsCollection).FindOne(ctx, bson.M{"_id": _id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var trip domain.TripModel
	err = result.Decode(&trip)
	if err != nil {
		return nil, err
	}

	return &trip, nil
}

func (r *mongoRepository) UpdateTrip(ctx context.Context, tripID string, status string, driver *pbd.Driver) error {
	_id, err := primitive.ObjectIDFromHex(tripID)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"status": status}}

	if driver != nil {
		update["$set"].(bson.M)["driver"] = driver
	}

	result, err := r.db.Collection(db.TripsCollection).UpdateOne(ctx, bson.M{"_id": _id}, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("trip not found: %s", tripID)
	}

	return nil
}

func (r *mongoRepository) SaveRideFare(ctx context.Context, fare *domain.RideFareModel) error {
	result, err := r.db.Collection(db.RideFaresCollection).InsertOne(ctx, fare)
	if err != nil {
		return err
	}

	fare.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *mongoRepository) GetRideFareByID(ctx context.Context, id string) (*domain.RideFareModel, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result := r.db.Collection(db.RideFaresCollection).FindOne(ctx, bson.M{"_id": _id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	var fare domain.RideFareModel
	err = result.Decode(&fare)
	if err != nil {
		return nil, err
	}

	return &fare, nil
}
