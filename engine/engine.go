package engine

import (
	"log"
	"os/exec"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

const (
	// Speed of the drone for left, right, up, down, forward, and backward.
	Speed = 40

	// RotateSpeed of the drone.
	RotateSpeed = 50
)

// Work is the main method where actions are performed on the drone.
func Work(drone *tello.Driver) {
	// initialize mplayer at 60 fps (less fps causes lag in video)
	mplayer := exec.Command("mplayer", "-fps", "60", "-")
	mplayerIn, _ := mplayer.StdinPipe()
	if err := mplayer.Start(); err != nil {
		log.Println(err)
		return
	}

	// start video streaming from drone
	drone.On(tello.ConnectedEvent, func(data interface{}) {
		log.Println("Connected event received")
		drone.StartVideo()
		drone.SetVideoEncoderRate(tello.VideoBitRateAuto)
		gobot.Every(100*time.Millisecond, func() {
			drone.StartVideo()
		})
	})

	// write video frames to mplayer
	drone.On(tello.VideoFrameEvent, func(data interface{}) {
		pkt := data.([]byte)
		if _, err := mplayerIn.Write(pkt); err != nil {
			log.Println(err)
		}
	})

	// get flight data events
	var flightData *tello.FlightData
	drone.On(tello.FlightDataEvent, func(data interface{}) {
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

// resetDroneMovement will set all drone movement to 0.
func resetDroneMovement(drone *tello.Driver) {
	log.Println("Reset drone position")
	drone.Forward(0)
	drone.Backward(0)
	drone.Up(0)
	drone.Down(0)
	drone.Left(0)
	drone.Right(0)
	drone.Clockwise(0)
}
