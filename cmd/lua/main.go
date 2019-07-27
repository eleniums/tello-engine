package main

import (
	"flag"
	"log"
	"os"

	"github.com/eleniums/tello-engine/engine"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	flag.BoolVar(&engine.Debug, "debug", false, "true to enable debug logging")
	flag.Parse()

	// get script command-line argument
	script := os.Args[len(os.Args)-1]
	if _, err := os.Stat(script); len(os.Args) < 2 || os.IsNotExist(err) {
		log.Fatalln("Need to pass lua script as command-line argument")
	}

	// prepare to handle input and start the drone
	e := engine.NewEngine()
	e.AddKeyboardDevice()
	e.Start(false)
	defer e.Stop()

	// initialize lua engine
	L := lua.NewState()
	defer L.Close()

	// register engine functions with lua
	register(L, e)

	// execute script
	log.Println("----- Starting script -----")
	err := L.DoFile(script)
	if err != nil {
		log.Println("***** ERROR: INITIATING EMERGENCY LANDING *****")
		e.Stop()
		log.Fatalf("Error while running lua script: %v", err)
	}
	log.Println("----- Script completed -----")
}

// register functions with lua.
func register(L *lua.LState, e *engine.Engine) {
	sleep := func(L *lua.LState) int {
		milliseconds := L.ToInt(1)
		engine.Sleep(milliseconds)
		return 0
	}
	L.SetGlobal("sleep", L.NewFunction(sleep))

	logging := func(L *lua.LState) int {
		message := L.ToString(1)
		engine.Log(message)
		return 0
	}
	L.SetGlobal("log", L.NewFunction(logging))

	getLastKeyPress := func(L *lua.LState) int {
		lastKey := e.GetLastKeyPressed()
		ls := lua.LString(lastKey)
		L.Push(ls)
		return 1
	}
	L.SetGlobal("getlastkeypress", L.NewFunction(getLastKeyPress))

	startVideo := func(L *lua.LState) int {
		e.StartVideoStream()
		return 0
	}
	L.SetGlobal("startvideo", L.NewFunction(startVideo))

	takeOff := func(L *lua.LState) int {
		e.TakeOff()
		return 0
	}
	L.SetGlobal("takeoff", L.NewFunction(takeOff))

	land := func(L *lua.LState) int {
		e.Land()
		return 0
	}
	L.SetGlobal("land", L.NewFunction(land))

	forward := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Forward(speed)
		return 0
	}
	L.SetGlobal("forward", L.NewFunction(forward))

	backward := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Backward(speed)
		return 0
	}
	L.SetGlobal("backward", L.NewFunction(backward))

	left := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Left(speed)
		return 0
	}
	L.SetGlobal("left", L.NewFunction(left))

	right := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Right(speed)
		return 0
	}
	L.SetGlobal("right", L.NewFunction(right))

	up := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Up(speed)
		return 0
	}
	L.SetGlobal("up", L.NewFunction(up))

	down := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.Down(speed)
		return 0
	}
	L.SetGlobal("down", L.NewFunction(down))

	rotateLeft := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.RotateLeft(speed)
		return 0
	}
	L.SetGlobal("rotateleft", L.NewFunction(rotateLeft))

	rotateRight := func(L *lua.LState) int {
		speed := L.ToInt(1)
		e.RotateRight(speed)
		return 0
	}
	L.SetGlobal("rotateright", L.NewFunction(rotateRight))

	frontFlip := func(L *lua.LState) int {
		e.FrontFlip()
		return 0
	}
	L.SetGlobal("frontflip", L.NewFunction(frontFlip))

	backFlip := func(L *lua.LState) int {
		e.BackFlip()
		return 0
	}
	L.SetGlobal("backflip", L.NewFunction(backFlip))

	leftFlip := func(L *lua.LState) int {
		e.LeftFlip()
		return 0
	}
	L.SetGlobal("leftflip", L.NewFunction(leftFlip))

	rightFlip := func(L *lua.LState) int {
		e.RightFlip()
		return 0
	}
	L.SetGlobal("rightflip", L.NewFunction(rightFlip))

	bounce := func(L *lua.LState) int {
		e.Bounce()
		return 0
	}
	L.SetGlobal("bounce", L.NewFunction(bounce))

	stop := func(L *lua.LState) int {
		e.ResetMovement()
		return 0
	}
	L.SetGlobal("stop", L.NewFunction(stop))
}
