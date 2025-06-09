export interface Trip {
    trip: Trip;
    id: string;
    userID: string;
    status: string;
    selectedFare: RouteFare;
    route: Route;
    driver?: Driver;
}

export interface RequestRideProps {
    pickup: [number, number],
    destination: [number, number],
}

export interface Coordinate {
    latitude: number,
    longitude: number,
}

export interface Route {
    geometry: {
        coordinates: Coordinate[]
    }[],
    duration: number,
    distance: number,
}

export enum CarPackageSlug {
    SEDAN = "sedan",
    SUV = "suv",
    VAN = "van",
    LUXURY = "luxury",
}

export interface RouteFare {
    id: string,
    packageSlug: CarPackageSlug,
    basePrice: number,
    totalPriceInCents?: number,
    expiresAt: Date,
    route: Route,
}


export interface HTTPTripStartResponse {
    tripID: string;
}

export interface TripPreview {
    tripID: string,
    route: [number, number][],
    rideFares: RouteFare[],
    duration: number,
    distance: number,
}


export interface Driver {
    id: string;
    location: Coordinate;
    geohash: string;
    name: string;
    profilePicture: string;
    carPlate: string;
}
