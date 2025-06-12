package messaging

import (
	pbd "ride-sharing/shared/proto/driver"
	pb "ride-sharing/shared/proto/trip"
)

const (
	FindAvailableDriversQueue       = "find_available_drivers"
	DriverCmdTripRequestQueue       = "driver_cmd_trip_request"
	DriverTripResponseQueue         = "driver_trip_response"
	NotifyDriverNoDriversFoundQueue = "notify_driver_no_drivers_found"
	NotifyDriverAssignQueue         = "notify_driver_assign"
)

type TripEventData struct {
	Trip *pb.Trip `json:"trip"`
}

type DriverTripResponseData struct {
	Driver  *pbd.Driver `json:"driver"`
	TripID  string      `json:"tripID"`
	RiderID string      `json:"riderID"`
}
