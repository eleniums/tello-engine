package main

import (
	"log"
	"os"

	"github.com/eleniums/tello-engine/engine"
	"github.com/yuin/gopher-lua"
)

func main() {
	// get script command-line argument
	if len(os.Args) < 2 {
		log.Fatalln("need to pass lua script as command-line argument")
	}
	script := os.Args[1]

	// prepare to handle input and start the drone
	e := engine.NewEngine()
	e.Start(false)
	defer e.Stop()

	// initialize lua engine and run the script
	L := lua.NewState()
	defer L.Close()

	// register engine functions with lua
	register(L, e)

	// execute script
	err := L.DoFile(script)
	if err != nil {
		log.Fatalf("error while running lua script: %v", err)
	}
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
		line := L.ToString(1)
		engine.Log(line)
		return 0
	}
	L.SetGlobal("log", L.NewFunction(logging))

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

	stop := func(L *lua.LState) int {
		e.ResetMovement()
		return 0
	}
	L.SetGlobal("stop", L.NewFunction(stop))
}
