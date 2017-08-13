package car

import (
	"fmt"

	"vaahan/shape"
)

type Car struct {
	Body *shape.Rectangle `json:"body"`
}

var (
	car *Car
)

func GetCar() *Car {
	if car == nil {
		car = New()
	}
	return car
}

func New() *Car {
	startVector := shape.NewRayByPointAndEquation(&shape.Point{0, 250}, 0, 250)
	fmt.Println(startVector)
	tailSlope := -(1 / startVector.GetSlope())
	fmt.Println(tailSlope)
	tailYIntercept := shape.GetYInterceptByPointAndSlope(startVector.GetStartPoint(), tailSlope)
	fmt.Println(tailYIntercept)
	return &Car{}
}
