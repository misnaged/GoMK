package service

import (
	"github.com/faiface/pixel/pixelgl"
)

type KeyPressed int

const (
	KeyLeft KeyPressed = iota
	KeyRight
	KeyUp
	KeyHighKick
)

func Pressed(win *pixelgl.Window, key pixelgl.Button) (pressed bool) {
	if win.JustPressed(key) {
		switch key {
		case pixelgl.KeyD:
			return true
		case pixelgl.KeyA:
			return true
		case pixelgl.KeyR:
			return true
		}
	}
	return false
}
