package gogeo

import (
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

func NewRayByPointAndDirection(start *Point, angle Angle) *Ray {
	return &Ray{
		start: start,
		angle: angle,
	}
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
