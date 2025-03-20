import Geohash from 'latlon-geohash';

export function decodeGeoHash(geohash: string) {
  const bounds = Geohash.bounds(geohash);
  return {
    latitude: [bounds.sw.lat, bounds.ne.lat],
    longitude: [bounds.sw.lon, bounds.ne.lon],
  };
}

// Function to create grid bounds from geohash
export const getGeohashBounds = (geohash: string) => {
  const { latitude: [minLat, maxLat], longitude: [minLng, maxLng] } = decodeGeoHash(geohash);
  return [
    [minLat, minLng],
    [maxLat, maxLng],
  ];
};