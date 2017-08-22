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
			start: &Point{0, 0},
			end:   &Point{0, 5},
			want:  float64(5),
		},
		{
			start: &Point{1, 3},
			end:   &Point{5, 10},
			want:  float64(8.06),
		},
	}

	for _, test := range testCases {
		distance := test.start.DistanceFrom(test.end)
		if distance != test.want {
			t.Errorf("%v.DistanceFrom(%v) want: %v got: %v", test.start, test.end, test.want, distance)
		}
	}
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		p1   *Point
		p2   *Point
		want bool
	}{
		{
			p1:   &Point{0, 0},
			p2:   &Point{0, 5},
			want: false,
		},
		{
			p1:   &Point{1, 3},
			p2:   &Point{1, 3},
			want: true,
		},
	}

	for _, test := range testCases {
		result := test.p1.Equal(test.p2)
		if result != test.want {
			t.Errorf("%v.Equal(%v) want: %v got: %v", test.p1, test.p2, test.want, result)
		}
	}
}
