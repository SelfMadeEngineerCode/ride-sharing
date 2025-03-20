import { Trip } from "../types"
import { TripOverviewCard } from "./TripOverviewCard"
import { Button } from "./ui/button"
import { TripEvents } from "../contracts"

interface DriverTripOverviewProps {
  trip?: Trip | null,
  status?: TripEvents | null,
  onAcceptTrip?: () => void,
  onDeclineTrip?: () => void
}

export const DriverTripOverview = ({ trip, status, onAcceptTrip, onDeclineTrip }: DriverTripOverviewProps) => {
  if (!trip) {
    return (
      <TripOverviewCard
        title="Waiting for a rider..."
        description="Waiting for a rider to request a trip..."
      />
    )
  }

  if (status === TripEvents.DriverTripRequest) {
    return (
      <TripOverviewCard
        title="Trip request received!"
        description="A trip has been requested, check the route and accept the trip if you can take it."
      >
        <div className="flex flex-col gap-2">
          <Button onClick={onAcceptTrip}>Accept trip</Button>
          <Button variant="outline" onClick={onDeclineTrip}>Decline trip</Button>
        </div>
      </TripOverviewCard>
    )
  }

  if (status === TripEvents.DriverTripAccept) {
    return (
      <TripOverviewCard
        title="All set!"
        description="You can now start the trip"
      >
        <div className="flex flex-col gap-4">
          <div className="flex flex-col gap-2">
            <h3 className="text-lg font-bold">Trip details</h3>
            <p className="text-sm text-gray-500">
              Trip ID: {trip.id}
              <br />
              Rider ID: {trip.userID}
            </p>
          </div>
        </div>
      </TripOverviewCard>
    )
  }

  return null
}