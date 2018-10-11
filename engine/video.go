package engine

import (
	"log"
	"os/exec"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

// StartVideoStream will initialize the video player and begin streaming from the drone.
func (e *Engine) StartVideoStream() {
	// initialize mplayer at 60 fps (less fps causes lag in video)
	mplayer := exec.Command("mplayer", "-fps", "60", "-")
	mplayerIn, err := mplayer.StdinPipe()
	if err != nil {
		log.Printf("Error creating pipe for video player: %v", err)
		return
	}
	if err := mplayer.Start(); err != nil {
		log.Printf("Error starting video player: %v", err)
		return
	}

	// start video streaming from drone
	e.drone.On(tello.ConnectedEvent, func(data interface{}) {
		log.Println("Connected and starting video stream")
		check(e.drone.StartVideo())
		check(e.drone.SetVideoEncoderRate(tello.VideoBitRateAuto))
		gobot.Every(100*time.Millisecond, func() {
			check(e.drone.StartVideo())
		})
	})

	// write video frames to mplayer
	e.drone.On(tello.VideoFrameEvent, func(data interface{}) {
		pkt := data.([]byte)
		if _, err := mplayerIn.Write(pkt); err != nil {
			log.Printf("Error writing to video: %v", err)
		}
	})
}
