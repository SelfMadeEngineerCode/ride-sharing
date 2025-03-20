import { Driver, CarPackageSlug } from "../types";
import Image from "next/image";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

export const DriverCard = ({ driver, packageSlug }: { driver?: Driver | null, packageSlug?: CarPackageSlug }) => {
  if (!driver) return null;

  const CarPlate = ({ plate }: { plate: string }) => (
    <span className="inline-flex items-center px-2.5 py-0.5 rounded-md text-sm font-medium bg-gray-100 text-gray-800 font-mono tracking-wider">
      {plate.toUpperCase()}
    </span>
  );

  return (
    <Card className="">
      <CardHeader>
        <CardTitle>{driver.name}</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-2 items-center">
        {driver.profilePicture && (
          <Image
            className="rounded-full"
            src={driver.profilePicture}
            alt={`${driver.name}'s profile picture`}
            width={50}
            height={50}
          />
        )}

        {driver.carPlate && (
          <p className="text-sm">
            <CarPlate plate={driver.carPlate} />
          </p>
        )}

        {packageSlug && (
          <p className="text-sm">
            <span className="font-mono">{packageSlug}</span> driver
          </p>
        )}
      </CardContent>
    </Card>
  )
};
