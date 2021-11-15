package main

import (
	"GoMK/internal/models"
	"GoMK/internal/service"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	pic, err := service.LoadPicture("arena.png")
	if err != nil {
		panic(err)
	}

	Subzero := &models.SubZero{}

	kick, err := Subzero.HighKick()
	if err != nil {
		panic(err)
	}

	move, err := Subzero.Move()
	if err != nil {
		panic(err)
	}

	idle, err := Subzero.Idle()
	if err != nil {
		panic(err)
	}

	mv := pixel.Vec{X: 0, Y: 0}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	SwitchIdle := false
	SwitchHighKick := false
	SwitchMoveFw := false
	SwitchMoveBw := false

	for !win.Closed() {
		go sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 3.5).Moved(win.Bounds().Center()))

		if win.JustPressed(pixelgl.KeyR) {
			kick.CurrentSpriteIndex = 0
			SwitchHighKick = true
			SwitchMoveBw = false
			SwitchIdle = false
			SwitchMoveFw = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyR) {

		}

		if win.JustPressed(pixelgl.Key5) {
			kick.CurrentSpriteIndex = 0
			fmt.Println(" index:", kick.CurrentSpriteIndex)
		}

		if win.JustPressed(pixelgl.KeyD) {
			SwitchMoveFw = true
			SwitchMoveBw = false
			SwitchIdle = false
			SwitchHighKick = false

			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyD) {
			fmt.Println("RELEASED!")
			SwitchIdle = true
			SwitchMoveFw = false
			SwitchHighKick = false
			SwitchMoveBw = false
		}
		if win.JustPressed(pixelgl.KeyA) {
			SwitchMoveFw = false
			SwitchMoveBw = true
			SwitchIdle = false
			SwitchHighKick = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyA) {
			fmt.Println("RELEASED!")
			SwitchIdle = true
			SwitchMoveBw = false
			SwitchMoveFw = false
			SwitchHighKick = false
		}

		if SwitchHighKick == true {
			go kick.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))

		}
		if SwitchIdle == true {
			idle.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
		}

		if SwitchMoveFw == true {
			mv.X += 0.5

			move.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))

		}

		if SwitchMoveBw == true {
			mv.X -= 0.5

			move.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))

		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
