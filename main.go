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

	// TODO Refract is needed in the future! Remove DRY-cases, etc
	mv := pixel.Vec{X: 400, Y: 100}
	zv := pixel.Vec{X: 2000, Y: 100}
	mv1 := pixel.Vec{X: 10, Y: 10}

	player1col := service.LineCollider{}
	player2col := service.LineCollider{}
	player1col.Collider.A = mv
	player1col.Collider.B = mv.Add(mv1)

	player2col.Collider.A = zv
	player2col.Collider.B.X = zv.X - 10
	player2col.Collider.B.Y = zv.Y + 10

	var wallright, wallleft pixel.Line

	wallleft.A = pixel.Vec{X: 0, Y: 100}
	wallleft.B = wallleft.A.Sub(pixel.Vec{X: -10, Y: -500})

	wallright.A = pixel.Vec{X: 2600, Y: 100}
	wallright.B = wallright.A.Add(pixel.Vec{X: 10, Y: 500})

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
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

	sprite := pixel.NewSprite(pic, pic.Bounds())

	SwitchIdleJax := false
	SwitchJaxMoveFw := false
	SwitchJaxMoveBw := false
	SwitchIdle := false
	SwitchHighKick := false
	SwitchMoveFw := false
	SwitchMoveBw := false
	for !win.Closed() {

		sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 3.5).Moved(win.Bounds().Center()))
		//win.SwapBuffers()
		colleft := player1col.DetectCollision(player1col.Collider, wallleft)
		colright := player1col.DetectCollision(player1col.Collider, wallright)
		pl2colleft := player2col.DetectCollision(player2col.Collider, wallleft)
		pl2collright := player2col.DetectCollision(player2col.Collider, wallright)

		if win.JustPressed(pixelgl.KeyR) {
			kick.CurrentSpriteIndex = 0
			SwitchHighKick = true
			SwitchMoveBw = false
			SwitchIdle = false
			SwitchMoveFw = false
			win.Clear(colornames.Black)
		} else if win.JustReleased(pixelgl.KeyR) {

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
		if win.JustPressed(pixelgl.Key6) {
			zv.X += 400.0
			idlejx.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))

		}
		if win.JustPressed(pixelgl.Key7) {
			fmt.Println(player1col.Collider)
			fmt.Println(player2col.Collider)
			fmt.Println(zv)
			fmt.Println(mv)
			fmt.Println(wallleft)
			fmt.Println(wallright)

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

		if SwitchHighKick == true {
			kick.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))

		}
		if SwitchIdle == true {
			idle.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
		}

		if SwitchMoveFw == true {

			movefw.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			if !colright {
				mv.X += 0.6
				player1col.Collider.A.X += 0.6
				player1col.Collider.B.X += 0.6
			}
		}

		if SwitchMoveBw == true {

			movebw.Draw(win, pixel.IM.Scaled(mv, 1.5).Moved(win.GetPos().Add(mv)))
			if !colleft {
				mv.X -= 0.6
				player1col.Collider.A.X -= 0.6
				player1col.Collider.B.X -= 0.6
			}
		}
		if SwitchJaxMoveFw == true {

			movejaxfw.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
			if !pl2colleft {
				zv.X -= 0.6
				player2col.Collider.A.X -= 0.6
				player2col.Collider.B.X -= 0.6
			}
		}

		if SwitchJaxMoveBw == true {
			movejaxbw.Draw(win, pixel.IM.Scaled(zv, 1.5).Moved(win.GetPos().Add(zv)))
			if !pl2collright {
				zv.X += 0.6
				player2col.Collider.A.X += 0.6
				player2col.Collider.B.X += 0.6
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
