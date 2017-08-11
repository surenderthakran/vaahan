package car

type Car struct {
	Length int `json:"length"`
	Width  int `json:"width"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

func New() Car {
	return Car{50, 30, 0, 0}
}
