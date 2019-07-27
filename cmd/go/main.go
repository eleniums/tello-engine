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
	if win.JustPressed(pixelgl.KeyX) {
		e.TakeOff()
	} else if win.JustPressed(pixelgl.KeyZ) {
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

	if win.JustPressed(pixelgl.KeyLeft) {
		e.Left(speed)
	} else if win.JustPressed(pixelgl.KeyRight) {
		e.Right(speed)
	} else if win.Pressed(pixelgl.KeyLeft) || win.Pressed(pixelgl.KeyRight) {
		// do nothing
	} else {
		e.Left(0)
	}

	if win.JustPressed(pixelgl.KeyUp) {
		e.Forward(speed)
	} else if win.JustPressed(pixelgl.KeyDown) {
		e.Backward(speed)
	} else if win.Pressed(pixelgl.KeyUp) || win.Pressed(pixelgl.KeyDown) {
		// do nothing
	} else {
		e.Forward(0)
	}

	if win.JustPressed(pixelgl.KeyA) {
		e.RotateLeft(speed)
	} else if win.JustPressed(pixelgl.KeyD) {
		e.RotateRight(speed)
	} else if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyD) {
		// do nothing
	} else {
		e.RotateLeft(0)
	}

	if win.JustPressed(pixelgl.KeyW) {
		e.Up(speed)
	} else if win.JustPressed(pixelgl.KeyS) {
		e.Down(speed)
	} else if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyS) {
		// do nothing
	} else {
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
		e.ResetMovement()
	}
}
