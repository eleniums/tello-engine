package engine

import (
	"log"

	"gobot.io/x/gobot/platforms/keyboard"
)

// handleKeyPress handles a key press event.
func (e *Engine) handleKeyPress(key keyboard.KeyEvent) {
	log.Printf("Key pressed: %s", key.Char)
	switch key.Key {
	case keyboard.A:
		e.drone.Right(0)
		e.drone.Left(e.Speed)
	case keyboard.D:
		e.drone.Left(0)
		e.drone.Right(e.Speed)
	case keyboard.W:
		e.drone.Backward(0)
		e.drone.Forward(e.Speed)
	case keyboard.S:
		e.drone.Forward(0)
		e.drone.Backward(e.Speed)
	case keyboard.ArrowUp:
		e.drone.Down(0)
		e.drone.Up(e.Speed)
	case keyboard.ArrowDown:
		e.drone.Up(0)
		e.drone.Down(e.Speed)
	case keyboard.ArrowLeft:
		e.drone.Clockwise(-e.Speed)
	case keyboard.ArrowRight:
		e.drone.Clockwise(e.Speed)
	case keyboard.Z:
		e.drone.Land()
	case keyboard.X:
		e.drone.TakeOff()
	case keyboard.Q:
		e.drone.LeftFlip()
	case keyboard.E:
		e.drone.RightFlip()
	case keyboard.R:
		e.drone.BackFlip()
	case keyboard.F:
		e.drone.FrontFlip()
	case keyboard.B:
		e.drone.Bounce()
	case keyboard.Escape, keyboard.Spacebar:
		fallthrough
	default:
		e.ResetMovement()
	}
}
