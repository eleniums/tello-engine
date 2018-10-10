package engine

import (
	"log"
)

// ResetDroneMovement will set all drone movement to 0.
func (e *Engine) ResetDroneMovement() {
	log.Println("Reset drone position")
	e.drone.Forward(0)
	e.drone.Backward(0)
	e.drone.Up(0)
	e.drone.Down(0)
	e.drone.Left(0)
	e.drone.Right(0)
	e.drone.Clockwise(0)
}
