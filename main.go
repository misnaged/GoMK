package main

import (
	"GoMK/internal/core"
	"GoMK/internal/models"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "image/png"
	"log"
)

type Game struct {
	keys    []ebiten.Key
	subzira *models.SubzeroModel
	scene   models.Scene
}

func NewGame() *Game {
	return &Game{subzira: models.NewSubzeroModel()}
}
func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	g.subzira.Subzero.FramesCount++
	g.Move()

	if g.Push3() {
		fmt.Println(g.subzira.Subzero.X0, g.subzira.Subzero.X1, g.subzira.Subzero.Y0, g.subzira.Subzero.Y1, g.subzira.Subzero.FrameNum)
	}
	return nil
}

func (g *Game) SubzeroIdle(screen *ebiten.Image) error {
	g.subzira.Init()
	err := g.subzira.SubzeroIdle(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
func (g *Game) SubzeroMove(screen *ebiten.Image) error {
	err := g.subzira.SubzeroMvFw(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}
func (g *Game) SubzeroMoveBw(screen *ebiten.Image) error {
	err := g.subzira.SubzeroMvBw(screen)
	if err != nil {
		return fmt.Errorf("error while creating sub-zero %w", err)
	}
	return nil
}

// TODO: Refactor to make it looking good

func (g *Game) Move() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.subzira.Subzero.Idling = false
		g.subzira.Subzero.MovingBw = false
		g.subzira.Subzero.Moving = true
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		fmt.Println("released")
		g.subzira.Subzero.Idling = true
		g.subzira.Subzero.MovingBw = false
		g.subzira.Subzero.Moving = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.subzira.Subzero.Idling = false
		g.subzira.Subzero.Moving = false
		g.subzira.Subzero.MovingBw = true

	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		fmt.Println("released")
		g.subzira.Subzero.Idling = true
		g.subzira.Subzero.Moving = false
		g.subzira.Subzero.MovingBw = false
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
	err := g.scene.Screen(screen)
	if err != nil {
		panic(err)
	}
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
	if g.subzira.Subzero.MovingBw {
		err := g.SubzeroMoveBw(screen)
		if err != nil {
			panic(err)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return core.ScreenWidth, core.ScreenHeight
}

func main() {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	game := NewGame()
	game.subzira.Subzero.Idling = true
	ebiten.SetWindowSize(core.ScreenWidth, core.ScreenHeight)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
