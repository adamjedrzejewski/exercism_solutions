package elon

import "fmt"

func (car *Car) Drive() {
	if car == nil {
		return
	}
	if car.batteryDrain > car.battery {
		return
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
}

func (car *Car) DisplayDistance() string {
	if car == nil {
		return "Car is nil"
	}
	return fmt.Sprintf("Driven %d meters", car.distance)
}

func (car *Car) DisplayBattery() string {
	if car == nil {
		return "Car is nil"
	}
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

func (car *Car) CanFinish(trackDistance int) bool {
	if car == nil {
		return false
	}
	if car.batteryDrain == 0 {
		return true
	}
	distance := car.battery / car.batteryDrain
	return distance*car.speed >= trackDistance
}
