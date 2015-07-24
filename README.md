# elevator

Elevator Control System in Go

## Usage

```go
package main

import (
    "fmt"

    "github.com/alixaxel/elevator"
)

func main() {
    elevators := []elevator.Elevator{
        elevator.NewElevator(1),
        elevator.NewElevator(2),
        elevator.NewElevator(3),
    }

    control := elevator.NewElevatorControl(elevators[0], elevators[1], elevators[2])

    fmt.Println(control.Step(), control.Status())
    fmt.Println(control.Step(), control.Status())

    control.Pickup(5, -1)

    fmt.Println(control.Step(), control.Status())
    fmt.Println(control.Step(), control.Status())
    fmt.Println(control.Step(), control.Status())

    control.Pickup(92, -1)
    control.Pickup(2, +1)

    fmt.Println(control.Step(), control.Status())
    fmt.Println(control.Step(), control.Status())

    control.Pickup(1, +1)
    control.Pickup(0, +1)
    control.Pickup(64, +1)

    fmt.Println(control.Step(), control.Status())
    fmt.Println(control.Step(), control.Status())
}
```

## Output

```
0 [[1 0 0] [2 0 0] [3 0 0]]
0 [[1 0 0] [2 0 0] [3 0 0]]
1 [[1 5 5] [2 0 0] [3 0 0]]
0 [[1 5 5] [2 0 0] [3 0 0]]
0 [[1 5 5] [2 0 0] [3 0 0]]
2 [[1 92 92] [2 2 2] [3 0 0]]
0 [[1 92 92] [2 2 2] [3 0 0]]
2 [[1 64 64] [2 1 1] [3 0 0]]
0 [[1 64 64] [2 1 1] [3 0 0]]
```

## Scheduling Algorithm

In order to choose which elevator should be requested to serve a pickup request, the following cost heuristic is applied:

 - If the elevator is not moving, ```cost = abs(currentFloor - callingFloor)```
 - If the elevator is moving in the same direction of the pickup request:
  - ```cost = abs(currentFloor, min(callingFloor, nextFloor))``` if going up
  - ```cost = abs(currentFloor, max(callingFloor, nextFloor))``` if going down
 - If the elevator is not moving in the same direction of the pickup request, then:
  - ```cost = abs(currentFloor - nextFloor) + abs(nextFloor - callingFloor)```

The elevator with the lowest cost is then chosen to answer the pickup request.

## Install

```shell
go get github.com/alixaxel/elevator
```
