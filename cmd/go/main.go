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
}

func processKeys(e *engine.Engine, *pixelgl.Window) {
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
		e.Left()
	} else if win.Pressed(pixelgl.KeyRight) {
		e.Right()
	}

	if win.Pressed(pixelgl.KeyUp) {
		e.Forward()
	} else if win.Pressed(pixelgl.KeyDown) {
		e.Backward()
	} 

	if win.Pressed(pixelgl.KeyA) {
		e.RotateLeft()
	} else if win.Pressed(pixelgl.KeyD) {
		e.RotateRight()
	}

	if win.Pressed(pixelgl.KeyW) {
		e.Up()
	} else if win.Pressed(pixelgl.KeyS) {
		e.Down()
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

	if win.Pressed(pixelgl.KeySpacebar) {
		e.ResetMovement()
	}
}