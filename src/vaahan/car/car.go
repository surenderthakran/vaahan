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
	vector    *geo.Ray
	track     *track.Track
	obstacles []*geo.LineSegment
	specs     *CarSpecs
	sensors   []*sensor
	Points    Points    `json:"points"`
	Status    CarStatus `json:"status"`
}

type CarStatus string

type Points struct {
	FL          *geo.Point `json:"front_left"`
	FR          *geo.Point `json:"front_right"`
	BL          *geo.Point `json:"back_left"`
	BR          *geo.Point `json:"back_right"`
	FrontCenter *geo.Point `json:"front_center"`
	BackCenter  *geo.Point `json:"back_center"`
}

type CarSpecs struct {
	length                              float64
	width                               float64
	internalAngle                       geo.Angle
	distanceOfCornersFromOppositeCenter float64
	speed                               float64
	turningAngle                        geo.Angle
}

type sensor struct {
	ray          *geo.Ray
	intersection *geo.Point
	distance     float64
}

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
			if car.collision() {
				car.Status = carSTOP
			} else {
				car.turnRight()
				car.updateCorners()
			}
		}
		time.Sleep(time.Second / 2)
	}
}

func (car *Car) moveForward() {
	glog.Info("inside car.moveForward()")
	point := car.vector.FindPointAtDistance(car.specs.speed)
	car.vector = geo.NewRayByPointAndDirection(point, car.vector.Angle())
}

func (car *Car) turnRight() {
	glog.Info("inside car.turnRight()")
	car.vector.SetAngle(car.vector.Angle() - car.specs.turningAngle)
	car.moveForward()
}

func (car *Car) turnLeft() {
	glog.Info("inside car.turnLeft()")
	car.vector.SetAngle(car.vector.Angle() + car.specs.turningAngle)
	car.moveForward()
}

func (car *Car) collision() bool {
	glog.Info("inside car.collision()")
	for _, sensor := range car.sensors {
		glog.Infof("sensor: %v, %v", sensor.ray.StartPoint(), sensor.ray.Angle())
		for _, obstacle := range car.obstacles {
			glog.Infof("obstacle: %v, %v", obstacle.StartPoint(), obstacle.EndPoint())
		}
	}
	return true
}

func (car *Car) updateCorners() {
	var alpha float64
	if car.specs.internalAngle == 0 {
		tan := car.specs.width / (2 * car.specs.length)
		alpha = math.Tanh(tan)
		car.specs.internalAngle = geo.Angle(alpha)
		car.specs.distanceOfCornersFromOppositeCenter = car.specs.length / math.Cos(alpha)
	} else {
		alpha = car.specs.internalAngle.Radians()
	}

	thetaForFR := car.vector.Angle().Radians() - alpha
	xRelativeForFR := math.Cos(thetaForFR) * car.specs.distanceOfCornersFromOppositeCenter
	yRelativeForFR := math.Sin(thetaForFR) * car.specs.distanceOfCornersFromOppositeCenter
	frontRight := geo.NewPoint(car.vector.StartPoint().X+xRelativeForFR, car.vector.StartPoint().Y+yRelativeForFR)
	frontRight.RoundTo(2)
	car.Points.FR = frontRight

	thetaForFL := car.vector.Angle().Radians() + alpha
	xRelativeForFL := math.Cos(thetaForFL) * car.specs.distanceOfCornersFromOppositeCenter
	yRelativeForFL := math.Sin(thetaForFL) * car.specs.distanceOfCornersFromOppositeCenter
	frontLeft := geo.NewPoint(car.vector.StartPoint().X+xRelativeForFL, car.vector.StartPoint().Y+yRelativeForFL)
	frontLeft.RoundTo(2)
	car.Points.FL = frontLeft

	thetaForBack := math.Pi - (car.vector.Angle().Radians() + (math.Pi / 2))
	xRelativeForBackCorners := math.Cos(thetaForBack) * (car.specs.width / 2)
	yRelativeForBackCorners := math.Sin(thetaForBack) * (car.specs.width / 2)

	var backLeft *geo.Point
	backLeft = geo.NewPoint(car.vector.StartPoint().X-xRelativeForBackCorners, car.vector.StartPoint().Y+yRelativeForBackCorners)
	backLeft.RoundTo(2)
	car.Points.BL = backLeft

	var backRight *geo.Point
	backRight = geo.NewPoint(car.vector.StartPoint().X+xRelativeForBackCorners, car.vector.StartPoint().Y-yRelativeForBackCorners)
	backRight.RoundTo(2)
	car.Points.BR = backRight

	car.Points.FrontCenter = car.vector.FindPointAtDistance(car.specs.length)
	car.Points.BackCenter = car.vector.StartPoint()
}

func (car *Car) initSensors() {
	car.sensors = []*sensor{
		&sensor{
			ray: geo.NewRayByPointAndDirection(car.Points.FrontCenter, car.vector.Angle()),
		},
	}
}

func (car *Car) readObstacles() {
	car.obstacles = car.track.Boundary.Sides()
}

func New(track *track.Track) *Car {
	if car == nil {
		car = &Car{
			track:  track,
			vector: track.StartVector,
			specs: &CarSpecs{
				length:       40,
				width:        18,
				speed:        10,
				turningAngle: math.Pi / 8,
			},
			Status: carSTOP,
		}
	} else {
		car.vector = track.StartVector
		car.Status = carSTOP
	}
	car.updateCorners()
	car.initSensors()
	car.readObstacles()

	return car
}
