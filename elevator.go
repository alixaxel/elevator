package elevator

import (
	"math"
)

type (
	Floor     int8
	Direction int8
)

type elevator struct {
	id           int8
	floor        Floor
	destinations map[Floor]int
}

type Elevator interface {
	ID() int8
	Step() bool
	Cost(floor Floor, direction Direction) float64
	Floor() Floor
	Request(floor Floor) int
	Direction() Direction
	Destination() Floor
	Destinations() map[Floor]float64
}

// Initializes an elevator to the ground (0) floor.
func NewElevator(id int8) Elevator {
	return &elevator{
		id:           id,
		floor:        0,
		destinations: map[Floor]int{},
	}
}

// Returns the supplied ID for the elevator.
func (self elevator) ID() int8 {
	return self.id
}

// Advances to the next goal floor and returns true.
// If there are no goal floors, false is returned instead.
func (self *elevator) Step() bool {
	direction := self.Direction()
	destination := self.Destination()

	if len(self.destinations) > 0 {
		self.floor = destination; delete(self.destinations, destination)
	}

	return direction != Direction(0)
}

// Compute how expensive it is to accept a pickup request.
func (self elevator) Cost(floor Floor, direction Direction) float64 {
	result := math.Inf(+1)
	moving := self.Direction()

	if moving == Direction(0) {
		result = math.Abs(float64(self.Floor()) - float64(floor))
	} else if moving == direction {
		if moving == Direction(+1) {
			result = math.Abs(float64(self.Floor()) - math.Min(float64(floor), float64(self.Destination())))
		} else if moving == Direction(-1) {
			result = math.Abs(float64(self.Floor()) - math.Max(float64(floor), float64(self.Destination())))
		}
	} else {
		result = math.Abs(float64(self.Floor()) - float64(self.Destination())) + math.Abs(float64(self.Destination()) - float64(floor))
	}

	return result
}

// Returns the floor where the elevator is at.
func (self elevator) Floor() Floor {
	return self.floor
}

// Add floor to the destination queue.
func (self elevator) Request(floor Floor) int {
	_, ok := self.destinations[floor]; if ok {
		self.destinations[floor]++
	} else {
		self.destinations[floor] = 1
	}

	return self.destinations[floor]
}

// Returns +1 if the elevator is moving up or -1 if it's moving down.
// If the elevator is not moving, 0 is returned instead.
func (self elevator) Direction() Direction {
	floor := self.Floor()
	destination := self.Destination()

	if destination > floor {
		return Direction(+1)
	} else if destination < floor {
		return Direction(-1)
	}

	return Direction(0)
}

// Returns the nearest destination floor.
func (self elevator) Destination() Floor {
	result := self.Floor()
	minimum := math.Inf(+1)

	for floor, distance := range self.Destinations() {
		if distance < minimum {
			minimum = distance; result = floor
		}
	}

	return result
}

// Returns a map of floors and respective distance from current position.
func (self elevator) Destinations() map[Floor]float64 {
	result := map[Floor]float64{}

	for target, _ := range self.destinations {
		result[target] = math.Abs(float64(self.Floor() - target))
	}

	return result
}
