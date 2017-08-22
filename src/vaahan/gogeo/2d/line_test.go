package gogeo

import (
	"math"
	"testing"
)

func TestRayFindPointAtDistance(t *testing.T) {
	testCases := []struct {
		ray      *Ray
		distance float64
		want     *Point
	}{
		{
			ray: &Ray{
				start: &Point{0, 0},
				angle: 0,
			},
			distance: float64(5),
			want:     NewPoint(5, 0),
		},
		{
			ray: &Ray{
				start: &Point{0, 0},
				angle: math.Pi / 2,
			},
			distance: float64(5),
			want:     NewPoint(0, 5),
		},
		{
			ray: &Ray{
				start: &Point{0, 0},
				angle: math.Pi / 4,
			},
			distance: float64(5),
			want:     NewPoint(3.54, 3.54),
		},
	}

	for _, test := range testCases {
		point := test.ray.FindPointAtDistance(test.distance)
		if !point.Equal(test.want) {
			t.Errorf("%v.FindPointAtDistance(%v) want: %v got: %v", test.ray, test.distance, test.want, point)
		}
	}
}
