package main

import pb "ride-sharing/shared/proto/driver"

type driverInMap struct {
	Driver *pb.Driver
	// Index int
	// TODO: route
}

type Service struct {
	drivers []*driverInMap
}

func NewService() *Service {
	return &Service{
		drivers: make([]*driverInMap, 0),
	}
}

// TODO: Register and unregister methods