package geo

import "math"

// Coordinates is a Lat Long struct.
type Coordinates struct {
	Latitude, Longitude float64
}

// NewCoordinatesFromDegrees return Coordinates from lat and long in Degree.
func NewCoordinatesFromDegrees(latDeg, longDeg float64) *Coordinates {
	return &Coordinates{
		Latitude:  DegreeToRadians(latDeg),
		Longitude: DegreeToRadians(longDeg),
	}
}

// NewCoordinatesFromRadians return Coordinates from lat and long in Radian.
func NewCoordinatesFromRadians(latRads, longRads float64) *Coordinates {
	return &Coordinates{
		Latitude:  latRads,
		Longitude: longRads,
	}
}

// EarthRadiusKilometers ...
const EarthRadiusKilometers = 6371

// Distance type.
type Distance float64

// ToKilometers convert Distance to Kilometers.
func (d Distance) ToKilometers() float64 {
	return float64(d) * EarthRadiusKilometers
}

// ToMeters convert Distance to Meters.
func (d Distance) ToMeters() float64 {
	return d.ToKilometers() * 1000
}

// Distance between two points in radians (angular distance).
func (c *Coordinates) Distance(c2 *Coordinates) Distance {
	sinLat := math.Sin((c2.Latitude - c.Latitude) / 2)
	sinLng := math.Sin((c2.Longitude - c.Longitude) / 2)

	a := sinLat*sinLat + math.Cos(c.Latitude)*math.Cos(c2.Latitude)*sinLng*sinLng

	return Distance(2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)))
}
