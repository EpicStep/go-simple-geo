package geo

import (
	"math"
)

// Radian type.
type Radian float64

// ToDegrees convert radians to degrees
func (r Radian) ToDegrees() Degree {
	return Degree(float64(r) * (unfoldedAngle / math.Pi))
}

// Float64 return float64 value of Radian type.
func (r Radian) Float64() float64 {
	return float64(r)
}
