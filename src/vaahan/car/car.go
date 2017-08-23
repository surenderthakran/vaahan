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
	Length                              float64
	Width                               float64
	acceleration                        float64
	speed                               float64
	topSpeed                            float64
	turningRadius                       float64
	FrontCenter                         *geo.Point `json:"front_center"`
	BackCenter                          *geo.Point `json:"back_center"`
	vector                              *geo.Ray
	internalAngle                       geo.Angle
	distanceOfCornersFromOppositeCenter float64
	FL                                  *geo.Point `json:"front_left"`
	FR                                  *geo.Point `json:"front_right"`
	BL                                  *geo.Point `json:"back_left"`
	BR                                  *geo.Point `json:"back_right"`
	Status                              string     `json:"status"`
}

var (
	car *Car
)

func (car *Car) moveForward(units float64) {
	glog.Info("inside car.moveForward()")
	if units == 0 {
		units = 1
	}
	car.FL.X = car.FL.X + units
	// car.FL.Y = car.FL.Y + units

	car.FR.X = car.FR.X + units
	// car.FR.Y = car.FR.Y + units

	car.BL.X = car.BL.X + units
	// car.BL.Y = car.BL.Y + units

	car.BR.X = car.BR.X + units
	// car.BR.Y = car.BR.Y + units
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

func (car *Car) updateCorners() {
	var alpha float64
	if car.internalAngle == 0 {
		tan := car.Width / (2 * car.Length)
		alpha = math.Tanh(tan)
		car.internalAngle = geo.Angle(alpha)
		car.distanceOfCornersFromOppositeCenter = car.Length / math.Cos(alpha)
	} else {
		alpha = car.internalAngle.Radians()
	}

	thetaForFR := car.vector.GetAngle().Radians() - alpha
	xRelativeForFR := math.Cos(thetaForFR) * car.distanceOfCornersFromOppositeCenter
	yRelativeForFR := math.Sin(thetaForFR) * car.distanceOfCornersFromOppositeCenter
	frontRight := geo.NewPoint(car.vector.GetStartPoint().X+xRelativeForFR, car.vector.GetStartPoint().Y+yRelativeForFR)
	frontRight.RoundTo(2)
	car.FR = frontRight

	thetaForFL := car.vector.GetAngle().Radians() + alpha
	xRelativeForFL := math.Cos(thetaForFL) * car.distanceOfCornersFromOppositeCenter
	yRelativeForFL := math.Sin(thetaForFL) * car.distanceOfCornersFromOppositeCenter
	frontLeft := geo.NewPoint(car.vector.GetStartPoint().X+xRelativeForFL, car.vector.GetStartPoint().Y+yRelativeForFL)
	frontLeft.RoundTo(2)
	car.FL = frontLeft

	thetaForBack := math.Pi - (car.vector.GetAngle().Radians() + (math.Pi / 2))
	xRelativeForBackCorners := math.Cos(thetaForBack) * (car.Width / 2)
	yRelativeForBackCorners := math.Sin(thetaForBack) * (car.Width / 2)

	var backLeft *geo.Point
	backLeft = geo.NewPoint(car.vector.GetStartPoint().X-xRelativeForBackCorners, car.vector.GetStartPoint().Y+yRelativeForBackCorners)
	backLeft.RoundTo(2)
	car.BL = backLeft

	var backRight *geo.Point
	backRight = geo.NewPoint(car.vector.GetStartPoint().X+xRelativeForBackCorners, car.vector.GetStartPoint().Y-yRelativeForBackCorners)
	backRight.RoundTo(2)
	car.BR = backRight

	car.FrontCenter = car.vector.FindPointAtDistance(car.Length)
	car.BackCenter = car.vector.GetStartPoint()
}

func New(track *track.Track) *Car {
	car = &Car{
		Length:        38.5,
		Width:         16.95,
		acceleration:  67567.57,
		speed:         0,
		topSpeed:      430.55,
		turningRadius: 48,
		vector:        track.StartVector,
		Status:        "stopped",
	}

	car.updateCorners()

	// go car.driver()
	return car
}
