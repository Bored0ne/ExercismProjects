// the space package provides a utility to calculate ones age in seconds
package space

type Planet string

const EarthYearInSeconds = 31557600

const MERCURY = "Mercury"
const VENUS = "Venus"
const EARTH = "Earth"
const MARS = "Mars"
const JUPITER = "Jupiter"
const SATURN = "Saturn"
const URANUS = "Uranus"
const NEPTUNE = "Neptune"

// Age accepts a currentAgeInSeconds argument and a Planet Argument and returns the number of years you are in float64 format.
func Age(currentAgeInSeconds float64, planet Planet) float64 {
	if planet == MERCURY {
		return currentAgeInSeconds / (0.2408467 * EarthYearInSeconds)
	}
	if planet == VENUS {
		return currentAgeInSeconds / (0.61519726 * EarthYearInSeconds)
	}
	if planet == EARTH {
		return currentAgeInSeconds / EarthYearInSeconds
	}
	if planet == MARS {
		return currentAgeInSeconds / (1.8808158 * EarthYearInSeconds)
	}
	if planet == JUPITER {
		return currentAgeInSeconds / (11.862615 * EarthYearInSeconds)
	}
	if planet == SATURN {
		return currentAgeInSeconds / (29.447498 * EarthYearInSeconds)
	}
	if planet == URANUS {
		return currentAgeInSeconds / (84.016846 * EarthYearInSeconds)
	}
	if planet == NEPTUNE {
		return currentAgeInSeconds / (164.79132 * EarthYearInSeconds)
	}
	return -1
}
