package core

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"sync"
)

func (img *Img) GetSx() (int, int) {
	i := (img.FramesCount / len(img.animCount)) % img.FrameNum
	sx, sy := img.X0+i*img.X1, img.Y0-i
	return sx, sy
}

func (img *Img) SetAnimationFramesLen(num, cap int) {
	img.animCount = make([]int, num, cap)
	for i := 0; i <= num; i++ {
		img.animCount = append(img.animCount, i)
	}
	if num > cap {
		num = 0
	}
}

var done1, done2 bool

func (img *Img) LpPos() *image.Rectangle {
	if !done1 {
		img.DrawOpts.GeoM.Translate(-float64(img.X1)*3.5, -float64(img.Y1)+20.5)
		img.DrawOpts.GeoM.Translate(ScreenWidth/2, ScreenHeight/1.3)
	}
	sx, sy := img.GetSx()
	img.Rect = image.Rect(sx, sy, sx+img.X1, sy+img.Y1)

	return &img.Rect
}

func (img *Img) LpIdle(screen *ebiten.Image) error {
	fighter, err := img.PrepareImg()
	done2 = true
	if err != nil {
		return fmt.Errorf("error while preparing image %w", err)
	}
	var wg sync.WaitGroup
	for range img.animCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			subImg := fighter.SubImage(*img.LpPos()).(*ebiten.Image)
			screen.DrawImage(subImg, &img.DrawOpts)
			done1 = true
		}()
		wg.Wait()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		fmt.Println(img.Rect)
	}
	return nil
}

func (img *Img) LpMoveFw(screen *ebiten.Image) error {
	fighter, err := img.PrepareImg()
	if err != nil {
		return fmt.Errorf("error while preparing image %w", err)
	}
	var wg sync.WaitGroup
	img.SubZera = fighter.SubImage(*img.LpPos()).(*ebiten.Image)
	for range img.animCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer img.SubZera.Clear()
			img.DrawOpts.GeoM.Translate(3.0/float64(len(img.animCount))*1.5, 0)
			screen.DrawImage(img.SubZera, &img.DrawOpts)
		}()
		wg.Wait()
	}

	return nil
}

func (img *Img) RightDirPath() []*image.Point {
	p1 := &image.Point{
		X: img.X0 + 1,
		Y: 0,
	}
	p2 := &image.Point{
		X: img.X0 + 1,
		Y: 0,
	}
	img.Path = append(img.Path, p1, p2)
	return img.Path
}

// ----- Position ----- //

// GetCurrent is
func (img *Img) GetCurrent() *image.Rectangle {
	return &img.Rect
}

/*
op := &ebiten.DrawImageOptions{}
op.GeoM.Translate(-float64(img.X1)/2, -float64(img.Y1)/2)
op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
i := (img.FramesCount / 5) % img.FrameNum
sx, sy := img.X0+i*img.X1, img.Y0
img.Rect = image.Rect(img.posX, sy, sx+img.X1, sy+img.Y1)
fighter, err := img.PrepareImg()

	if err != nil {
		return fmt.Errorf("error while preparing image %w", err)
	}

screen.DrawImage(fighter.SubImage(img.Rect).(*ebiten.Image), op)

	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		fmt.Println("posX =", img.posX)
	}	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {

	fmt.Printf("\n\n\n\n\n\n\n %d \n", len(img.animCount))
	fmt.Printf("max.x: %v, max.y: %v \n", img.Rect.Max.X, img.Rect.Max.Y)
	fmt.Printf("min.x: %v, min.y: %v \n", img.Rect.Min.X, img.Rect.Min.Y)

}
*/
