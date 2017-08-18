package driver

import (
	"fmt"

	"vaahan/car"
	"vaahan/track"

	glog "github.com/golang/glog"
)

type Driver struct {
	car    *car.Car
	track  *track.Track
	status string
}

var (
	driver *Driver
)

func (driver Driver) Drive() {
	driver.status = "drive"
	car.moveForward(1)
}

func (driver Driver) Pause() {

}

func (driver Driver) Resume() {

}

func (driver Driver) Stop() {

}

func GetDriver() (*Driver, error) {
	if driver == nil {
		glog.Error("driver not found")
		return nil, fmt.Errorf("driver not found")
	}
	return driver, nil
}

func New(car *car.Car, track *track.Track) *Driver {
	driver = &Driver{
		car:   car,
		track: track,
	}
	return driver
}
