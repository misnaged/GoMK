package core

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"sync"
)

type Img struct {
	Img, SubZera                     *ebiten.Image
	DrawOpts, ResizeOpt, DefaultOpts ebiten.DrawImageOptions
	FrameNum, FramesCount, Kuka      int
	ImgBytes                         []byte
	Rect                             image.Rectangle
	Path                             []*image.Point
	X0, X1, Y0, Y1                   int
	animCount                        []int
	mux                              sync.Mutex
	Idling, Moving, MovingBw         bool
}

// PrepareImg is
func (img *Img) PrepareImg() (*ebiten.Image, error) {
	img2prep, _, err := image.Decode(bytes.NewReader(img.ImgBytes))
	if err != nil {
		return nil, fmt.Errorf("error while decoding %w", err)
	}
	img.Img = ebiten.NewImageFromImage(img2prep)
	if !done2 {
		img.DrawOpts.GeoM.Scale(2, 2)
	}
	return img.Img, nil
}

func (img *Img) ClearImg(image2clear *ebiten.Image, cleanCond bool) {
	if !cleanCond {
		image2clear.Clear()
	}
}
