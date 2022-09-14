package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

func (img *Img) DrawBackGround(screen *ebiten.Image) error {

	scene, err := img.PrepareImg()
	if err != nil {
		return fmt.Errorf("error while preparing image %w", err)
	}
	w, h := scene.Size()
	scaleW := ScreenWidth / float64(w)
	scaleH := ScreenHeight / float64(h)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(2.6, 3.2)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
	screen.DrawImage(scene, op)
	return nil
}
