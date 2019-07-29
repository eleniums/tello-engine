# tello-engine

[![GoDoc](https://godoc.org/github.com/eleniums/tello-engine?status.svg)](https://godoc.org/github.com/eleniums/tello-engine) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/eleniums/gohost/blob/master/LICENSE)

Engine created in Go to control a Tello drone using Gobot, with Lua as a scripting language. Yet another fun hackathon project.

## Installation
```
go get -u github.com/eleniums/tello-playground
```

Install MPlayer for video:
```
brew install mplayer
```

## Run
First, connect to the Tello drone via Wi-Fi. It will be named something similar to "TELLO-XXXXXX". Then run the program:
```
go run ./cmd/lua/main.go ./scripts/basic.lua
```

For a pure Go, non-scripted experience, run:
```
go run ./cmd/go/main.go
```

## Technologies
- DJI Tello drone:
    - https://store.dji.com/shop/tello-series
- Gobot framework with Tello driver:
    - https://gobot.io/documentation/platforms/tello
    - https://godoc.org/gobot.io/x/gobot/platforms/dji/tello
    - https://github.com/hybridgroup/gobot
- MPlayer:
    - http://www.mplayerhq.hu
- Pixel for keyboard input:
    - https://github.com/faiface/pixel

## Available lua commands
- `sleep(milliseconds)`: Sleep for the given number of milliseconds.
- `log(message)`: Log a message.
- `getlastkeypress()`: Return the last key that was pressed or an empty string if none.
- `startvideo()`: Initialize the video player and begin streaming from the drone.
- `takeoff()`: Start the blades and raise the drone to a normal flying height.
- `land()`: Lower the drone to the ground and stop the blades.
- `forward(speed)`: Move the drone forward at the given speed.
- `backward(speed)`: Move the drone backward at the given speed.
- `left(speed)`: Move the drone left at the given speed.
- `right(speed)`: Move the drone right at the given speed.
- `up(speed)`: Move the drone up at the given speed.
- `down(speed)`: Move the drone down at the given speed.
- `rotateleft(speed)`: Rotate the drone left at the given speed.
- `rotateright(speed)`: Rotate the drone right at the given speed.
- `frontflip()`: Command the drone to perform a front flip.
- `backflip()`: Command the drone to perform a back flip.
- `leftflip()`: : Command the drone to perform a left flip.
- `rightflip()`: : Command the drone to perform a right flip.
- `bounce()`: Toggle "bouncing" the drone up and down.
- `stop()`: Set all drone movement to 0.