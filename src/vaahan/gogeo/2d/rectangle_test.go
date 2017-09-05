package gogeo

import (
	"testing"
)

func TestContainsPoint(t *testing.T) {
	testCases := []struct {
		rectangle *Rectangle
		point     *Point
		want      bool
	}{
		{
			rectangle: &Rectangle{
				Height:      11.2,
				Width:       39.4,
				TopLeft:     NewPoint(17, 35),
				TopRight:    NewPoint(51, 15),
				BottomRight: NewPoint(46, 6),
				BottomLeft:  NewPoint(25, 25),
			},
			point: NewPoint(30, 13),
			want:  false,
		},
		{
			rectangle: &Rectangle{
				Height:      11.2,
				Width:       39.4,
				TopLeft:     NewPoint(17, 35),
				TopRight:    NewPoint(51, 15),
				BottomRight: NewPoint(46, 6),
				BottomLeft:  NewPoint(12, 25),
			},
			point: NewPoint(50, 10),
			want:  false,
		},
		{
			rectangle: &Rectangle{
				Height:      11.2,
				Width:       39.4,
				TopLeft:     NewPoint(17, 35),
				TopRight:    NewPoint(51, 15),
				BottomRight: NewPoint(46, 6),
				BottomLeft:  NewPoint(12, 25),
			},
			point: NewPoint(15, 30),
			want:  true,
		},
	}

	for _, test := range testCases {
		result := test.rectangle.ContainsPoint(test.point)
		if result != test.want {
			t.Errorf("%v.ContainsPoint(%v) want: %v got: %v", test.rectangle, test.point, test.want, result)
		}
	}
}

// func TestNewRectangleByCorners(t *testing.T) {
// 	testCases := []struct {
// 		sw   *Point
// 		ne   *Point
// 		want *Rectangle
// 	}{
// 		{
// 			sw: NewPoint(12, 25),
// 			ne: NewPoint(51, 15),
// 			want: &Rectangle{
// 				Height:      11.2,
// 				Width:       39.4,
// 				TopLeft:     NewPoint(17, 35),
// 				BottomRight: NewPoint(46, 6),
// 			},
// 		},
// 	}
//
// 	for _, test := range testCases {
// 		rectangle := NewRectangleByCorners(test.sw, test.ne)
// 		if !rectangle.Equal(test.want) {
// 			t.Errorf("NewRectangleByCorners(%v, %v) want: %v got: %v", test.sw, test.ne, test.want, rectangle)
// 		}
// 	}
// }
