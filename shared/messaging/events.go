package messaging

import pb "ride-sharing/shared/proto/trip"

const (
	FindAvailableDriversQueue       = "find_available_drivers"
	DriverCmdTripRequestQueue       = "driver_cmd_trip_request"
	DriverTripResponseQueue         = "driver_trip_response"
	NotifyDriverNoDriversFoundQueue = "notify_driver_no_drivers_found"
)

type TripEventData struct {
	Trip *pb.Trip `json:"trip"`
}
