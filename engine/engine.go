package engine

import (
	"log"
	"os/exec"
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
	// Speed of the drone for left, right, up, down, forward, and backward.
	Speed int

	// RotateSpeed of the drone.
	RotateSpeed int

	drone *tello.Driver
	robot *gobot.Robot
}

// NewEngine will create a new instance of the engine.
func NewEngine() *Engine {
	drone := tello.NewDriver("8888")
	keys := keyboard.NewDriver()

	e := Engine{
		Speed:       defaultSpeed,
		RotateSpeed: defaultRotateSpeed,
		drone:       drone,
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
	e.robot.Start(autoRun)
}

// work is the main method where actions are performed on the drone.
func (e *Engine) work() {
	// initialize mplayer at 60 fps (less fps causes lag in video)
	mplayer := exec.Command("mplayer", "-fps", "60", "-")
	mplayerIn, _ := mplayer.StdinPipe()
	if err := mplayer.Start(); err != nil {
		log.Println(err)
		return
	}

	// start video streaming from drone
	e.drone.On(tello.ConnectedEvent, func(data interface{}) {
		log.Println("Connected event received")
		e.drone.StartVideo()
		e.drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
		gobot.Every(100*time.Millisecond, func() {
			e.drone.StartVideo()
		})
	})

	// write video frames to mplayer
	e.drone.On(tello.VideoFrameEvent, func(data interface{}) {
		pkt := data.([]byte)
		if _, err := mplayerIn.Write(pkt); err != nil {
			log.Println(err)
		}
	})

	// get flight data events
	var flightData *tello.FlightData
	e.drone.On(tello.FlightDataEvent, func(data interface{}) {
		flightData = data.(*tello.FlightData)
	})

	// handle flight data events on a timed interval
	gobot.Every(5*time.Second, func() {
		if flightData != nil {
			log.Printf("Battery percentage: %v%%, Height: %v\n", flightData.BatteryPercentage, flightData.Height)
			if flightData.BatteryLow || flightData.BatteryLower {
				log.Println("***** WARNING: BATTERY LOW! *****")
			}
		}
	})
}
