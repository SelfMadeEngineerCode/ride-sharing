import { useMapEvents } from 'react-leaflet'

interface MapClickHandlerProps {
  onClick: (e: L.LeafletMouseEvent) => void;
}

export function MapClickHandler({ onClick }: MapClickHandlerProps) {
  useMapEvents({
    click: onClick,
  })
  return null
}

