package main

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
)

//type SubZero struct{
//	Idle []*pixel.Sprite
//}


func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1600, 900),
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
	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite.Draw(win, pixel.IM.Scaled(Vec2, 3.5).Moved(win.Bounds().Center()))
	fmt.Println("sprite", sprite)

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

	 var back pixel.Vec
	back.X = -1.0
	back.Y = 1.0


	for !win.Closed() {
		win.Update()
		idle.Draw(win, pixel.IM.Scaled(Vec2, 1.5).Moved(win.Bounds().Center()))
	}
}
func main() {
	pixelgl.Run(run)
}