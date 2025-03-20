import { Bus, Truck, Crown } from "lucide-react";
import { Car } from "lucide-react";
import { CarPackageSlug } from "../types";

export const PackagesMeta: Record<CarPackageSlug, {
  name: string,
  icon: React.ReactNode,
  description: string,
}> = {
  [CarPackageSlug.SEDAN]: {
    name: "Sedan",
    icon: <Car />,
    description: "Economic and comfortable",
  },
  [CarPackageSlug.SUV]: {
    name: "SUV",
    icon: <Truck />,
    description: "Spacious ride for groups",
  },
  [CarPackageSlug.VAN]: {
    name: "Van",
    icon: <Bus />,
    description: "Perfect for larger groups",
  },
  [CarPackageSlug.LUXURY]: {
    name: "Luxury",
    icon: <Crown />,
    description: "Premium experience",
  },
}