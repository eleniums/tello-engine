package engine

import (
	"time"
)

// TakeOff will start the blades and raise the drone to a normal flying height.
func (e *Engine) TakeOff() {
	check(e.drone.TakeOff())
	time.Sleep(5000 * time.Millisecond)
}

// Land will lower the drone to the ground and stop the blades.
func (e *Engine) Land() {
	check(e.drone.Land())
	time.Sleep(5000 * time.Millisecond)
}

// Forward will move the drone forward at the given speed.
func (e *Engine) Forward(speed int) {
	check(e.drone.Forward(speed))
}

// ResetMovement will set all drone movement to 0.
func (e *Engine) ResetMovement() {
	check(e.drone.Forward(0))
	check(e.drone.Backward(0))
	check(e.drone.Up(0))
	check(e.drone.Down(0))
	check(e.drone.Left(0))
	check(e.drone.Right(0))
	check(e.drone.Clockwise(0))
}
