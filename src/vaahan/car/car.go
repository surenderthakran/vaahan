package car

import (
	"fmt"
	"math"
	"time"

	geo "vaahan/gogeo/2d"
	"vaahan/track"

	glog "github.com/golang/glog"
)

type Car struct {
	Length         float64
	Width          float64
	frontCenter    *geo.Point
	backCenter     *geo.Point
	LeftHeadlight  *geo.Point `json:"left_headlight"`
	RightHeadlight *geo.Point `json:"right_headlight"`
	LeftTaillight  *geo.Point `json:"left_taillight"`
	RightTaillight *geo.Point `json:"right_taillight"`
	Status         string     `json:"status"`
}

var (
	car *Car
)

func (car *Car) moveForward(units float64) {
	glog.Info("inside car.moveForward()")
	if units == 0 {
		units = 1
	}
	car.LeftHeadlight.X = car.LeftHeadlight.X + units
	// car.LeftHeadlight.Y = car.LeftHeadlight.Y + units

	car.RightHeadlight.X = car.RightHeadlight.X + units
	// car.RightHeadlight.Y = car.RightHeadlight.Y + units

	car.LeftTaillight.X = car.LeftTaillight.X + units
	// car.LeftTaillight.Y = car.LeftTaillight.Y + units

	car.RightTaillight.X = car.RightTaillight.X + units
	// car.RightTaillight.Y = car.RightTaillight.Y + units
}

func (car *Car) driver() {
	glog.Info("inside car.driver()")
	// instruction := ""
	for {
		glog.Info("reading from channel")
		if car.Status == "driving" {
			car.moveForward(0)
		}
		time.Sleep(time.Second * 1)
	}
}

func (car *Car) Drive() {
	car.Status = "driving"
}

func (car *Car) Pause() {
	car.Status = "paused"
}

func GetCar() (*Car, error) {
	if car == nil {
		glog.Error("car not found")
		return nil, fmt.Errorf("car not found")
	}
	return car, nil
}

func (car *Car) SetRightHeadlight(startVector *geo.Ray) {
	tan := car.Width / (2 * car.Length)
	alpha := math.Tanh(tan)
	distanceOfHeadlightFromStart := car.Length / math.Cos(alpha)
	theta := startVector.GetAngle().Radians() - alpha

	x := math.Cos(theta) * distanceOfHeadlightFromStart
	y := math.Sin(theta) * distanceOfHeadlightFromStart

	rightHeadlight := geo.NewPoint(geo.RoundTo(startVector.GetStartPoint().X+x, 2), geo.RoundTo(startVector.GetStartPoint().Y+y, 2))

	car.RightHeadlight = rightHeadlight
}

func New(track *track.Track) *Car {
	startVector := track.StartVector

	length := float64(50)
	width := float64(30)
	car = &Car{
		Length:         length,
		Width:          width,
		frontCenter:    startVector.FindPointAtDistance(length),
		backCenter:     startVector.GetStartPoint(),
		LeftHeadlight:  geo.NewPoint(50, 265),
		LeftTaillight:  geo.NewPoint(0, 265),
		RightTaillight: geo.NewPoint(0, 235),
		Status:         "stopped",
	}

	car.SetRightHeadlight(startVector)

	go car.driver()
	return car
}
