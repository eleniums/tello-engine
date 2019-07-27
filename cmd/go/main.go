package main

import (
	"flag"

	"github.com/eleniums/tello-engine/engine"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
		processKeys(win)
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

func processKeys(*pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		e.Left()
	}
	if win.Pressed(pixelgl.KeyRight) {
		e.Right()
	}
	if win.Pressed(pixelgl.KeyUp) {
		e.Forward()
	}
	if win.Pressed(pixelgl.KeyDown) {
		e.Backward()
	}
}