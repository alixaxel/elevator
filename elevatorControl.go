package elevator

import (
	"math"
)

type elevatorControl struct {
	elevators []Elevator
}

type ElevatorControl interface {
	Pickup(origin Floor, direction Direction) Elevator
	Status() [][]int8
	Step() uint8
}

// Returns a new Elevator Control System instance.
func NewElevatorControl(elevators ...Elevator) ElevatorControl {
	return &elevatorControl{
		elevators: elevators,
	}
}

// Returns the nearest available elevator, given origin floor and target direction.
func (self elevatorControl) Pickup(origin Floor, direction Direction) Elevator {
	var result Elevator

	minimum := math.Inf(+1)

	for _, elevator := range self.elevators {
		if cost := elevator.Cost(origin, direction); cost < minimum {
			minimum = cost; result = elevator
		}
	}

	result.Request(origin)

	return result
}

// Return a bi-dimensional slice, where the each value triplet represents Elevator ID, Floor Number and Goal Floor Number.
func (self elevatorControl) Status() [][]int8 {
	result := [][]int8{}

	for _, elevator := range self.elevators {
		result = append(result, []int8{
			int8(elevator.ID()),
			int8(elevator.Floor()),
			int8(elevator.Destination()),
		})
	}

	return result
}

// Time-steps all the elevators and returns how many where moved in this step.
func (self elevatorControl) Step() uint8 {
	result := uint8(0)

	for _, elevator := range self.elevators {
		if elevator.Step() == true {
			result++
		}
	}

	return result
}
