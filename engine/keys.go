package engine

import (
	"log"

	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

// KeyHandler handles all key press events.
func KeyHandler(drone *tello.Driver, key keyboard.KeyEvent) {
	log.Printf("Key pressed: %s", key.Char)
	switch key.Key {
	case keyboard.A:
		drone.Right(0)
		drone.Left(Speed)
	case keyboard.D:
		drone.Left(0)
		drone.Right(Speed)
	case keyboard.W:
		drone.Backward(0)
		drone.Forward(Speed)
	case keyboard.S:
		drone.Forward(0)
		drone.Backward(Speed)
	case keyboard.ArrowUp:
		drone.Down(0)
		drone.Up(Speed)
	case keyboard.ArrowDown:
		drone.Up(0)
		drone.Down(Speed)
	case keyboard.ArrowLeft:
		drone.Clockwise(-RotateSpeed)
	case keyboard.ArrowRight:
		drone.Clockwise(RotateSpeed)
	case keyboard.Z:
		drone.Land()
	case keyboard.X:
		drone.TakeOff()
	case keyboard.Q:
		drone.LeftFlip()
	case keyboard.E:
		drone.RightFlip()
	case keyboard.R:
		drone.BackFlip()
	case keyboard.F:
		drone.FrontFlip()
	case keyboard.B:
		drone.Bounce()
	case keyboard.Escape, keyboard.Spacebar:
		fallthrough
	default:
		resetDroneMovement(drone)
	}
}
