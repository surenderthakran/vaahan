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
	speed                               float64
	turningAngle                        geo.Angle
	FrontCenter                         *geo.Point `json:"front_center"`
	BackCenter                          *geo.Point `json:"back_center"`
	vector                              *geo.Ray
	internalAngle                       geo.Angle
	distanceOfCornersFromOppositeCenter float64
	commChannel                         chan string
	FL                                  *geo.Point `json:"front_left"`
	FR                                  *geo.Point `json:"front_right"`
	BL                                  *geo.Point `json:"back_left"`
	BR                                  *geo.Point `json:"back_right"`
	Status                              CarStatus  `json:"status"`
}

type CarStatus string

var (
	car      *Car
	carSTOP  CarStatus = "STOP"
	carDRIVE CarStatus = "DRIVE"
)

func (car *Car) Drive() {
	if car.Status == carSTOP {
		go car.drive()
	}
	car.Status = carDRIVE
}

func (car *Car) Stop() {
	car.Status = carSTOP
}

func GetCar() (*Car, error) {
	if car == nil {
		glog.Error("car not found")
		return nil, fmt.Errorf("car not found")
	}
	return car, nil
}

func (car *Car) drive() {
	for {
		if car.Status == carSTOP {
			glog.Info("stopping car")
			break
		} else if car.Status == carDRIVE {
			glog.Info("car is moving")
			car.turnRight()
			car.updateCorners()
		}
		time.Sleep(time.Second / 2)
	}
}

func (car *Car) moveForward() {
	glog.Info("inside car.moveForward()")
	point := car.vector.FindPointAtDistance(car.speed)
	car.vector = geo.NewRayByPointAndDirection(point, car.vector.Angle())
}

func (car *Car) turnRight() {
	glog.Info("inside car.turnRight()")
	car.vector.SetAngle(car.vector.Angle() - car.turningAngle)
	car.moveForward()
}

func (car *Car) turnLeft() {
	glog.Info("inside car.turnLeft()")
	car.vector.SetAngle(car.vector.Angle() + car.turningAngle)
	car.moveForward()
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

	thetaForFR := car.vector.Angle().Radians() - alpha
	xRelativeForFR := math.Cos(thetaForFR) * car.distanceOfCornersFromOppositeCenter
	yRelativeForFR := math.Sin(thetaForFR) * car.distanceOfCornersFromOppositeCenter
	frontRight := geo.NewPoint(car.vector.StartPoint().X+xRelativeForFR, car.vector.StartPoint().Y+yRelativeForFR)
	frontRight.RoundTo(2)
	car.FR = frontRight

	thetaForFL := car.vector.Angle().Radians() + alpha
	xRelativeForFL := math.Cos(thetaForFL) * car.distanceOfCornersFromOppositeCenter
	yRelativeForFL := math.Sin(thetaForFL) * car.distanceOfCornersFromOppositeCenter
	frontLeft := geo.NewPoint(car.vector.StartPoint().X+xRelativeForFL, car.vector.StartPoint().Y+yRelativeForFL)
	frontLeft.RoundTo(2)
	car.FL = frontLeft

	thetaForBack := math.Pi - (car.vector.Angle().Radians() + (math.Pi / 2))
	xRelativeForBackCorners := math.Cos(thetaForBack) * (car.Width / 2)
	yRelativeForBackCorners := math.Sin(thetaForBack) * (car.Width / 2)

	var backLeft *geo.Point
	backLeft = geo.NewPoint(car.vector.StartPoint().X-xRelativeForBackCorners, car.vector.StartPoint().Y+yRelativeForBackCorners)
	backLeft.RoundTo(2)
	car.BL = backLeft

	var backRight *geo.Point
	backRight = geo.NewPoint(car.vector.StartPoint().X+xRelativeForBackCorners, car.vector.StartPoint().Y-yRelativeForBackCorners)
	backRight.RoundTo(2)
	car.BR = backRight

	car.FrontCenter = car.vector.FindPointAtDistance(car.Length)
	car.BackCenter = car.vector.StartPoint()
}

func New(track *track.Track) *Car {
	if car == nil {
		car = &Car{
			Length:       40,
			Width:        18,
			speed:        10,
			turningAngle: math.Pi / 8,
			vector:       track.StartVector,
			Status:       carSTOP,
		}
	} else {
		car.vector = track.StartVector
		car.Status = carSTOP
	}
	car.updateCorners()

	return car
}
