package models

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"math/rand"
	"sync"
)

type Players struct {
	PlayerLeft  SubZero
	PlayerRight Jax
	mu          sync.Mutex
	glfw.Window
}

func (p *Players) DrawKick(win *pixelgl.Window) {
	g := &service.Geometry{}
	lv := g.GetLeftVector()
	var treesFrames []pixel.Rect

	pressed := service.Pressed(win, pixelgl.KeyR)
	//kick, err := p.PlayerLeft.HighKick()
	pic, err := service.LoadPicture("spritesheet.png")
	if err != nil {
		panic(err)
	}
	for x := pic.Bounds().Min.X; x < pic.Bounds().Max.X; x += 32 {
		for y := pic.Bounds().Min.Y; y < pic.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)
	var wg sync.WaitGroup

	if pressed {
		wg.Add(1)
		sprite := pixel.NewSprite(pic, treesFrames[rand.Intn(len(treesFrames))])

		fmt.Println("unlocked")
		sprite.Draw(batch, pixel.IM.Scaled(lv, 1.5).Moved(lv))
		batch.Draw(win)
		if !pressed {
			wg.Wait()
			goto Label1
		}
		wg.Done()
	}

	goto Label1

Label1:
	if pressed {
		fmt.Println("escaped!")
	}
}

func ReadShit(win *pixelgl.Window) {

	g := &service.Geometry{}
	lv := g.GetLeftVector()
	pic, err := service.LoadPicture("spritesheet.png")
	if err != nil {
		panic(err)
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, pic)
	//lr := g.GetLeftRect()
	//tree := pixel.NewSprite(pic, lr)
	sprite := pixel.NewSprite(pic, pic.Bounds())

	sprite.Draw(batch, pixel.IM.Scaled(lv, 1.5).Moved(lv))

}
