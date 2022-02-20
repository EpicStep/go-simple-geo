package geo_test

import (
	"fmt"
	"testing"

	"github.com/EpicStep/go-simple-geo/v2/geo"
)

func TestRadian_ToDegrees(t *testing.T) {
	type test struct {
		v    float64
		want string
	}

	tests := []test{
		{
			v:    1,
			want: "57.30",
		},
		{
			v:    0.5,
			want: "28.65",
		},
		{
			v:    4,
			want: "229.18",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", geo.RadiansToDegrees(tt.v)); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", tt.v, tt.want, x)
		}
	}
}
