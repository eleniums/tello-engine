package engine

import (
	"gobot.io/x/gobot/platforms/keyboard"
)

// GetLastKeyPressed will return the last key that was pressed or an empty string if none.
func (e *Engine) GetLastKeyPressed() string {
	lastKey := e.lastKeyPressed
	e.lastKeyPressed = ""
	return lastKey
}

// handleKeyPress handles a key press event.
func (e *Engine) handleKeyPress(key keyboard.KeyEvent) {
	// convert some common key presses into human readable strings
	var lastKeyPressed string
	switch key.Key {
	case keyboard.ArrowUp:
		lastKeyPressed = "arrowup"
	case keyboard.ArrowDown:
		lastKeyPressed = "arrowdown"
	case keyboard.ArrowLeft:
		lastKeyPressed = "arrowleft"
	case keyboard.ArrowRight:
		lastKeyPressed = "arrowright"
	case keyboard.Spacebar:
		lastKeyPressed = "spacebar"
	case keyboard.Escape:
		lastKeyPressed = "escape"
	default:
		lastKeyPressed = string(key.Key)
	}

	e.lastKeyPressed = lastKeyPressed

	debug("Key pressed: %s (%d)", lastKeyPressed, key.Key)
}
