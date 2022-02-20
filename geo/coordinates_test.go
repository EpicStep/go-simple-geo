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
		{
			v:    geo.NewCoordinatesFromDegrees(40.7143, -74.006),
			v2:   geo.NewCoordinatesFromDegrees(37.7749, -122.419),
			want: "4129.02",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.Distance(tt.v2)); x != tt.want {
			t.Fatalf("incorrect answer for (%.2f %.2f) and (%.2f %.2f). Want: %s. Got: %s", tt.v.Latitude, tt.v.Longitude, tt.v2.Latitude, tt.v.Longitude, tt.want, x)
		}
	}
}

func TestNewCoordinatesFromRadians(t *testing.T) {
	c := geo.NewCoordinatesFromRadians(0.1, 0.1)

	if c.Latitude != 0.1 || c.Longitude != 0.1 {
		t.Fail()
	}
}

func BenchmarkCoordinates_Distance(b *testing.B) {
	d := geo.NewCoordinatesFromDegrees(55.9606, 38.0456)
	d2 := geo.NewCoordinatesFromDegrees(59.9386, 30.3141)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if dist := d.Distance(d2); dist != 634.6971900186131 {
			b.Fail()
		}
	}
}
