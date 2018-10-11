package engine

import (
	"time"
)

// Sleep for the given number of milliseconds.
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
