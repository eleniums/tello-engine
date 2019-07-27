package main

import (
	"flag"
	"fmt"

	"github.com/eleniums/tello-engine/engine"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const (
	minSpeed = 0
	maxSpeed = 100
)

var (
	speed = 50
)

func main() {
	flag.BoolVar(&engine.Debug, "debug", false, "true to enable debug logging")
	flag.Parse()

	// prepare to handle input and start the drone
	e := engine.NewEngine()
	e.Start(false)
	defer e.Stop()

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
		Title:  "tello-go",
		Bounds: pixel.R(0, 0, 1024, 768),
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
	if win.Pressed(pixelgl.KeyX) {
		e.TakeOff()
	} else if win.Pressed(pixelgl.KeyZ) {
		e.Land()
	}

	if win.JustPressed(pixelgl.KeyO) {
		speed -= 10
		if speed < minSpeed {
			speed = minSpeed
		}
	} else if win.JustPressed(pixelgl.KeyP) {
		speed += 10
		if speed > maxSpeed {
			speed = maxSpeed
		}
	}

	if win.Pressed(pixelgl.KeyLeft) {
		e.Left(speed)
	} else if win.Pressed(pixelgl.KeyRight) {
		e.Right(speed)
	} else {
		e.Left(0)
	}

	if win.Pressed(pixelgl.KeyUp) {
		e.Forward(speed)
	} else if win.Pressed(pixelgl.KeyDown) {
		e.Backward(speed)
	} else {
		e.Forward(0)
	}

	if win.Pressed(pixelgl.KeyA) {
		e.RotateLeft(speed)
	} else if win.Pressed(pixelgl.KeyD) {
		e.RotateRight(speed)
	} else {
		e.RotateLeft(0)
	}

	if win.Pressed(pixelgl.KeyW) {
		e.Up(speed)
	} else if win.Pressed(pixelgl.KeyS) {
		e.Down(speed)
	} else {
		e.Up(0)
	}

	if win.Pressed(pixelgl.KeyQ) {
		e.LeftFlip()
	} else if win.Pressed(pixelgl.KeyE) {
		e.RightFlip()
	} else if win.Pressed(pixelgl.KeyF) {
		e.FrontFlip()
	} else if win.Pressed(pixelgl.KeyR) {
		e.BackFlip()
	}

	if win.Pressed(pixelgl.KeyB) {
		e.Bounce()
	}

	if win.Pressed(pixelgl.KeySpace) {
		e.ResetMovement()
	}
}
