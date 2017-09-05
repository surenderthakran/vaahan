package gogeo

import (
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundTo(input float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(input*output)) / output
}

func counterClockWise(a, b, c *Point) bool {
	return (b.X-a.X)*(c.Y-a.Y) > (b.Y-a.Y)*(c.X-a.X)
}
