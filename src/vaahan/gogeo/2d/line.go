package gogeo

import (
	"fmt"
	"math"
)

type Line struct {
	slope      float64
	yIntercept float64
}

func (line *Line) Slope() float64 {
	return line.slope
}

func (line *Line) YIntercept() float64 {
	return line.yIntercept
}

type Ray struct {
	start *Point
	angle Angle
}

func NewRayByPointAndDirection(start *Point, angle Angle) (*Ray, error) {
	if !start.Valid() {
		return nil, fmt.Errorf("unable to create ray: starting point must be valid.")
	}
	return &Ray{
		start: start,
		angle: angle,
	}, nil
}

func (ray *Ray) StartPoint() *Point {
	return ray.start
}

func (ray *Ray) Angle() Angle {
	return ray.angle
}

func (ray *Ray) SetAngle(angle Angle) {
	ray.angle = angle
}

func (ray *Ray) FindPointAtDistance(distance float64) *Point {
	x := math.Cos(ray.angle.Radians()) * distance
	y := math.Sin(ray.angle.Radians()) * distance
	point := NewPoint(ray.start.X+x, ray.start.Y+y)
	point.RoundTo(2)
	return point
}

type LineSegment struct {
	start, end *Point
	slope      float64
	yIntercept float64
}

func (segment *LineSegment) StartPoint() *Point {
	return segment.start
}

func (segment *LineSegment) EndPoint() *Point {
	return segment.end
}

func (segment *LineSegment) Slope() float64 {
	return segment.slope
}

func (segment *LineSegment) YIntercept() float64 {
	return segment.yIntercept
}

func (segment *LineSegment) HasPoint(point *Point) bool {
	fmt.Println("inside segment.HasPoint()")
	fmt.Println(point)
	AB := segment.StartPoint().DistanceFrom(point)
	BC := point.DistanceFrom(segment.EndPoint())
	AC := segment.StartPoint().DistanceFrom(segment.EndPoint())
	if AB+BC == AC {
		return true
	}
	return false
}

func (segment *LineSegment) String() string {
	return fmt.Sprintf("LineSegment{%v, %v}", segment.StartPoint(), segment.EndPoint())
}

func NewLineSegmentByPoints(start, end *Point) (*LineSegment, error) {
	if start.Equal(end) {
		return nil, fmt.Errorf("unable to create line segment: starting and ending points cannot be same.")
	}
	if !start.Valid() || !end.Valid() {
		return nil, fmt.Errorf("unable to create line segment: starting and ending points must be valid.")
	}
	segment := LineSegment{
		start: start,
		end:   end,
	}
	slope, yIntercept := GetSlopeAndYInterceptByPoints(start, end)
	segment.slope = slope
	segment.yIntercept = yIntercept
	return &segment, nil
}

func GetSlopeAndYInterceptByPoints(start, end *Point) (float64, float64) {
	slope := (end.Y - start.Y) / (end.X - end.X)
	yIntercept := start.Y - (slope * start.X)
	if math.IsNaN(yIntercept) {
		yIntercept = 0
	}
	return slope, yIntercept
}

func GetYInterceptByPointAndSlope(point *Point, slope float64) float64 {
	return point.Y - (slope * point.X)
}
