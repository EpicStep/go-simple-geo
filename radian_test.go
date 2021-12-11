package geo_test

import (
	"fmt"
	"testing"

	geo "github.com/EpicStep/go-simple-geo"
)

func TestRadian_ToDegrees(t *testing.T) {
	type test struct {
		v    geo.Radian
		want string
	}

	tests := []test{
		{
			v:    geo.Radian(1),
			want: "57.30",
		},
		{
			v:    geo.Radian(0.5),
			want: "28.65",
		},
	}

	for _, tt := range tests {
		if x := fmt.Sprintf("%.2f", tt.v.ToDegrees().Float64()); x != tt.want {
			t.Fatalf("incorrect answer for %.2f. Want: %s. Got: %s", tt.v.Float64(), tt.want, x)
		}
	}
}
