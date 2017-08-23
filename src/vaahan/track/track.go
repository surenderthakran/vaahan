package track

import (
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
	boundary := geo.NewRectangleByCorners(origin, oppositeOrigin)
	track := &Track{
		ID:       "1",
		Height:   500,
		Width:    1000,
		Boundary: boundary,
		// StartVector: geo.NewRayByPointAndDirection(&geo.Point{20, 250}, geo.Angle(0)),
		StartVector: geo.NewRayByPointAndDirection(&geo.Point{500, 250}, geo.Angle(math.Pi/10)),
	}
	return track, nil
}
