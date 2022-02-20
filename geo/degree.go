package geo

import (
	"math"
)

// DegreeToRadians convert degrees to radians
func DegreeToRadians(deg float64) float64 {
	return deg * (math.Pi / unfoldedAngle)
}
