
export function convertSecondsToMinutes(seconds: number) {
  return `${Math.floor(seconds / 60)} minutes`
}

export function convertMetersToKilometers(meters: number) {
  return `${(meters / 1000).toFixed(2)} km`
}