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

## Install

```shell
go get github.com/alixaxel/elevator
```
