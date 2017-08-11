package mapper

type Map struct {
	ID     string `json:"id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Road   []Line `json:"road"`
}

type Line struct {
	StartX int `json:"startX"`
	StartY int `json:"startY"`
	EndX   int `json:"endX"`
	EndY   int `json:"endY"`
}

var (
	map1 = Map{
		ID:     "1",
		Height: 500,
		Width:  1500,
		Road: []Line{
			{0, 200, 1500, 200},
			{0, 300, 1500, 300},
		},
	}
)

func GetMap() (Map, error) {
	return map1, nil
}
