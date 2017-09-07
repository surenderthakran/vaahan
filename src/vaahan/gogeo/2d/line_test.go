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
				Start: &Point{0, 0},
				angle: 0,
			},
			distance: float64(5),
			want:     NewPoint(5, 0),
		},
		{
			ray: &Ray{
				Start: &Point{0, 0},
				angle: math.Pi / 2,
			},
			distance: float64(5),
			want:     NewPoint(0, 5),
		},
		{
			ray: &Ray{
				Start: &Point{0, 0},
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

func TestLineSegmentHasPoint(t *testing.T) {
	testCases := []struct {
		segment *LineSegment
		point   *Point
		want    bool
	}{
		{
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{0, 500},
			},
			point: &Point{0, 100},
			want:  true,
		},
		{
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{0, -500},
			},
			point: &Point{0, 100},
			want:  false,
		},
		{
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{500, 500},
			},
			point: &Point{100, 100},
			want:  true,
		},
		{
			segment: &LineSegment{
				start: &Point{10, 20},
				end:   &Point{50, 10},
			},
			point: &Point{30, 15},
			want:  true,
		},
		{
			segment: &LineSegment{
				start: &Point{30, 24},
				end:   &Point{40, 8},
			},
			point: &Point{35, 16},
			want:  true,
		},
	}

	for _, test := range testCases {
		hasPoint := test.segment.HasPoint(test.point)
		if hasPoint != test.want {
			t.Errorf("%v.HasPoint(%v) want: %v got: %v", test.segment, test.point, test.want, hasPoint)
		}
	}
}

func TestLineSegmentMidPoint(t *testing.T) {
	testCases := []struct {
		segment *LineSegment
		want    *Point
	}{
		{
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{0, 500},
			},
			want: &Point{0, 250},
		},
		{
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{500, 500},
			},
			want: &Point{250, 250},
		},
		{
			segment: &LineSegment{
				start: &Point{10, 20},
				end:   &Point{50, 10},
			},
			want: &Point{30, 15},
		},
		{
			segment: &LineSegment{
				start: &Point{30, 24},
				end:   &Point{40, 8},
			},
			want: &Point{35, 16},
		},
		{
			segment: &LineSegment{
				start: &Point{30, 20},
				end:   &Point{-8, -14},
			},
			want: &Point{11, 3},
		},
	}

	for _, test := range testCases {
		point := test.segment.MidPoint()
		if !point.Equal(test.want) {
			t.Errorf("%v.MidPoint() want: %v got: %v", test.segment, test.want, point)
		}
	}
}

func TestRayLineSegmentIntersection(t *testing.T) {
	testCases := []struct {
		ray     *Ray
		segment *LineSegment
		want    *Point
	}{
		{
			ray: &Ray{
				Start: &Point{0, 15},
				angle: 0,
			},
			segment: &LineSegment{
				start: &Point{30, 25},
				end:   &Point{30, 5},
			},
			want: &Point{30, 15},
		},
		{
			ray: &Ray{
				Start: &Point{0, 15},
				angle: 0,
			},
			segment: &LineSegment{
				start: &Point{-30, 25},
				end:   &Point{-30, 5},
			},
			want: nil,
		},
		{
			ray: &Ray{
				Start: &Point{0, 15},
				angle: 0,
			},
			segment: &LineSegment{
				start: &Point{30, 25},
				end:   &Point{15, 5},
			},
			want: &Point{22.5, 15},
		},
		{
			ray: &Ray{
				Start: &Point{5, 0},
				angle: math.Pi / 2,
			},
			segment: &LineSegment{
				start: &Point{0, 0},
				end:   &Point{10, 10},
			},
			want: &Point{5, 5},
		},
	}

	for _, test := range testCases {
		point := test.ray.Intersection(test.segment)
		if !point.Equal(test.want) {
			t.Errorf("%v.Intersection(%v) want: %v got: %v", test.ray, test.segment, test.want, point)
		}
	}
}
