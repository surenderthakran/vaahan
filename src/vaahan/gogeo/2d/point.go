package gogeo

import (
	"math"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p1 *Point) DistanceFrom(p2 *Point) float64 {
	distance := math.Pow(math.Pow(p2.X-p1.X, 2)+math.Pow(p2.Y-p1.Y, 2), 0.5)
	return RoundTo(distance, 2)
}

func (p *Point) RoundTo(precision int) *Point {
	p.X = RoundTo(p.X, precision)
	p.Y = RoundTo(p.Y, precision)
	return p
}

func (p1 *Point) Equal(p2 *Point) bool {
	if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
		return false
	}
	if p1 == nil && p2 == nil {
		return true
	}
	return Equal(p1.X, p2.X) && Equal(p1.Y, p2.Y)
}

func (p *Point) Valid() bool {
	if math.IsNaN(p.X) || math.IsNaN(p.Y) || math.IsInf(p.X, 0) || math.IsInf(p.Y, 0) {
		return false
	}
	return true
}

func (p *Point) Vector() *Vector {
	return NewVector(p.X, p.Y)
}
