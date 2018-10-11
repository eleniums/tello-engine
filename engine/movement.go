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
	check(e.drone.Backward(0))
	check(e.drone.Forward(speed))
}

// Backward will move the drone backward at the given speed.
func (e *Engine) Backward(speed int) {
	check(e.drone.Forward(0))
	check(e.drone.Backward(speed))
}

// Left will move the drone left at the given speed.
func (e *Engine) Left(speed int) {
	check(e.drone.Right(0))
	check(e.drone.Left(speed))
}

// Right will move the drone right at the given speed.
func (e *Engine) Right(speed int) {
	check(e.drone.Left(0))
	check(e.drone.Right(speed))
}

// Up will move the drone up at the given speed.
func (e *Engine) Up(speed int) {
	check(e.drone.Down(0))
	check(e.drone.Up(speed))
}

// Down will move the drone down at the given speed.
func (e *Engine) Down(speed int) {
	check(e.drone.Up(0))
	check(e.drone.Down(speed))
}

// RotateLeft will rotate the drone left at the given speed.
func (e *Engine) RotateLeft(speed int) {
	check(e.drone.Clockwise(-speed))
}

// RotateRight will rotate the drone right at the given speed.
func (e *Engine) RotateRight(speed int) {
	check(e.drone.Clockwise(speed))
}

// FrontFlip will cause the drone to perform a front flip.
func (e *Engine) FrontFlip() {
	check(e.drone.FrontFlip())
	time.Sleep(3000 * time.Millisecond)
}

// BackFlip will cause the drone to perform a back flip.
func (e *Engine) BackFlip() {
	check(e.drone.BackFlip())
	time.Sleep(3000 * time.Millisecond)
}

// LeftFlip will cause the drone to perform a left flip.
func (e *Engine) LeftFlip() {
	check(e.drone.LeftFlip())
	time.Sleep(3000 * time.Millisecond)
}

// RightFlip will cause the drone to perform a right flip.
func (e *Engine) RightFlip() {
	check(e.drone.RightFlip())
	time.Sleep(3000 * time.Millisecond)
}

// Bounce will toggle "bouncing" the drone up and down.
func (e *Engine) Bounce() {
	check(e.drone.Bounce())
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
