package main

import (
	"GoMK/internal/service"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/png"
)


func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync: false,
	}
	err := recover()
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	pic, err := service.LoadPicture("arena.png")
	if err != nil {
		panic(err)
	}

	var Vec2 pixel.Vec

	//coordX := 800.0
	//coordY := 800.0
	//var matrix = pixel.Matrix{1.5, 0, 0, 1.5, coordX, coordY}
	sprite := pixel.NewSprite(pic, pic.Bounds())


	subzero2 := &service.Fighter{}
	subzero2.Pathpics = []string{"sz/highkick/01.png","sz/highkick/02.png","sz/highkick/03.png","sz/highkick/04.png","sz/highkick/05.png","sz/highkick/06.png","sz/highkick/07.png","sz/highkick/08.png","sz/highkick/09.png","sz/highkick/10.png","sz/highkick/11.png","sz/highkick/12.png","sz/highkick/13.png"}

	subzero2.Pics = subzero2.BuildPics()
	subzero2.Sprites = subzero2.BuildSprites()
	kick, err :=  subzero2.BuildAnimation(65); if err != nil{
		panic(err)
	}


	// TODO all these things have to be implemented by interface using models for each fighters in Game
	// This implementation is very very dirt!
	subzero := &service.Fighter{}
	subzero.Pathpics = []string{"sz/idle/02.png","sz/idle/03.png", "sz/idle/04.png", "sz/idle/05.png", "sz/idle/06.png", "sz/idle/07.png", "sz/idle/08.png", "sz/idle/09.png","sz/idle/10.png", "sz/idle/11.png", "sz/idle/12.png"}

	subzero.Pics = subzero.BuildPics()
	subzero.Sprites = subzero.BuildSprites()
	idle, err :=  subzero.BuildAnimation(77); if err != nil{
		panic(err)
	}
	//

	SwitcherA := false
	SwitcherB := false
	SwitcherA = false
	SwitcherB = false
	
	for !win.Closed() {
		sprite.Draw(win, pixel.IM.Scaled(Vec2, 3.5).Moved(win.Bounds().Center()))

		if win.JustPressed(pixelgl.KeyR){
			SwitcherB = true
			SwitcherA = false
			win.Clear(colornames.Black)
		}

		if win.JustPressed(pixelgl.KeyF){
			SwitcherA = true
			SwitcherB = false
			win.Clear(colornames.Black)
		}
		if SwitcherB == true{
			 kick.Draw(win, pixel.IM.Scaled(Vec2, 1.5).Moved(win.Bounds().Center()))
		}

		if SwitcherA == true{
			 idle.Draw(win, pixel.IM.Scaled(Vec2, 1.5).Moved(win.Bounds().Center()))
		}
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}