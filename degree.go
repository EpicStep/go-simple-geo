package geo

import "math"

// Degree type.
type Degree float64

// ToRadians convert degrees to radians
func (d Degree) ToRadians() Radian {
	return Radian(float64(d) * (math.Pi / angle))
}

// Float64 return float64 value of Degree type.
func (d Degree) Float64() float64 {
	return float64(d)
}
