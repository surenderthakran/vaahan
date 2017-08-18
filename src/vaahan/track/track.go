package track

import (
	"vaahan/shape"

	glog "github.com/golang/glog"
)

type Track struct {
	ID          string           `json:"id"`
	Height      int              `json:"height"`
	Width       int              `json:"width"`
	Boundary    *shape.Rectangle `json:"boundary"`
	StartVector *shape.Line      `json:"start_vector"`
}

var (
	origin         = shape.NewPoint(0, 0)
	oppositeOrigin = shape.NewPoint(1000, 500)
)

func GetTrack(trackID string) (*Track, error) {
	boundary := shape.NewRectangleByCorners(origin, oppositeOrigin)
	glog.Info(boundary)
	track := &Track{
		ID:          "1",
		Height:      500,
		Width:       1000,
		Boundary:    boundary,
		StartVector: shape.NewRayByPointAndEquation(&shape.Point{0, 250}, 0, 250),
	}
	return track, nil
}
