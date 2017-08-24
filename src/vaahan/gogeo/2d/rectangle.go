package gogeo

type Rectangle struct {
	Height  float64 `json:"height"`
	Width   float64 `json:"width"`
	TopLeft *Point  `json:"top_left"`
	sides   []*LineSegment
}

func NewRectangleByCorners(sw, ne *Point) *Rectangle {
	nw := NewPoint(sw.X, ne.Y)
	se := NewPoint(ne.X, sw.Y)
	rectangle := &Rectangle{
		Height:  nw.DistanceFrom(sw),
		Width:   nw.DistanceFrom(ne),
		TopLeft: nw,
		sides: []*LineSegment{
			NewLineSegmentByPoints(sw, nw),
			NewLineSegmentByPoints(nw, ne),
			NewLineSegmentByPoints(ne, se),
			NewLineSegmentByPoints(se, sw),
		},
	}
	return rectangle
}

func (rect *Rectangle) Sides() []*LineSegment {
	return rect.sides
}
