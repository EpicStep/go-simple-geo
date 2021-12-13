package geo_test

import (
	"fmt"
	"testing"

	"github.com/EpicStep/go-simple-geo/geo"
)

func TestDegree_ToRadians(t *testing.T) {
	type test struct {
		v    geo.Degree
		want string
	}

	tests := []test{
		{
			v:    geo.Degree(100),
			want: "1.75",
		},
		{
			v:    geo.Degree(57.3),
			want: "1.00",
		},
		{
			v:    geo.Degree(1),
			want: "0.02",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.ToRadians().Float64()); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", tt.v.Float64(), tt.want, x)
		}
	}
}
