package engine

import (
	"time"
)

// TakeOff will start the blades and raise the drone to a normal flying height.
func (e *Engine) TakeOff() {
	debug("Take off")
	check(e.drone.TakeOff())
	time.Sleep(5000 * time.Millisecond)
}

// Land will lower the drone to the ground and stop the blades.
func (e *Engine) Land() {
	debug("Land")
	check(e.drone.Land())
	time.Sleep(5000 * time.Millisecond)
}

// Forward will move the drone forward at the given speed.
func (e *Engine) Forward(speed int) {
	debug("Forward, speed: %v", speed)
	check(e.drone.Backward(0))
	check(e.drone.Forward(speed))
}

// Backward will move the drone backward at the given speed.
func (e *Engine) Backward(speed int) {
	debug("Backward, speed: %v", speed)
	check(e.drone.Forward(0))
	check(e.drone.Backward(speed))
}

// Left will move the drone left at the given speed.
func (e *Engine) Left(speed int) {
	debug("Left, speed: %v", speed)
	check(e.drone.Right(0))
	check(e.drone.Left(speed))
}

// Right will move the drone right at the given speed.
func (e *Engine) Right(speed int) {
	debug("Right, speed: %v", speed)
	check(e.drone.Left(0))
	check(e.drone.Right(speed))
}

// Up will move the drone up at the given speed.
func (e *Engine) Up(speed int) {
	debug("Up, speed: %v", speed)
	check(e.drone.Down(0))
	check(e.drone.Up(speed))
}

// Down will move the drone down at the given speed.
func (e *Engine) Down(speed int) {
	debug("Down, speed: %v", speed)
	check(e.drone.Up(0))
	check(e.drone.Down(speed))
}

// RotateLeft will rotate the drone left at the given speed.
func (e *Engine) RotateLeft(speed int) {
	debug("Rotate left, speed: %v", speed)
	check(e.drone.Clockwise(-speed))
}

// RotateRight will rotate the drone right at the given speed.
func (e *Engine) RotateRight(speed int) {
	debug("Rotate right, speed: %v", speed)
	check(e.drone.Clockwise(speed))
}

// FrontFlip will cause the drone to perform a front flip.
func (e *Engine) FrontFlip() {
	debug("Front flip")
	check(e.drone.FrontFlip())
	time.Sleep(3000 * time.Millisecond)
}

// BackFlip will cause the drone to perform a back flip.
func (e *Engine) BackFlip() {
	debug("Back flip")
	check(e.drone.BackFlip())
	time.Sleep(3000 * time.Millisecond)
}

// LeftFlip will cause the drone to perform a left flip.
func (e *Engine) LeftFlip() {
	debug("Left flip")
	check(e.drone.LeftFlip())
	time.Sleep(3000 * time.Millisecond)
}

// RightFlip will cause the drone to perform a right flip.
func (e *Engine) RightFlip() {
	debug("Right flip")
	check(e.drone.RightFlip())
	time.Sleep(3000 * time.Millisecond)
}

// Bounce will toggle "bouncing" the drone up and down.
func (e *Engine) Bounce() {
	debug("Toggle bounce")
	check(e.drone.Bounce())
}

// ResetMovement will set all drone movement to 0.
func (e *Engine) ResetMovement() {
	debug("Reset movement")
	check(e.drone.Forward(0))
	check(e.drone.Backward(0))
	check(e.drone.Up(0))
	check(e.drone.Down(0))
	check(e.drone.Left(0))
	check(e.drone.Right(0))
	check(e.drone.Clockwise(0))
}
