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
