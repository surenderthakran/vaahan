package gogeo

import (
	"fmt"
)

type Rectangle struct {
	Height      float64 `json:"height"`
	Width       float64 `json:"width"`
	TopLeft     *Point  `json:"top_left"`
	BottomRight *Point  `json:"bottom_right"`
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
	d1 := rect.TopLeft.DistanceFrom(point)
	d2 := rect.BottomRight.DistanceFrom(point)
	return d1+d2 < rect.Height+rect.Width
}

func (rect *Rectangle) Equal(rect2 *Rectangle) bool {
	return rect.Height == rect2.Height && rect.Width == rect2.Width && rect.TopLeft.Equal(rect2.TopLeft) && rect.BottomRight.Equal(rect2.BottomRight)
}

func (rect *Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Height: %v, Width: %v, TopLeft: %v, BottomRight: %v}", rect.Height, rect.Width, rect.TopLeft, rect.BottomRight)
}

func NewRectangleByCorners(sw, ne *Point) *Rectangle {
	nw := NewPoint(sw.X, ne.Y)
	se := NewPoint(ne.X, sw.Y)

	rectangle := &Rectangle{
		Height:      nw.DistanceFrom(sw),
		Width:       nw.DistanceFrom(ne),
		TopLeft:     nw,
		BottomRight: se,
	}
	return rectangle
}
