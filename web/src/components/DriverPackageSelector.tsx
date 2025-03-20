import { PackagesMeta } from './PackagesMeta'
import { CarPackageSlug } from '../types'
import { cn } from "../lib/utils"

interface DriverPackageSelectorProps {
  onSelect: (packageSlug: CarPackageSlug) => void
}

export function DriverPackageSelector({ onSelect }: DriverPackageSelectorProps) {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <div className="bg-white w-full h-full sm:h-auto sm:rounded-2xl sm:shadow-lg sm:max-w-md sm:mx-4 p-4 sm:p-6">
        <h2 className="text-lg sm:text-xl font-semibold mb-2">Select your car type</h2>
        <p className="text-sm text-gray-500 mb-6">Choose the type of car you&apos;ll be driving</p>
        <div className="space-y-3 sm:space-y-4">
          {Object.entries(PackagesMeta).map(([slug, meta]) => (
            <div
              key={slug}
              className={cn(
                "flex items-center gap-3 sm:gap-4 p-3 sm:p-4 sm:rounded-lg sm:border transition-all cursor-pointer",
                "hover:border-primary hover:bg-primary/5",
              )}
              onClick={() => onSelect(slug as CarPackageSlug)}
            >
              <div className="p-1.5 sm:p-2 bg-gray-100 rounded-lg">
                {meta?.icon}
              </div>
              <div>
                <h3 className="font-medium text-sm sm:text-base">{meta?.name}</h3>
                <p className="text-xs sm:text-sm text-gray-500">{meta?.description}</p>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}