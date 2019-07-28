package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/eleniums/tello-engine/engine"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	minSpeed           = 0
	maxSpeed           = 100
	defaultMoveSpeed   = 50
	defaultRotateSpeed = 80
)

var (
	moveSpeed   = defaultMoveSpeed
	rotateSpeed = defaultRotateSpeed

	moveXAxis  = false
	moveYAxis  = false
	moveZAxis  = false
	moveRotate = false
)

func main() {
	flag.BoolVar(&engine.Debug, "debug", false, "true to enable debug logging")
	flag.Parse()

	// prepare to handle input and start the drone
	e := engine.NewEngine()
	e.Start(false)
	defer e.Stop()

	e.StartVideoStream()

	pixelgl.Run(func() {
		// open window for receiving keyboard input
		win, err := run()
		if err != nil {
			panic(err)
		}

		// main loop to handle keyboard input
		for !win.Closed() {
			win.UpdateInput()
			processKeys(e, win)
			if win.Pressed(pixelgl.KeyEscape) {
				break
			}
		}
	})
}

func run() (*pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "tello-engine",
		Bounds: pixel.R(0, 0, 480, 360),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return nil, fmt.Errorf("Error creating window: %v", err)
	}

	win.SetSmooth(true)

	return win, nil
}

func processKeys(e *engine.Engine, win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyX) {
		e.TakeOff()
	} else if win.JustPressed(pixelgl.KeyZ) {
		e.Land()
	}

	if win.JustPressed(pixelgl.KeyO) {
		moveSpeed -= 10
		if moveSpeed < minSpeed {
			moveSpeed = minSpeed
		}
		if engine.Debug {
			log.Printf("Move speed decreased to: %v", moveSpeed)
		}
	} else if win.JustPressed(pixelgl.KeyP) {
		moveSpeed += 10
		if moveSpeed > maxSpeed {
			moveSpeed = maxSpeed
		}
		if engine.Debug {
			log.Printf("Move speed increased to: %v", moveSpeed)
		}
	}

	if win.JustPressed(pixelgl.KeyU) {
		rotateSpeed -= 10
		if rotateSpeed < minSpeed {
			rotateSpeed = minSpeed
		}
		if engine.Debug {
			log.Printf("Rotate speed decreased to: %v", rotateSpeed)
		}
	} else if win.JustPressed(pixelgl.KeyI) {
		rotateSpeed += 10
		if rotateSpeed > maxSpeed {
			rotateSpeed = maxSpeed
		}
		if engine.Debug {
			log.Printf("Rotate speed increased to: %v", rotateSpeed)
		}
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		moveXAxis = true
		e.Left(moveSpeed)
	} else if win.JustPressed(pixelgl.KeyRight) {
		moveXAxis = true
		e.Right(moveSpeed)
	} else if moveXAxis && !win.Pressed(pixelgl.KeyLeft) && !win.Pressed(pixelgl.KeyRight) {
		moveXAxis = false
		e.Left(0)
	}

	if win.JustPressed(pixelgl.KeyUp) {
		moveYAxis = true
		e.Forward(moveSpeed)
	} else if win.JustPressed(pixelgl.KeyDown) {
		moveYAxis = true
		e.Backward(moveSpeed)
	} else if moveYAxis && !win.Pressed(pixelgl.KeyUp) && !win.Pressed(pixelgl.KeyDown) {
		moveYAxis = false
		e.Forward(0)
	}

	if win.JustPressed(pixelgl.KeyA) {
		moveRotate = true
		e.RotateLeft(rotateSpeed)
	} else if win.JustPressed(pixelgl.KeyD) {
		moveRotate = true
		e.RotateRight(rotateSpeed)
	} else if moveRotate && !win.Pressed(pixelgl.KeyA) && !win.Pressed(pixelgl.KeyD) {
		moveRotate = false
		e.RotateLeft(0)
	}

	if win.JustPressed(pixelgl.KeyW) {
		moveZAxis = true
		e.Up(moveSpeed)
	} else if win.JustPressed(pixelgl.KeyS) {
		moveZAxis = true
		e.Down(moveSpeed)
	} else if moveZAxis && !win.Pressed(pixelgl.KeyW) && !win.Pressed(pixelgl.KeyS) {
		moveZAxis = false
		e.Up(0)
	}

	if win.JustPressed(pixelgl.KeyQ) {
		e.LeftFlip()
	} else if win.JustPressed(pixelgl.KeyE) {
		e.RightFlip()
	} else if win.JustPressed(pixelgl.KeyF) {
		e.FrontFlip()
	} else if win.JustPressed(pixelgl.KeyR) {
		e.BackFlip()
	}

	if win.JustPressed(pixelgl.KeyB) {
		e.Bounce()
	}

	if win.JustPressed(pixelgl.KeySpace) {
		moveSpeed = defaultMoveSpeed
		rotateSpeed = defaultRotateSpeed
		e.ResetMovement()
	}
}
