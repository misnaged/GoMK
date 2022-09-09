package models

import (
	"GoMK/internal/core"
	"GoMK/sz"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Subzero struct {
	Subzero core.Img
}

func (g *Subzero) SubzeroIdle(screen *ebiten.Image) error {
	//
	var (
		frameOX     = 0
		frameOY     = 0
		frameWidth  = 82
		frameHeight = 150
		frameNum    = 11
	)
	g.Subzero.X0, g.Subzero.X1, g.Subzero.Y0, g.Subzero.Y1, g.Subzero.FrameNum = frameOX, frameWidth, frameOY, frameHeight, frameNum

	//
	g.Subzero.ImgBytes = sz.SubzeroIdle
	err := g.Subzero.DrawFighter(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
func (g *Subzero) SubzeroMvFw(screen *ebiten.Image) error {
	//
	var (
		frameOX     = 0
		frameOY     = 0
		frameWidth  = 82
		frameHeight = 150
		frameNum    = 9
	)
	g.Subzero.X0, g.Subzero.X1, g.Subzero.Y0, g.Subzero.Y1, g.Subzero.FrameNum = frameOX, frameWidth, frameOY, frameHeight, frameNum

	//
	g.Subzero.ImgBytes = sz.SubzeroMove
	err := g.Subzero.Move(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
