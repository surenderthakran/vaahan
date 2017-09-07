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

// counterClockWise checks if a path from a -> b -> c turns in counter clockwise direction.
func counterClockWise(a, b, c *Point) bool {
	return (b.X-a.X)*(c.Y-a.Y) > (b.Y-a.Y)*(c.X-a.X)
}

func percentage(value, percent float64) float64 {
	return (value * percent) / 100
}

func Equal(a, b float64) bool {
	a = math.Abs(a)
	b = math.Abs(b)
	difference := math.Abs(a - b)
	return difference <= percentage(a, 0.1) && difference <= percentage(b, 0.1)
}

func NormalizeRadian(angle Angle) Angle {
	angle = Angle(RoundTo(float64(angle), 2))
	for angle > math.Pi {
		angle = angle - (2 * math.Pi)
		angle = Angle(RoundTo(float64(angle), 2))
	}
	for angle < -math.Pi {
		angle = angle + (2 * math.Pi)
		angle = Angle(RoundTo(float64(angle), 2))
	}
	return angle
}
