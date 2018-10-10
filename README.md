# tello-engine
Engine created in Go to control a Tello drone using Gobot, with Lua as a scripting language.

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
go run ./cmd/main.go
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
