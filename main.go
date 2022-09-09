package main

import (
	"GoMK/internal/core"
	"GoMK/internal/models"

	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	subzira models.Subzero
}

func (g *Game) Update() error {
	g.subzira.Subzero.FramesCount++
	if g.Push() {
		g.subzira.Subzero.Idling = false
		g.subzira.Subzero.Moving = true
	}
	if g.Push2() {
		g.subzira.Subzero.Moving = false
		g.subzira.Subzero.Idling = true
	}
	if g.Push3() {
		fmt.Println(g.subzira.Subzero.X0, g.subzira.Subzero.X1, g.subzira.Subzero.Y0, g.subzira.Subzero.Y1, g.subzira.Subzero.FrameNum)
	}
	return nil
}

func (g *Game) SubzeroIdle(screen *ebiten.Image) error {
	err := g.subzira.SubzeroIdle(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
func (g *Game) SubzeroMove(screen *ebiten.Image) error {
	_ = g.subzira.SubzeroMvFw(screen)
	err := g.subzira.Subzero.Move(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}

func (g *Game) Push() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	} else {
		return false
	}
}
func (g *Game) Push2() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyAltLeft) {
		return true
	} else {
		return false
	}
}
func (g *Game) Push3() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyControlLeft) {
		return true
	} else {
		return false
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.subzira.Subzero.Idling {
		err := g.SubzeroIdle(screen)
		if err != nil {
			panic(err)
		}
	}
	if g.subzira.Subzero.Moving {
		err := g.SubzeroMove(screen)
		if err != nil {
			panic(err)
		}
	}
	//screen.Clear()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return core.ScreenWidth, core.ScreenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	ebiten.SetWindowSize(core.ScreenWidth*2, core.ScreenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
