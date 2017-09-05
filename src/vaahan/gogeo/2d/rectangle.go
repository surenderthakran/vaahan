package gogeo

import (
	"fmt"
)

type Rectangle struct {
	Height      float64 `json:"height"`
	Width       float64 `json:"width"`
	TopLeft     *Point  `json:"top_left"`
	TopRight    *Point  `json:"top_right"`
	BottomRight *Point  `json:"bottom_right"`
	BottomLeft  *Point  `json:"bottom_left"`
	sides       []*LineSegment
}

func (rect *Rectangle) Sides() ([]*LineSegment, error) {
	sw := NewPoint(rect.TopLeft.X, rect.BottomRight.Y)
	ne := NewPoint(rect.BottomRight.X, rect.TopLeft.Y)

	swNWSide, err := NewLineSegmentByPoints(sw, rect.TopLeft)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	nwNESide, err := NewLineSegmentByPoints(rect.TopLeft, ne)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	neSESide, err := NewLineSegmentByPoints(ne, rect.BottomRight)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	seSESide, err := NewLineSegmentByPoints(rect.BottomRight, sw)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}

	return []*LineSegment{
		swNWSide,
		nwNESide,
		neSESide,
		seSESide,
	}, nil
}

func (rect *Rectangle) ContainsPoint(point *Point) bool {
	var ccwResult []bool

	ccwResult = append(ccwResult, counterClockWise(rect.TopLeft, rect.BottomLeft, point))
	ccwResult = append(ccwResult, counterClockWise(rect.BottomLeft, rect.BottomRight, point))
	ccwResult = append(ccwResult, counterClockWise(rect.BottomRight, rect.TopRight, point))
	ccwResult = append(ccwResult, counterClockWise(rect.TopRight, rect.TopLeft, point))

	return ccwResult[0] == ccwResult[1] && ccwResult[1] == ccwResult[2] && ccwResult[2] == ccwResult[3] && ccwResult[0] == ccwResult[3]
}

func (rect *Rectangle) Equal(rect2 *Rectangle) bool {
	return rect.Height == rect2.Height && rect.Width == rect2.Width && rect.TopLeft.Equal(rect2.TopLeft) && rect.BottomRight.Equal(rect2.BottomRight)
}

func (rect *Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Height: %v, Width: %v, TopLeft: %v, BottomRight: %v}", rect.Height, rect.Width, rect.TopLeft, rect.BottomRight)
}

// TODO(surenderthakran): Fix for tilted rectangles.
func NewRectangleByCorners(sw, ne *Point) *Rectangle {
	nw := NewPoint(sw.X, ne.Y)
	se := NewPoint(ne.X, sw.Y)

	rectangle := &Rectangle{
		Height:      nw.DistanceFrom(sw),
		Width:       nw.DistanceFrom(ne),
		TopLeft:     nw,
		TopRight:    ne,
		BottomRight: se,
		BottomLeft:  sw,
	}
	return rectangle
}
