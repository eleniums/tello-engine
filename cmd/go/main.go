package main

import (
	"flag"

	"github.com/eleniums/tello-engine/engine"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	// open window for receiving keyboard input
	win := run()

	for !win.Closed() {
		processKeys(e, win)
	}
}

func run() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "tello-go",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	return win
}

func processKeys(e *engine.Engine, win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyX) {
		e.TakeOff()
	} else if win.Pressed(pixelgl.KeyZ) {
		e.Land()
	}

	if win.JustPressed(pixelgl.KeyO) {
		speed -= 10
	} else if win.JustPressed(pixelgl.KeyP) {
		speed += 10
	}

	if win.Pressed(pixelgl.KeyLeft) {
		e.Left(speed)
	} else if win.Pressed(pixelgl.KeyRight) {
		e.Right(speed)
	}

	if win.Pressed(pixelgl.KeyUp) {
		e.Forward(speed)
	} else if win.Pressed(pixelgl.KeyDown) {
		e.Backward(speed)
	}

	if win.Pressed(pixelgl.KeyA) {
		e.RotateLeft(speed)
	} else if win.Pressed(pixelgl.KeyD) {
		e.RotateRight(speed)
	}

	if win.Pressed(pixelgl.KeyW) {
		e.Up(speed)
	} else if win.Pressed(pixelgl.KeyS) {
		e.Down(speed)
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
