package gogeo

type Line struct {
	start, end *Point
	slope      float64
	yIntercept float64
}

func (line Line) GetStartPoint() *Point {
	return line.start
}

func (line Line) GetSlope() float64 {
	return line.slope
}

func (line Line) GetYIntercept() float64 {
	return line.yIntercept
}

func GetSlopeAndYInterceptByPoints(start, end *Point) (float64, float64) {
	slope := (end.Y - start.Y) / (end.X - end.X)
	yIntercept := start.Y - (slope * start.X)
	return slope, yIntercept
}

func GetYInterceptByPointAndSlope(point *Point, slope float64) float64 {
	return point.Y - (slope * point.X)
}

func NewLineSegmentByPoints(start, end *Point) *Line {
	line := Line{
		start: start,
		end:   end,
	}
	slope, yIntercept := GetSlopeAndYInterceptByPoints(start, end)
	line.slope = slope
	line.yIntercept = yIntercept
	return &line
}

func NewRayByPointAndEquation(start *Point, slope, yIntercept float64) *Line {
	return &Line{
		start:      start,
		slope:      slope,
		yIntercept: yIntercept,
	}
}
