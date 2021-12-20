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

	// TODO Refract is needed in the future! Remove DRY-cases, etc
	mv := pixel.Vec{X: 400, Y: 100}
	mvMin := pixel.Vec{X: 200, Y: 100}
	mvMax := pixel.Vec{X: 800, Y: 200}

	zv := pixel.Vec{X: 2000, Y: 100}

	zvMin := pixel.Vec{X: 2200, Y: 100}
	zvMax := pixel.Vec{X: 1900, Y: 300}

	mv1 := pixel.Vec{X: 50, Y: 0}
	mRect := pixel.Rect{
		Min: mvMin,
		Max: mvMax,
	}
	zRect := pixel.Rect{
		Min: zvMin,
		Max: zvMax,
	}

	player1col := service.Collider{}
	player2col := service.Collider{}
	player1col.Horizontal.Scaled(1)
	player1col.Horizontal.A = mv
	player1col.Horizontal.B = mv.Add(mv1)
	player1col.Vertical.A = mv
	player1col.Vertical.B.X = mv.X
	player1col.Vertical.B.Y = mv.Y + 100
	player1col.Box = mRect

	player2col.Horizontal.A = zv
	player2col.Horizontal.B.X = zv.X - 50
	player2col.Horizontal.B.Y = zv.Y
	player2col.Vertical.A = zv
	player2col.Vertical.B.X = zv.X
	player2col.Vertical.B.Y = zv.Y + 100
	player2col.Box = zRect

	var wallright, wallleft pixel.Line

	wallleft.A = pixel.Vec{X: 0, Y: 100}
	wallleft.B = wallleft.A.Sub(pixel.Vec{X: -10, Y: -500})

	wallright.A = pixel.Vec{X: 2600, Y: 100}
	wallright.B = wallright.A.Add(pixel.Vec{X: 10, Y: 500})

	pic, err := service.LoadPicture("arena.png")
	if err != nil {
		panic(err)
	}

	//var wg sync.WaitGroup
	Subzero := &models.SubZero{}
	Jax := &models.Jax{}

	idlejx, err := Jax.Idle()
	if err != nil {
		panic(err)
	}
	hitjax, err := Jax.HitJax()
	if err != nil {
		panic(err)
	}
	movejaxfw, err := Jax.MoveFW()
	if err != nil {
		panic(err)
	}

	movejaxbw, err := Jax.MoveBW()
	if err != nil {
		panic(err)
	}

	kick, err := Subzero.HighKick()
	if err != nil {
		panic(err)
	}

	movefw, err := Subzero.MoveFW()
	if err != nil {
		panic(err)
	}

	movebw, err := Subzero.MoveBW()
	if err != nil {
		panic(err)
	}

	idle, err := Subzero.Idle()
	if err != nil {
		panic(err)
	}

	jumpUp, err := Subzero.JumpUp()
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	SwitchIdleJax := false
	SwitchJaxMoveFw := false
	SwitchJaxMoveBw := false
	SwitchIdle := false
	SwitchHighKick := false
	SwitchMoveFw := false
	SwitchMoveBw := false
	SwitchSzJumpedUp := false
	JaxWhaped := false

	// TODO: Jumping is buggy!

	for !win.Closed() {

		sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 3.5).Moved(win.Bounds().Center()))
		colleft := player1col.DetectLineCollision(player1col.Horizontal, wallleft)   // left limits for player 1
		colright := player1col.DetectLineCollision(player1col.Horizontal, wallright) // right limits for player 1

		pl2colleft := player2col.DetectLineCollision(player2col.Horizontal, wallleft)    // left limits for player 2
		pl2collright := player2col.DetectLineCollision(player2col.Horizontal, wallright) // right limits for player 2

		pl1pl2 := player1col.DetectLineCollision(player1col.Horizontal, player2col.Vertical)
		pl2pl1 := player2col.DetectLineCollision(player2col.Horizontal, player1col.Vertical)

		boxCollision := player1col.DetectBoxCollision(player1col.Box, player2col.Box)

		if boxCollision {
			fmt.Println("box collision!")
		}
		if win.JustPressed(pixelgl.KeyR) {
			kick.CurrentSpriteIndex = 0
			SwitchHighKick = true
			SwitchMoveBw = false
			SwitchIdle = false
			SwitchMoveFw = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyR) {
			SwitchHighKick = false
			SwitchIdle = true
		}

		//JAX--------------------//
		if win.JustPressed(pixelgl.KeyL) {
			SwitchJaxMoveBw = true
			SwitchJaxMoveFw = false
			SwitchIdleJax = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyL) {
			fmt.Println("RELEASED!")
			SwitchIdleJax = true
			SwitchJaxMoveFw = false
			SwitchJaxMoveBw = false
		}
		if win.JustPressed(pixelgl.KeyJ) {
			SwitchJaxMoveFw = true
			SwitchJaxMoveBw = false
			SwitchIdleJax = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyJ) {
			fmt.Println("RELEASED!")
			SwitchIdleJax = true
			SwitchJaxMoveFw = false
			SwitchJaxMoveBw = false
		}

		if win.JustPressed(pixelgl.Key5) {
			SwitchIdleJax = true
			SwitchIdle = true
			idlejx.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
			idle.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			win.Clear(colornames.Black)
		}

		if win.JustPressed(pixelgl.Key7) {
			fmt.Println("subziraColliderHorizontal", player1col.Horizontal)
			fmt.Println("subziraColliderVertical", player1col.Vertical)
			fmt.Println("JaxColVertical", player2col.Vertical)
			//fmt.Println("JaxBox", player2col.Box)
			//fmt.Println("SZbox", player1col.Box)
			fmt.Println("subziraPos", mv)
			fmt.Println("JaxPos", zv)
		}

		if win.JustPressed(pixelgl.KeyW) {
			jumpUp.CurrentSpriteIndex = 0
			SwitchSzJumpedUp = true
			SwitchMoveFw = false
			SwitchMoveBw = false
			SwitchIdle = false
			SwitchHighKick = false
			win.Clear(colornames.Black)
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

		if SwitchIdleJax == true {
			idlejx.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
		}
		if JaxWhaped == true {
			hitjax.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
		}
		if JaxWhaped == true && hitjax.CurrentSpriteIndex == 5 {
			JaxWhaped = false
		}
		if SwitchHighKick == true {
			kick.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			player1col.ExpandLineCollider(player1col.Horizontal)
			hitjax.CurrentSpriteIndex = 0

			if boxCollision && kick.CurrentSpriteIndex > 6 {
				JaxWhaped = true
				zv.X += 200.0
				zvMax.X += 200.0
				player2col.Horizontal.A.X += 200.0
				player2col.Horizontal.B.X += 200.0
				player2col.Vertical.A.X += 200.0
				player2col.Vertical.B.X += 200.0
				player2col.Box.Max.X += 200.0
				player2col.Box.Min.X += 200.0
			}
		}
		if SwitchIdle == true {
			idle.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))

		}

		if SwitchSzJumpedUp == true {
			jumpUp.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			if !colright && !pl1pl2 && jumpUp.CurrentSpriteIndex < 4 {
				mv.Y += 1.5
				player1col.Horizontal.B.Y += 1.5
				player1col.Vertical.B.Y += 1.5
			}
			if jumpUp.CurrentSpriteIndex == 4 {
				SwitchSzJumpedUp = false
				SwitchIdle = true
			}
			if SwitchSzJumpedUp != true && jumpUp.CurrentSpriteIndex == 4 {
				mv.Y = 100.0
				player1col.Horizontal.B.Y = 100.0
				player1col.Vertical.B.Y = 200.0
			}

		}

		if SwitchMoveFw == true {

			movefw.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			if !colright && !pl1pl2 {
				mv.X += 0.6
				player1col.Horizontal.A.X += 0.6
				player1col.Horizontal.B.X += 0.6
				player1col.Vertical.A.X += 0.6
				player1col.Vertical.B.X += 0.6
				player1col.Box.Max.X += 0.6
				player1col.Box.Min.X += 0.6
			}
		}

		if SwitchMoveBw == true {

			movebw.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			if !colleft {
				mv.X -= 0.6
				player1col.Horizontal.A.X -= 0.6
				player1col.Horizontal.B.X -= 0.6
				player1col.Vertical.A.X -= 0.6
				player1col.Vertical.B.X -= 0.6
				player1col.Box.Max.X -= 0.6
				player1col.Box.Min.X -= 0.6
			}
		}
		if SwitchJaxMoveFw == true {

			movejaxfw.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
			if !pl2colleft && !pl2pl1 {
				zv.X -= 0.6
				player2col.Horizontal.A.X -= 0.6
				player2col.Horizontal.B.X -= 0.6
				player2col.Vertical.A.X -= 0.6
				player2col.Vertical.B.X -= 0.6
				player2col.Box.Max.X -= 0.6
				player2col.Box.Min.X -= 0.6
			}
		}

		if SwitchJaxMoveBw == true {
			movejaxbw.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
			if !pl2collright {
				zv.X += 0.6
				player2col.Horizontal.A.X += 0.6
				player2col.Horizontal.B.X += 0.6
				player2col.Vertical.A.X += 0.6
				player2col.Vertical.B.X += 0.6
				player2col.Box.Max.X += 0.6
				player2col.Box.Min.X += 0.6
			}
		}

		win.Update()

	}
}

func main() {
	//defer profile.Start().Stop()
	//defer profile.Start(profile.MemProfile).Stop()
	pixelgl.Run(run)
}
