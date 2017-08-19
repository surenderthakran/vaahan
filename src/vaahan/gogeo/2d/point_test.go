package gogeo

import (
	"testing"
)

func TestDistanceFrom(t *testing.T) {
	testCases := []struct {
		start *Point
		end   *Point
		want  float64
	}{
		{
			start: NewPoint(0, 0),
			end:   NewPoint(0, 5),
			want:  float64(5),
		},
		// {
		// 	start: NewPoint(1, 3),
		// 	end:   NewPoint(5, 10),
		// 	want:  float64(5),
		// },
	}

	for _, test := range testCases {
		distance := test.start.DistanceFrom(test.end)
		if distance != test.want {
			t.Errorf("%v.DistanceFrom(%v) want: %v got: %v", test.start, test.end, test.want, distance)
		}
	}
}
