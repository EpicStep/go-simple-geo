package geo_test

import (
	"fmt"
	"testing"

	"github.com/EpicStep/go-simple-geo/geo"
)

func TestCoordinates_Distance(t *testing.T) {
	type test struct {
		v    *geo.Coordinates
		v2   *geo.Coordinates
		want string
	}

	tests := []test{
		{
			v:    geo.NewCoordinatesFromDegrees(55.9606, 38.0456),
			v2:   geo.NewCoordinatesFromDegrees(59.9386, 30.3141),
			want: "634.70",
		},
		{
			v:    geo.NewCoordinatesFromDegrees(55.7522, 37.6156),
			v2:   geo.NewCoordinatesFromDegrees(54.7818, 32.0401),
			want: "369.22",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.Distance(tt.v2).ToKilometers()); x != tt.want {
			t.Fatalf("incorrect answer for (%.2f %.2f) and (%.2f %.2f). Want: %s. Got: %s", tt.v.Latitude, tt.v.Longitude, tt.v2.Latitude, tt.v.Longitude, tt.want, x)
		}
	}
}

func TestDistance_ToKilometers(t *testing.T) {
	type test struct {
		v    geo.Distance
		want string
	}

	tests := []test{
		{
			v:    geo.Distance(0.1),
			want: "637.10",
		},
		{
			v:    geo.Distance(1),
			want: "6371.00",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.ToKilometers()); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", float64(tt.v), tt.want, x)
		}
	}
}

func TestDistance_ToMeters(t *testing.T) {
	type test struct {
		v    geo.Distance
		want string
	}

	tests := []test{
		{
			v:    geo.Distance(0.1),
			want: "637100.00",
		},
		{
			v:    geo.Distance(1),
			want: "6371000.00",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.ToMeters()); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", float64(tt.v), tt.want, x)
		}
	}
}

func TestNewCoordinatesFromRadians(t *testing.T) {
	c := geo.NewCoordinatesFromRadians(geo.Radian(0.1), geo.Radian(0.1))

	if c.Latitude != geo.Radian(0.1) || c.Longitude != geo.Radian(0.1) {
		t.Fail()
	}
}

func BenchmarkCoordinates_Distance(b *testing.B) {
	d := geo.NewCoordinatesFromDegrees(55.9606, 38.0456)
	d2 := geo.NewCoordinatesFromDegrees(59.9386, 30.3141)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if dist := d.Distance(d2).ToKilometers(); dist != 634.6971900186131 {
			b.Fail()
		}
	}
}
