package geo

import (
	"math"
)

// Vec3d is a 3D floating point structure.
type Vec3d struct {
	X, Y, Z float64
}

// SquareDistance return a square of the distance between two points.
func (vec *Vec3d) SquareDistance(v *Vec3d) float64 {
	return math.Pow(vec.X-v.X, 2) + math.Pow(vec.Y-v.Y, 2) + math.Pow(vec.Z-v.Z, 2)
}

// ToVec3d return Vec3d from Coordinates.
func (c *Coordinates) ToVec3d() *Vec3d {
	r := math.Cos(c.Latitude.Float64())

	return &Vec3d{
		X: math.Cos(c.Longitude.Float64()) * r,
		Y: math.Sin(c.Longitude.Float64()) * r,
		Z: math.Sin(c.Latitude.Float64()),
	}
}
