package models

import (
	"GoMK/internal/core"
	"GoMK/sz"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	scene core.Img
}

func (s *Scene) Screen(screen *ebiten.Image) error {

	s.scene.ImgBytes = sz.Arena
	err := s.scene.DrawBackGround(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
