package engine

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"gobot.io/x/gobot/platforms/keyboard"
)

const (
	defaultSpeed       = 40
	defaultRotateSpeed = 50
)

// Engine is the main structure for controlling a drone.
type Engine struct {
	drone *tello.Driver
	robot *gobot.Robot

	lastKeyPressed string
}

// NewEngine will create a new instance of the engine.
func NewEngine() *Engine {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()

	e := Engine{
		drone: drone,
	}

	keys.On(keyboard.Key, func(data interface{}) {
		e.handleKeyPress(data.(keyboard.KeyEvent))
	})

	e.robot = gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{keys, drone},
		func() {
			e.work()
		},
	)

	return &e
}

// Start the drone and allow input. If autoRun is true, this function will block.
func (e *Engine) Start(autoRun bool) {
	check(e.robot.Start(autoRun))
}

// Stop the drone.
func (e *Engine) Stop() {
	check(e.robot.Stop())
}

// work is the main method where actions are performed on the drone.
func (e *Engine) work() {
	// get flight data events
	var flightData *tello.FlightData
	e.drone.On(tello.FlightDataEvent, func(data interface{}) {
		flightData = data.(*tello.FlightData)
	})

	// handle flight data events on a timed interval
	gobot.Every(5*time.Second, func() {
		if flightData != nil {
			debug("Battery percentage: %v%%\n", flightData.BatteryPercentage)
			if flightData.BatteryLow || flightData.BatteryLower {
				log.Println("***** WARNING: BATTERY LOW! *****")
			}
		}
	})
}
