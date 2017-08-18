package car

import (
	"fmt"

	"vaahan/shape"

	glog "github.com/golang/glog"
)

type Car struct {
	LeftHeadlight  *shape.Point `json:"left_headlight"`
	RightHeadlight *shape.Point `json:"right_headlight"`
	LeftTaillight  *shape.Point `json:"left_taillight"`
	RightTaillight *shape.Point `json:"right_taillight"`
}

var (
	car *Car
)

func (car Car) moveForward(units float64) {
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

func GetCar() (*Car, error) {
	if car == nil {
		glog.Error("car not found")
		return nil, fmt.Errorf("car not found")
	}
	return car, nil
}

func New(startVector *shape.Line) *Car {
	fmt.Println(startVector)
	tailSlope := -(1 / startVector.GetSlope())
	fmt.Println(tailSlope)
	tailYIntercept := shape.GetYInterceptByPointAndSlope(startVector.GetStartPoint(), tailSlope)
	fmt.Println(tailYIntercept)

	car = &Car{
		LeftHeadlight:  shape.NewPoint(50, 265),
		RightHeadlight: shape.NewPoint(50, 235),
		LeftTaillight:  shape.NewPoint(0, 265),
		RightTaillight: shape.NewPoint(0, 235),
	}
	return car
}
