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
	debug("Key pressed: %s (%d)", key.Char, key.Key)
	switch key.Key {
	case keyboard.ArrowUp:
		e.lastKeyPressed = "arrowup"
	case keyboard.ArrowDown:
		e.lastKeyPressed = "arrowdown"
	case keyboard.ArrowLeft:
		e.lastKeyPressed = "arrowleft"
	case keyboard.ArrowRight:
		e.lastKeyPressed = "arrowright"
	case keyboard.Spacebar:
		e.lastKeyPressed = "spacebar"
	case keyboard.Escape:
		e.lastKeyPressed = "escape"
	default:
		e.lastKeyPressed = key.Char
	}
}
