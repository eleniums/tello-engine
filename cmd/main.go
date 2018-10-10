package main

import (
	"github.com/eleniums/tello-engine/engine"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()

	keys.On(keyboard.Key, func(data interface{}) {
		engine.KeyHandler(drone, data.(keyboard.KeyEvent))
	})

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		func() {
			engine.Work(drone)
		},
	)

	robot.Start()
}
