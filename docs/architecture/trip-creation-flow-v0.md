```mermaid
sequenceDiagram
    participant User
    participant TripService
    participant DriverService
    participant Driver

    User->>TripService: Create Trip Request
    TripService-->>+DriverService: trip.event.created
    Note right of TripService: Event: Something happened

    DriverService->>Driver: driver.cmd.trip_request
    Note right of DriverService: Command: Please do this

    Driver-->>DriverService: driver.cmd.trip_accept
    Note left of Driver: Command: Response

    DriverService->>-TripService: trip.event.driver_assigned
    Note left of DriverService: Event: Something happened
```
