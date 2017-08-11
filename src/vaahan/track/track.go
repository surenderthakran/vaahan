package track

type Track struct {
	ID            string `json:"id"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	StartingLine  Line   `json:"startingLine"`
	FinishingLine Line   `json:"finishingLine"`
	Road          []Line `json:"road"`
}

type Line struct {
	StartX int `json:"startX"`
	StartY int `json:"startY"`
	EndX   int `json:"endX"`
	EndY   int `json:"endY"`
}

var (
	track1 = Track{
		ID:            "1",
		Height:        500,
		Width:         1500,
		StartingLine:  Line{50, 200, 50, 300},
		FinishingLine: Line{1450, 200, 1450, 300},
		Road: []Line{
			{0, 200, 0, 300},
			{0, 300, 1500, 300},
			{1500, 300, 1500, 200},
			{1500, 200, 0, 200},
		},
	}
)

func GetTrack() (Track, error) {
	return track1, nil
}
