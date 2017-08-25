package gogeo

import (
	"fmt"
)

type Rectangle struct {
	Height  float64 `json:"height"`
	Width   float64 `json:"width"`
	TopLeft *Point  `json:"top_left"`
	sides   []*LineSegment
}

func NewRectangleByCorners(sw, ne *Point) (*Rectangle, error) {
	nw := NewPoint(sw.X, ne.Y)
	se := NewPoint(ne.X, sw.Y)

	swNWSide, err := NewLineSegmentByPoints(sw, nw)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	nwNESide, err := NewLineSegmentByPoints(nw, ne)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	neSESide, err := NewLineSegmentByPoints(ne, se)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}
	seSESide, err := NewLineSegmentByPoints(se, sw)
	if err != nil {
		return nil, fmt.Errorf("unable to create rectangle: %s.", err)
	}

	rectangle := &Rectangle{
		Height:  nw.DistanceFrom(sw),
		Width:   nw.DistanceFrom(ne),
		TopLeft: nw,
		sides: []*LineSegment{
			swNWSide,
			nwNESide,
			neSESide,
			seSESide,
		},
	}
	return rectangle, nil
}

func (rect *Rectangle) Sides() []*LineSegment {
	return rect.sides
}
