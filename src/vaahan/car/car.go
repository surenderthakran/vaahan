package car

import (
	"fmt"
	"math"
	"time"

	"vaahan/shape"
	"vaahan/track"

	glog "github.com/golang/glog"
)

type Car struct {
	LeftHeadlight  *shape.Point `json:"left_headlight"`
	RightHeadlight *shape.Point `json:"right_headlight"`
	LeftTaillight  *shape.Point `json:"left_taillight"`
	RightTaillight *shape.Point `json:"right_taillight"`
	Status         string       `json:"status"`
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

func New(track *track.Track) *Car {
	startVector := track.StartVector
	glog.Infof("startVector: %v", startVector)
	glog.Infof("slope: %v", startVector.GetSlope())
	glog.Infof("yIntercept: %v", startVector.GetYIntercept())
	glog.Info("tailCentre: %v", startVector.GetStartPoint())

	length := float64(50)
	interim := math.Sqrt(math.Pow(length, 2) / (1 + math.Pow(startVector.GetSlope(), 2)))

	x1 := startVector.GetStartPoint().X + interim
	y1 := (startVector.GetSlope() * x1) + startVector.GetYIntercept()
	glog.Infof("%v, %v", x1, y1)

	x2 := startVector.GetStartPoint().X - interim
	y2 := (startVector.GetSlope() * x2) + startVector.GetYIntercept()
	glog.Infof("%v, %v", x2, y2)

	dot1 := (startVector.GetStartPoint().X * x1) + (startVector.GetStartPoint().Y * y1)
	glog.Infof("dot1: %v", dot1)

	dot2 := (startVector.GetStartPoint().X * x2) + (startVector.GetStartPoint().Y * y2)
	glog.Infof("dot2: %v", dot2)

	tailSlope := -(1 / startVector.GetSlope())
	glog.Infof("tailSlope: %v", tailSlope)
	tailYIntercept := shape.GetYInterceptByPointAndSlope(startVector.GetStartPoint(), tailSlope)
	glog.Infof("tailYIntercept: %v", tailYIntercept)

	car = &Car{
		LeftHeadlight:  shape.NewPoint(50, 265),
		RightHeadlight: shape.NewPoint(50, 235),
		LeftTaillight:  shape.NewPoint(0, 265),
		RightTaillight: shape.NewPoint(0, 235),
		Status:         "stopped",
	}
	// go car.driver()
	return car
}
