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
	Length                             float64
	Width                              float64
	FrontCenter                        *geo.Point `json:"front_center"`
	BackCenter                         *geo.Point `json:"back_center"`
	internalAngle                      geo.Angle
	distanceOfLightsFromOppositeCenter float64
	LHL                                *geo.Point `json:"left_headlight"`
	RHL                                *geo.Point `json:"right_headlight"`
	LTL                                *geo.Point `json:"left_taillight"`
	RTL                                *geo.Point `json:"right_taillight"`
	Status                             string     `json:"status"`
}

var (
	car *Car
)

func (car *Car) moveForward(units float64) {
	glog.Info("inside car.moveForward()")
	if units == 0 {
		units = 1
	}
	car.LHL.X = car.LHL.X + units
	// car.LHL.Y = car.LHL.Y + units

	car.RHL.X = car.RHL.X + units
	// car.RHL.Y = car.RHL.Y + units

	car.LTL.X = car.LTL.X + units
	// car.LTL.Y = car.LTL.Y + units

	car.RTL.X = car.RTL.X + units
	// car.RTL.Y = car.RTL.Y + units
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

func (car *Car) setHeadlights(carVector *geo.Ray) {
	var alpha float64
	if car.internalAngle == 0 {
		tan := car.Width / (2 * car.Length)
		alpha = math.Tanh(tan)
		car.internalAngle = geo.Angle(alpha)
		car.distanceOfLightsFromOppositeCenter = car.Length / math.Cos(alpha)
	} else {
		alpha = car.internalAngle.Radians()
	}

	thetaForRHL := carVector.GetAngle().Radians() - alpha
	xRelativeForRHL := math.Cos(thetaForRHL) * car.distanceOfLightsFromOppositeCenter
	yRelativeForRHL := math.Sin(thetaForRHL) * car.distanceOfLightsFromOppositeCenter
	rightHeadlight := geo.NewPoint(carVector.GetStartPoint().X+xRelativeForRHL, carVector.GetStartPoint().Y+yRelativeForRHL)
	rightHeadlight.RoundTo(2)
	car.RHL = rightHeadlight

	thetaForLHL := carVector.GetAngle().Radians() + alpha
	xRelativeForLHL := math.Cos(thetaForLHL) * car.distanceOfLightsFromOppositeCenter
	yRelativeForLHL := math.Sin(thetaForLHL) * car.distanceOfLightsFromOppositeCenter
	leftHeadlight := geo.NewPoint(carVector.GetStartPoint().X+xRelativeForLHL, carVector.GetStartPoint().Y+yRelativeForLHL)
	leftHeadlight.RoundTo(2)
	car.LHL = leftHeadlight

	thetaForTail := math.Pi - (carVector.GetAngle().Radians() + (math.Pi / 2))
	xRelativeForTail := math.Cos(thetaForTail) * (car.Width / 2)
	yRelativeForTail := math.Sin(thetaForTail) * (car.Width / 2)
	var leftTaillight *geo.Point
	if carVector.GetAngle().Radians() >= 0 && carVector.GetAngle().Radians() < math.Pi/2 {
		leftTaillight = geo.NewPoint(carVector.GetStartPoint().X-xRelativeForTail, carVector.GetStartPoint().Y+yRelativeForTail)
	}
	leftTaillight.RoundTo(2)
	car.LTL = leftTaillight

	var rightTaillight *geo.Point
	if carVector.GetAngle().Radians() >= 0 && carVector.GetAngle().Radians() < math.Pi/2 {
		rightTaillight = geo.NewPoint(carVector.GetStartPoint().X+xRelativeForTail, carVector.GetStartPoint().Y-yRelativeForTail)
	}
	rightTaillight.RoundTo(2)
	car.RTL = rightTaillight
}

func New(track *track.Track) *Car {
	startVector := track.StartVector

	length := float64(50)
	width := float64(30)
	car = &Car{
		Length:      length,
		Width:       width,
		FrontCenter: startVector.FindPointAtDistance(length),
		BackCenter:  startVector.GetStartPoint(),
		RTL:         geo.NewPoint(0, 235),
		Status:      "stopped",
	}

	car.setHeadlights(startVector)

	// go car.driver()
	return car
}
