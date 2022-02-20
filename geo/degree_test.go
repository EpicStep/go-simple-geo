package geo_test

import (
	"fmt"
	"testing"

	"github.com/EpicStep/go-simple-geo/v2/geo"
)

func TestDegree_ToRadians(t *testing.T) {
	type test struct {
		v    float64
		want string
	}

	tests := []test{
		{
			v:    100,
			want: "1.75",
		},
		{
			v:    57.3,
			want: "1.00",
		},
		{
			v:    1,
			want: "0.02",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", geo.DegreeToRadians(tt.v)); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", tt.v, tt.want, x)
		}
	}
}
