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
	Start *Point `json:"start_point"`
	angle Angle
}

func NewRayByPointAndDirection(start *Point, angle Angle) (*Ray, error) {
	if !start.Valid() {
		return nil, fmt.Errorf("unable to create ray: starting point must be valid.")
	}
	return &Ray{
		Start: start,
		angle: NormalizeRadian(angle),
	}, nil
}

func (ray *Ray) Angle() Angle {
	return ray.angle
}

func (ray *Ray) SetAngle(angle Angle) {
	ray.angle = NormalizeRadian(angle)
}

func (ray *Ray) FindPointAtDistance(distance float64) *Point {
	x := math.Cos(ray.angle.Radians()) * distance
	y := math.Sin(ray.angle.Radians()) * distance
	point := NewPoint(ray.Start.X+x, ray.Start.Y+y)
	point.RoundTo(2)
	return point
}

func (ray *Ray) Intersection(segment *LineSegment) *Point {
	rayOrigin := ray.Start.Vector()
	rayDirection := NewVectorFromAngle(ray.Angle())
	point1 := segment.StartPoint().Vector()
	point2 := segment.EndPoint().Vector()

	v1 := rayOrigin.SubtractVector(point1)
	v2 := point2.SubtractVector(point1)
	v3 := NewVector(-rayDirection.Y, rayDirection.X)

	dot := v2.DotProduct(v3)
	if dot == 0 {
		return nil
	}

	t1 := v2.CrossProduct(v1) / dot
	t2 := v1.DotProduct(v3) / dot

	if t1 >= 0 && t2 >= 0 && t2 <= 1 {
		intersection := rayOrigin.AddVector(rayDirection.Multiply(t1))
		return intersection.Point().RoundTo(2)
	}

	return nil
}

func (ray *Ray) String() string {
	return fmt.Sprintf("Ray{%v, %v}", ray.Start, ray.Angle())
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
	AB := segment.StartPoint().DistanceFrom(point)
	BC := point.DistanceFrom(segment.EndPoint())
	AC := segment.StartPoint().DistanceFrom(segment.EndPoint())
	return Equal(AB+BC, AC)
}

func (segment *LineSegment) MidPoint() *Point {
	return NewPoint((segment.start.X+segment.end.X)/2, (segment.start.Y+segment.end.Y)/2)
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
	yIntercept := GetYInterceptByPointAndSlope(start, slope)
	if math.IsNaN(yIntercept) {
		yIntercept = 0
	}
	return slope, yIntercept
}

func GetYInterceptByPointAndSlope(point *Point, slope float64) float64 {
	return point.Y - (slope * point.X)
}
