package geo

import "math"

// Coordinates is a Lat Long struct.
type Coordinates struct {
	Latitude, Longitude Radian
}

// NewCoordinatesFromDegrees return Coordinates from lat and long in Degree.
func NewCoordinatesFromDegrees(latDeg, longDeg Degree) *Coordinates {
	return &Coordinates{
		Latitude:  latDeg.ToRadians(),
		Longitude: longDeg.ToRadians(),
	}
}

// NewCoordinatesFromRadians return Coordinates from lat and long in Radian.
func NewCoordinatesFromRadians(latRads, longRads Radian) *Coordinates {
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

// Distance between two points in Coordinates.
func (c *Coordinates) Distance(c2 *Coordinates) Distance {
	sinLat := math.Sin((c2.Latitude.Float64() - c.Latitude.Float64()) / 2)
	sinLng := math.Sin((c2.Longitude.Float64() - c.Longitude.Float64()) / 2)

	a := sinLat*sinLat + math.Cos(c.Latitude.Float64())*math.Cos(c2.Latitude.Float64())*sinLng*sinLng

	return Distance(2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a)))
}
