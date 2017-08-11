package track

type Track struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

func GetTrack() (Track, error) {
	return Track{1001, 2002}, nil
}
