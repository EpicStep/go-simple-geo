package geo

import (
	"math"
)

// RadiansToDegrees convert radians to degrees.
func RadiansToDegrees(rads float64) float64 {
	return rads * (unfoldedAngle / math.Pi)
}
