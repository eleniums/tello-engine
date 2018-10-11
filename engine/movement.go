package engine

import (
	"log"
	"time"
)

// TakeOff will start the blades and raise the drone to a normal flying height.
func (e *Engine) TakeOff() {
	err := e.drone.TakeOff()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(5000 * time.Millisecond)
}

// Land will lower the drone to the ground and stop the blades.
func (e *Engine) Land() {
	err := e.drone.Land()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(5000 * time.Millisecond)
}

// ResetMovement will set all drone movement to 0.
func (e *Engine) ResetMovement() {
	log.Println("Reset drone position")
	e.drone.Forward(0)
	e.drone.Backward(0)
	e.drone.Up(0)
	e.drone.Down(0)
	e.drone.Left(0)
	e.drone.Right(0)
	e.drone.Clockwise(0)
}
