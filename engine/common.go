package engine

import (
	"log"
	"time"
)

// Debug flag. Set to true to view debug logs.
var Debug = false

// Sleep for the given number of milliseconds.
func Sleep(milliseconds int) {
	debug("Sleep for %vms", milliseconds)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Log a message.
func Log(message string) {
	log.Printf("(lua) %s", message)
}

// debug will write a debug log.
func debug(format string, v ...interface{}) {
	if Debug {
		log.Printf(format, v...)
	}
}

// check is a helper method to reduce clutter and log errors.
func check(err error) {
	if err != nil {
		log.Printf("Error occurred: %v", err)
	}
}
