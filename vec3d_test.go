package geo_test

import (
	"math"
	"testing"

	geo "github.com/EpicStep/go-simple-geo"
)

func TestDistance(t *testing.T) {
	type test struct {
		v    *geo.Vec3d
		want float64
	}

	tests := []test{
		{
			v: &geo.Vec3d{
				X: 0,
				Y: 0,
				Z: 0,
			},
			want: 0,
		},
		{
			v: &geo.Vec3d{
				X: 1,
				Y: 0,
				Z: 0,
			},
			want: 1,
		},
		{
			v: &geo.Vec3d{
				X: 0,
				Y: 1,
				Z: 1,
			},
			want: 2,
		},
		{
			v: &geo.Vec3d{
				X: 1,
				Y: 1,
				Z: 1,
			},
			want: 3,
		},
		{
			v: &geo.Vec3d{
				X: 1,
				Y: 1,
				Z: 2,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		distance := tests[0].v.SquareDistance(tt.v)

		if distance != tt.want {
			t.Fatalf("incorrect answer for x:%2f, y:%2f, z:%2f. Want: %2f", tt.v.X, tt.v.Y, tt.v.Z, tt.want)
		}
	}
}

func TestGeo_ToVec3d(t *testing.T) {
	origin := geo.Vec3d{
		X: 0,
		Y: 0,
		Z: 0,
	}

	c1 := geo.Coordinates{
		Latitude:  0,
		Longitude: 0,
	}

	c2 := geo.Coordinates{
		Latitude:  math.Pi / 2,
		Longitude: 0,
	}

	c3 := geo.Coordinates{
		Latitude:  math.Pi,
		Longitude: 0,
	}

	p1Vec := c1.ToVec3d()

	if p1 := origin.SquareDistance(p1Vec); p1 != 1 {
		t.Fatalf("answer is incorrect. Want: %d, got: %2f", 1, p1)
	}

	p2Dist := p1Vec.SquareDistance(c2.ToVec3d())

	if p2Dist >= 2 && p2Dist <= 1.9 {
		t.Fatalf("answer is incorrect. Want: %d, got: %2f", 1, p2Dist)
	}

	if p3 := p1Vec.SquareDistance(c3.ToVec3d()); p3 != 4 {
		t.Fatalf("answer is incorrect. Want: %d, got: %2f", 1, p3)
	}
}
