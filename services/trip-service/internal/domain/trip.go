package domain

import (
	"context"
	"ride-sharing/shared/types"

	tripTypes "ride-sharing/services/trip-service/pkg/types"
	pb "ride-sharing/shared/proto/trip"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TripModel struct {
	ID       primitive.ObjectID
	UserID   string
	Status   string
	RideFare *RideFareModel
	Driver   *pb.TripDriver
}

type TripRepository interface {
	CreateTrip(ctx context.Context, trip *TripModel) (*TripModel, error)
	SaveRideFare(ctx context.Context, f *RideFareModel) error
	GetRideFareByID(ctx context.Context, id string) (*RideFareModel, error)
}

type TripService interface {
	CreateTrip(ctx context.Context, fare *RideFareModel) (*TripModel, error)
	GetRoute(ctx context.Context, pickup, destination *types.Coordinate, useOsrmApi bool) (*tripTypes.OsrmApiResponse, error)
	EstimatePackagesPriceWithRoute(route *tripTypes.OsrmApiResponse) []*RideFareModel
	GenerateTripFares(
		ctx context.Context,
		fares []*RideFareModel,
		userID string,
		Route *tripTypes.OsrmApiResponse,
	) ([]*RideFareModel, error)
	GetAndValidateFare(ctx context.Context, fareID, userID string) (*RideFareModel, error)
}
