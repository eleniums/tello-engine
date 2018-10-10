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

	// initialize lua engine and run the script
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile(script); err != nil {
		log.Fatalf("error while running lua script: %v", err)
	}
}
