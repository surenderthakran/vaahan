package gogeo

import (
	"math"
)

type Angle float64

const (
	Degree = (math.Pi / 180)
)

func (angle Angle) Degrees() float64 {
	return float64(angle / Degree)
}

func (angle Angle) Radians() float64 {
	return float64(angle)
}
