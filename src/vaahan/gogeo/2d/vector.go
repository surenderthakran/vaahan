package gogeo

import (
	"math"
)

type Vector Point

func NewVector(x, y float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
	}
}

func NewVectorFromAngle(angle Angle) *Vector {
	theta := float64(angle)
	return NewVector(math.Cos(theta), math.Sin(theta))
}

func (v *Vector) Add(value float64) *Vector {
	return NewVector(v.X+value, v.Y+value)
}

func (v *Vector) Subtract(value float64) *Vector {
	return NewVector(v.X-value, v.Y-value)
}

func (v *Vector) Multiply(value float64) *Vector {
	return NewVector(v.X*value, v.Y*value)
}

func (v1 *Vector) AddVector(v2 *Vector) *Vector {
	return NewVector(v1.X+v2.X, v1.Y+v2.Y)
}

func (v1 *Vector) SubtractVector(v2 *Vector) *Vector {
	return NewVector(v1.X-v2.X, v1.Y-v2.Y)
}

func (v1 *Vector) DotProduct(v2 *Vector) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y)
}

func (v1 *Vector) CrossProduct(v2 *Vector) float64 {
	return (v1.X * v2.Y) - (v1.Y * v2.X)
}

func (v *Vector) Point() *Point {
	return NewPoint(v.X, v.Y)
}
