package engine

import (
	"log"
	"time"
)

// Sleep for the given number of milliseconds.
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// Log a message
func Log(line string) {
	log.Printf("(lua) %s", line)
}

// check is a helper method to reduce clutter and log errors.
func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
