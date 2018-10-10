package main

import (
	"log"
	"os"

	"github.com/eleniums/tello-engine/engine"
	"github.com/yuin/gopher-lua"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	// get script command-line argument
	if len(os.Args) < 2 {
		log.Fatalln("need to pass lua script as command-line argument")
	}
	script := os.Args[1]

	// prepare to handle input and start the drone
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

	robot.Start(false)

	// initialize lua engine and run the script
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(script); err != nil {
		log.Fatalf("error while running lua script: %v", err)
	}
}
