package car

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	geo "vaahan/gogeo/2d"
	"vaahan/gomind"
	"vaahan/track"

	glog "github.com/golang/glog"
)

type Car struct {
	vector             *geo.Ray
	track              *track.Track
	obstacles          []*geo.LineSegment
	specs              *CarSpecs
	Sensors            []*sensor `json:"sensors"`
	Points             Points    `json:"points"`
	Status             CarStatus `json:"status"`
	RestartOnCollision bool      `json:"restartOnCollision"`
}

type CarStatus string

type Points struct {
	FL     *geo.Point `json:"front_left"`
	FR     *geo.Point `json:"front_right"`
	BL     *geo.Point `json:"back_left"`
	BR     *geo.Point `json:"back_right"`
	FC     *geo.Point `json:"front_center"`
	BC     *geo.Point `json:"back_center"`
	center *geo.Point `json:"center"`
}

type CarSpecs struct {
	length                              float64
	width                               float64
	internalAngle                       geo.Angle
	distanceOfCornersFromOppositeCenter float64
	speed                               float64 // distance travelled in oneTimeUnit.
	turningAngle                        geo.Angle
}

type sensor struct {
	name         string
	Ray          *geo.Ray   `json:"ray"`
	Intersection *geo.Point `json:"intersection"`
	distance     float64
}

var (
	car          *Car
	carSTOP      CarStatus = "STOP"
	carDRIVE     CarStatus = "DRIVE"
	carCOLLISION CarStatus = "COLLISION"
	carSUCCESS   CarStatus = "SUCCESS"
)

const (
	oneTimeUnit = time.Second / 4
)

func (car *Car) drive() {
	for {
		glog.Info("===============================================================")
		if car.Status == carSTOP {
			glog.Info("stopping car")
			break
		} else if car.Status == carCOLLISION {
			if car.RestartOnCollision {
				glog.Info("restarting car")
				time.Sleep(4 * oneTimeUnit)
				InitCar()
				car.Status = carDRIVE
			} else {
				car.Status = carSTOP
			}
		} else if car.Status == carDRIVE {
			if car.collision() {
				car.Status = carCOLLISION
			} else {
				glog.Info("moving car")

				// move car.
				car.moveForward()

				// update car coordinates.
				car.updatePoints()

				// update sensors readings.
				if err := car.updateSensors(); err != nil {
					glog.Infof("unable to read from sensors: %v", err)
					car.Status = carSTOP
				}
			}
		}
		time.Sleep(oneTimeUnit)
	}
}

func (car *Car) moveForward() error {
	point := car.vector.FindPointAtDistance(car.specs.speed)
	newCarVector, err := geo.NewRayByPointAndDirection(point, car.vector.Angle())
	if err != nil {
		return fmt.Errorf("invalid car vector: %v", err)
	}
	car.vector = newCarVector
	return nil
}

func (car *Car) turnRight() {
	car.vector.SetAngle(car.vector.Angle() - car.specs.turningAngle)
	car.moveForward()
}

func (car *Car) turnLeft() {
	car.vector.SetAngle(car.vector.Angle() + car.specs.turningAngle)
	car.moveForward()
}

func (car *Car) collision() bool {
	glog.Info("inside car.collision()")
	// check if the car is currently colliding.
	if !car.insideTrack() {
		return true
	}

	for _, sensor := range car.Sensors {
		if sensor.distance <= car.specs.speed {
			glog.Infof("sensor collided: %v, %v, %f", sensor.name, sensor.Ray.Start, sensor.Ray.Angle(), sensor.distance)
			glog.Info("COLLISION!!")
			return true
		}
	}
	return false
}

func (car *Car) insideTrack() bool {
	corners := []*geo.Point{
		car.Points.FL,
		car.Points.FR,
		car.Points.BL,
		car.Points.BR,
	}
	for _, corner := range corners {
		if !car.track.PointInTrack(corner) {
			return false
		}
	}
	return true
}

func (car *Car) updatePoints() {
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
	frontRight := geo.NewPoint(car.vector.Start.X+xRelativeForFR, car.vector.Start.Y+yRelativeForFR)
	frontRight.RoundTo(2)
	car.Points.FR = frontRight

	thetaForFL := car.vector.Angle().Radians() + alpha
	xRelativeForFL := math.Cos(thetaForFL) * car.specs.distanceOfCornersFromOppositeCenter
	yRelativeForFL := math.Sin(thetaForFL) * car.specs.distanceOfCornersFromOppositeCenter
	frontLeft := geo.NewPoint(car.vector.Start.X+xRelativeForFL, car.vector.Start.Y+yRelativeForFL)
	frontLeft.RoundTo(2)
	car.Points.FL = frontLeft

	thetaForBack := math.Pi - (car.vector.Angle().Radians() + (math.Pi / 2))
	xRelativeForBackCorners := math.Cos(thetaForBack) * (car.specs.width / 2)
	yRelativeForBackCorners := math.Sin(thetaForBack) * (car.specs.width / 2)

	var backLeft *geo.Point
	backLeft = geo.NewPoint(car.vector.Start.X-xRelativeForBackCorners, car.vector.Start.Y+yRelativeForBackCorners)
	backLeft.RoundTo(2)
	car.Points.BL = backLeft

	var backRight *geo.Point
	backRight = geo.NewPoint(car.vector.Start.X+xRelativeForBackCorners, car.vector.Start.Y-yRelativeForBackCorners)
	backRight.RoundTo(2)
	car.Points.BR = backRight

	car.Points.FC = car.vector.FindPointAtDistance(car.specs.length)
	car.Points.BC = car.vector.Start
}

func (car *Car) updateSensors() error {
	// update all sensor's location and orientation according to car.
	frontCenterRay, err := geo.NewRayByPointAndDirection(car.Points.FC, car.vector.Angle())
	if err != nil {
		return fmt.Errorf("unable to load sensors: %v", err)
	}
	frontLeftRay, err := geo.NewRayByPointAndDirection(car.Points.FL, car.vector.Angle()+(math.Pi/4))
	if err != nil {
		return fmt.Errorf("unable to load sensors: %v", err)
	}
	frontRightRay, err := geo.NewRayByPointAndDirection(car.Points.FR, car.vector.Angle()-(math.Pi/4))
	if err != nil {
		return fmt.Errorf("unable to load sensors: %v", err)
	}
	car.Sensors = []*sensor{
		&sensor{
			name: "Front Center Sensor",
			Ray:  frontCenterRay,
		},
		&sensor{
			name: "Front Left Sensor",
			Ray:  frontLeftRay,
		},
		&sensor{
			name: "Front Right Sensor",
			Ray:  frontRightRay,
		},
	}

	// get sensor readings
	for _, sensor := range car.Sensors {
		// reset sensor readings.
		sensor.Intersection = nil
		sensor.distance = math.Inf(0)

		// iterate over all obstacles on track.
		for _, obstacle := range car.obstacles {
			// find sensor's ray's intersection point with obstacle.
			if intersection := sensor.Ray.Intersection(obstacle); intersection != nil {
				distance := sensor.Ray.Start.DistanceFrom(intersection)
				// update sensor's intersection data if an obstacle is closer than the previous obstacle.
				if distance < sensor.distance {
					sensor.Intersection = intersection
					sensor.distance = distance
				}
			}
		}
		glog.Infof("sensor: %v\tdistance: %v", sensor.name, sensor.distance)
	}

	return nil
}

func (car *Car) readObstacles() {
	car.obstacles = car.track.Boundary.Sides
}

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
		InitCar()
	}
	return car, nil
}

func InitCar() (*Car, error) {
	glog.Info("inside car.InitCar()")
	track, err := track.GetTrack()
	if err != nil {
		return nil, fmt.Errorf("unable to get track: %v", err)
	}
	if car == nil {
		glog.Info("creating new car at starting vector")
		car = &Car{
			track:  track,
			vector: track.StartVector,
			specs: &CarSpecs{
				length:       40,
				width:        18,
				speed:        5,
				turningAngle: math.Pi / 16,
			},
			Status: carSTOP,
		}
	} else {
		glog.Info("moving car to starting vector")
		car.vector = track.StartVector
		car.Status = carSTOP
	}

	car.updatePoints()
	car.readObstacles()

	if err := car.updateSensors(); err != nil {
		return nil, fmt.Errorf("unable to start car: %v", err)
	}

	trainingSet := [][][]float64{
		[][]float64{[]float64{0.1, 0.2}, []float64{0.3}},
		[][]float64{[]float64{0.15, 0.25}, []float64{0.4}},
		[][]float64{[]float64{0.12, 0.22}, []float64{0.34}},
		[][]float64{[]float64{0.01, 0.02}, []float64{0.03}},
		[][]float64{[]float64{0.2, 0.3}, []float64{0.5}},
		[][]float64{[]float64{0.3, 0.4}, []float64{0.7}},
		[][]float64{[]float64{0.4, 0.5}, []float64{0.9}},
		[][]float64{[]float64{0.5, 0.1}, []float64{0.6}},
		[][]float64{[]float64{0.6, 0.2}, []float64{0.8}},
		[][]float64{[]float64{0.7, 0.2}, []float64{0.9}},
	}

	testSet := [][][]float64{
		[][]float64{[]float64{0.1, 0.3}},
		[][]float64{[]float64{0.2, 0.4}},
		[][]float64{[]float64{0.3, 0.5}},
		[][]float64{[]float64{0.4, 0.1}},
		[][]float64{[]float64{0.5, 0.2}},
		[][]float64{[]float64{0.05, 0.2}},
	}

	// trainingSet := [][][]float64{
	// 	[][]float64{[]float64{0.1, 0.2, 0.3}, []float64{0.32}},
	// 	[][]float64{[]float64{0.2, 0.3, 0.4}, []float64{0.46}},
	// 	[][]float64{[]float64{0.3, 0.4, 0.5}, []float64{0.62}},
	// 	[][]float64{[]float64{0.4, 0.5, 0.6}, []float64{0.8}},
	// 	[][]float64{[]float64{0.5, 0.1, 0.2}, []float64{0.25}},
	// 	[][]float64{[]float64{0.6, 0.2, 0.3}, []float64{0.42}},
	// 	[][]float64{[]float64{0.7, 0.2, 0.3}, []float64{0.44}},
	// 	[][]float64{[]float64{0.8, 0.3, 0.4}, []float64{0.64}},
	// 	[][]float64{[]float64{0.9, 0.4, 0.5}, []float64{0.86}},
	// 	[][]float64{[]float64{0.8, 0.5, 0.1}, []float64{0.5}},
	// 	[][]float64{[]float64{0.7, 0.6, 0.2}, []float64{0.62}},
	// 	[][]float64{[]float64{0.6, 0.2, 0.5}, []float64{0.62}},
	// }

	// trainingSet := [][][]float64{
	// 	[][]float64{[]float64{0.05, 0.10}, []float64{0.01, 0.99}},
	// }

	// trainingSet := [][][]float64{
	// 	[][]float64{[]float64{0.01}, []float64{0.02}},
	// 	[][]float64{[]float64{0.02}, []float64{0.04}},
	// 	[][]float64{[]float64{0.04}, []float64{0.08}},
	// 	[][]float64{[]float64{0.05}, []float64{0.10}},
	// 	[][]float64{[]float64{0.07}, []float64{0.14}},
	// 	[][]float64{[]float64{0.08}, []float64{0.16}},
	// 	[][]float64{[]float64{0.10}, []float64{0.20}},
	// 	[][]float64{[]float64{0.11}, []float64{0.22}},
	// 	[][]float64{[]float64{0.13}, []float64{0.26}},
	// 	[][]float64{[]float64{0.14}, []float64{0.28}},
	// 	[][]float64{[]float64{0.16}, []float64{0.32}},
	// 	[][]float64{[]float64{0.17}, []float64{0.34}},
	// }

	mind, err := gomind.NewNeuralNetwork(len(trainingSet[0][0]), 3, len(trainingSet[0][1]))
	if err != nil {
		glog.Info(err)
	}

	mind.Describe()
	fmt.Println("========================================================")
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		index := rand.Intn(len(trainingSet))

		fmt.Println(fmt.Sprintf("%v %v", trainingSet[index][0], trainingSet[index][1]))
		mind.Train(trainingSet[index][0], trainingSet[index][1])

		fmt.Println(mind.LastOutput())
		fmt.Println(fmt.Sprintf("Error: %v\n", mind.CalculateError(trainingSet[index][1])))
	}
	fmt.Println(fmt.Sprintf("\nTotal Error: %v", mind.CalculateTotalError(trainingSet)))
	fmt.Println("========================================================")
	mind.Describe()
	fmt.Println("========================================================")
	for _, test := range testSet {
		fmt.Println(test)
		fmt.Println(mind.CalculateOutput(test[0]))
	}
	fmt.Println("========================================================")

	return car, nil
}
