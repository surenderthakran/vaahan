package mapper

type Map struct {
	ID     string `json:"id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

var (
	map1 = Map{
		ID:     "1",
		Height: 500,
		Width:  1500,
	}
)

func GetMap() (Map, error) {
	return map1, nil
}
