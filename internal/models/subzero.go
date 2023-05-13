package models

import (
	"GoMK/internal/core"
	"GoMK/sz"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type SubzeroModel struct {
	Subzero *core.Img
}

func NewSubzeroModel() *SubzeroModel {
	return &SubzeroModel{Subzero: core.NewImg()}
}

func (g *SubzeroModel) Init() {
	var (
		frameOX     = 0
		frameOY     = 0
		frameWidth  = 82
		frameHeight = 152
	)
	g.Subzero.X0, g.Subzero.X1, g.Subzero.Y0, g.Subzero.Y1 = frameOX, frameWidth, frameOY, frameHeight
}
func (g *SubzeroModel) SubzeroIdle(screen *ebiten.Image) error {
	//
	g.Subzero.FrameNum = 11

	//
	g.Subzero.SetAnimationFramesLen(2, 5)

	g.Subzero.ImgBytes = sz.SubzeroIdle

	err := g.Subzero.LpIdle(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}

func (g *SubzeroModel) SubzeroMvFw(screen *ebiten.Image) error {
	//
	g.Subzero.FrameNum = 9

	g.Subzero.SetAnimationFramesLen(2, 5)

	//

	g.Subzero.ImgBytes = sz.SubzeroMove
	err := g.Subzero.LpMoveFw(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
func (g *SubzeroModel) SubzeroMvBw(screen *ebiten.Image) error {
	//
	g.Subzero.FrameNum = 9

	g.Subzero.SetAnimationFramesLen(2, 5)

	//

	g.Subzero.ImgBytes = sz.SubzeroMoveBw
	err := g.Subzero.LpMoveBw(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
