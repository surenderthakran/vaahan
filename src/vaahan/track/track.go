package track

import (
	"fmt"
	"math"

	geo "vaahan/gogeo/2d"
)

type Track struct {
	ID          string         `json:"id"`
	Height      int            `json:"height"`
	Width       int            `json:"width"`
	Boundary    *geo.Rectangle `json:"boundary"`
	StartVector *geo.Ray       `json:"start_vector"`
}

var (
	origin         = geo.NewPoint(0, 0)
	oppositeOrigin = geo.NewPoint(1000, 500)
)

func GetTrack(trackID string) (*Track, error) {
	boundary, err := geo.NewRectangleByCorners(origin, oppositeOrigin)
	if err != nil {
		return nil, fmt.Errorf("unable to define track boundary: %s.", err)
	}
	startVector, err := geo.NewRayByPointAndDirection(&geo.Point{100, 250}, geo.Angle(math.Pi/10))
	if err != nil {
		return nil, fmt.Errorf("unable to define track's starting vector: %s.", err)
	}
	track := &Track{
		ID:          "1",
		Height:      500,
		Width:       1000,
		Boundary:    boundary,
		StartVector: startVector,
	}
	return track, nil
}
