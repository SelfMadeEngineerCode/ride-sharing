package main

import (
	"context"
	pb "ride-sharing/shared/proto/driver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type driverGrpcHandler struct {
	pb.UnimplementedDriverServiceServer

	service *Service
}

func NewGrpcHandler(s *grpc.Server, service *Service) {
	handler := &driverGrpcHandler{
		service: service,
	}

	pb.RegisterDriverServiceServer(s, handler)
}

func (h *driverGrpcHandler) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	driver, err := h.service.RegisterDriver(req.GetDriverID(), req.GetPackageSlug())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to register driver")
	}

	return &pb.RegisterDriverResponse{
		Driver: driver,
	}, nil
}

func (h *driverGrpcHandler) UnregisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
	h.service.UnregisterDriver(req.GetDriverID())

	return &pb.RegisterDriverResponse{
		Driver: &pb.Driver{
			Id: req.GetDriverID(),
		},
	}, nil
}
