package gogeo

type Line struct {
	slope      float64
	yIntercept float64
}

func (line Line) GetSlope() float64 {
	return line.slope
}

func (line Line) GetYIntercept() float64 {
	return line.yIntercept
}

type Ray struct {
	start      *Point
	slope      float64
	yIntercept float64
}

func NewRayByPointAndEquation(start *Point, slope, yIntercept float64) *Ray {
	return &Ray{
		start:      start,
		slope:      slope,
		yIntercept: yIntercept,
	}
}

func (ray Ray) GetStartPoint() *Point {
	return ray.start
}

func (ray Ray) GetSlope() float64 {
	return ray.slope
}

func (ray Ray) GetYIntercept() float64 {
	return ray.yIntercept
}

type LineSegment struct {
	start, end *Point
	slope      float64
	yIntercept float64
}

func (segment LineSegment) GetStartPoint() *Point {
	return segment.start
}

func (segment LineSegment) GetSlope() float64 {
	return segment.slope
}

func (segment LineSegment) GetYIntercept() float64 {
	return segment.yIntercept
}

func NewLineSegmentByPoints(start, end *Point) *LineSegment {
	segment := LineSegment{
		start: start,
		end:   end,
	}
	slope, yIntercept := GetSlopeAndYInterceptByPoints(start, end)
	segment.slope = slope
	segment.yIntercept = yIntercept
	return &segment
}

func GetSlopeAndYInterceptByPoints(start, end *Point) (float64, float64) {
	slope := (end.Y - start.Y) / (end.X - end.X)
	yIntercept := start.Y - (slope * start.X)
	return slope, yIntercept
}

func GetYInterceptByPointAndSlope(point *Point, slope float64) float64 {
	return point.Y - (slope * point.X)
}
