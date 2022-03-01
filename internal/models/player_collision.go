package models

import (
	"GoMK/internal/service"
	"github.com/faiface/pixel"
)

type PlayerCollision struct {
	player1col, player2col service.Collider
	geo                    *service.Geometry
}

func (pc *PlayerCollision) PlayerOneCollision() *service.Collider {
	lv := pc.geo.GetLeftVector()
	lRect := pc.geo.GetLeftRect()
	lv1 := pixel.Vec{X: 50, Y: 0}
	pc.player1col.Horizontal.A = lv
	pc.player1col.Horizontal.B = lv.Add(lv1)
	pc.player1col.Vertical.A = lv
	pc.player1col.Vertical.B.X = lv.X
	pc.player1col.Vertical.B.Y = lv.Y + 100
	pc.player1col.Box = lRect

	return &pc.player1col
}

func (pc *PlayerCollision) PlayerTwoCollision() *service.Collider {
	rv := pc.geo.GetRightVector()
	rRect := pc.geo.GetRightRect()
	pc.player2col.Horizontal.A = rv
	pc.player2col.Horizontal.B.X = rv.X - 50
	pc.player2col.Horizontal.B.Y = rv.Y
	pc.player2col.Vertical.A = rv
	pc.player2col.Vertical.B.X = rv.X
	pc.player2col.Vertical.B.Y = rv.Y + 100
	pc.player2col.Box = rRect

	return &pc.player2col
}

func (pc *PlayerCollision) LeftEdge(player *service.Collider) (colleft bool) {
	var wallleft pixel.Line
	wallleft.A = pixel.Vec{X: 0, Y: 100}
	wallleft.B = wallleft.A.Sub(pixel.Vec{X: -10, Y: -500})
	switch player {
	case pc.PlayerOneCollision():
		colleft = pc.player1col.DetectLineCollision(pc.player1col.Horizontal, wallleft) // left limits for player 1
		return colleft
	case pc.PlayerTwoCollision():
		colleft = pc.player2col.DetectLineCollision(pc.player2col.Horizontal, wallleft) // left limits for player 2
		return colleft
	}

	return
}

func (pc *PlayerCollision) RightEdge(player *service.Collider) (colright bool) {
	var wallright pixel.Line
	wallright.A = pixel.Vec{X: 2600, Y: 100}
	wallright.B = wallright.A.Add(pixel.Vec{X: 10, Y: 500})
	switch player {
	case pc.PlayerOneCollision():
		colright = pc.player1col.DetectLineCollision(pc.player1col.Horizontal, wallright) // right limits for player 1
		return colright

	case pc.PlayerTwoCollision():
		collright := pc.player2col.DetectLineCollision(pc.player2col.Horizontal, wallright) // right limits for player 2
		return collright
	}

	return
}
func (pc *PlayerCollision) Player1ToPlayer2() bool {
	pl1pl2 := pc.player1col.DetectLineCollision(pc.player1col.Horizontal, pc.player2col.Vertical)
	return pl1pl2
}
func (pc *PlayerCollision) Player2ToPlayer1() bool {
	pl2pl1 := pc.player2col.DetectLineCollision(pc.player2col.Horizontal, pc.player1col.Vertical)

	return pl2pl1
}

func (pc *PlayerCollision) DetectBox() bool {
	boxCollision := pc.player1col.DetectBoxCollision(pc.player1col.Box, pc.player2col.Box)
	return boxCollision
}
