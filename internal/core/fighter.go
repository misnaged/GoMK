package core

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

//	type Fighter struct {
//		Img *Img
//	}
type Img struct {
	FigtherImg            *ebiten.Image
	DrawOpts              *ebiten.DrawImageOptions
	FrameNum, FramesCount int
	ImgBytes              []byte
	Rect                  image.Rectangle
	X0, X1, Y0, Y1        int
	Idling, Moving        bool
}

//type Position struct {
//	Rect           image.Rectangle
//	x0, x1, y0, y1 int
//}

//	func NewFighter(fighter *Fighter) IFighter{
//		return fighter
//	}

// PrepareImg is
func (img *Img) PrepareImg() (*ebiten.Image, error) {
	img2prep, _, err := image.Decode(bytes.NewReader(img.ImgBytes))
	if err != nil {
		return nil, fmt.Errorf("error while decoding %w", err)
	}
	img.FigtherImg = ebiten.NewImageFromImage(img2prep)
	return img.FigtherImg, nil
}

// DrawFighter is
func (img *Img) DrawFighter(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(img.X1)/2, -float64(img.Y1)/2)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
	i := (img.FramesCount / 5) % img.FrameNum
	sx, sy := img.X0+i*img.X1, img.Y0
	img.Rect = image.Rect(sx, sy, sx+img.X1, sy+img.Y1)
	fighter, err := img.PrepareImg()
	if err != nil {
		return fmt.Errorf("error while preparing image %w", err)
	}
	screen.DrawImage(fighter.SubImage(img.Rect).(*ebiten.Image), op)
	return nil
}
func (img *Img) Move(screen *ebiten.Image) error {
	screen.Clear()
	op := &ebiten.DrawImageOptions{}
	if img.Moving {

		op.GeoM.Translate(-float64(img.X1)/2, -float64(img.Y1)+float64(1+1))
		op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)

		i := (img.FramesCount / 5) % img.FrameNum
		sx, sy := img.X0+i*img.X1, img.Y0
		img.Rect = image.Rect(sx, sy, sx+img.X1, sy+img.Y1)
		fighter, err := img.PrepareImg()
		if err != nil {
			return fmt.Errorf("error while preparing image %w", err)
		}

		screen.DrawImage(fighter.SubImage(img.Rect).(*ebiten.Image), op)
	}
	return nil
}

// ----- Position ----- //

// GetCurrent is
func (img *Img) GetCurrent() *image.Rectangle {
	return &img.Rect
}
