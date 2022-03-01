package main

import (
	"GoMK/internal/models"
	"GoMK/internal/service"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	_ "image/png"
	"log"
	"net/http"
	_ "net/http/pprof"
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

	sprite := pixel.NewSprite(pic, pic.Bounds())

	p := &models.Players{}

	for !win.Closed() {
		sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 3.5).Moved(win.Bounds().Center()))

		p.DrawKick(win)
		win.Update()
	}
}
func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	fmt.Println("hello")
	//go tool pprof http://localhost:6060/debug/pprof/heap//

	//defer profile.Start().Stop()
	//defer profile.Start(profile.MemProfile).Stop()
	pixelgl.Run(run)
}

/*

	type Players struct {
		PlayerLeft *models.SubZero
		PlayerRight *models.Jax
	}



	func PlayerLeft() *models.SubZero {
		Subzero := &models.SubZero{}


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
		return Subzero
	}


	//var wg sync.WaitGroup

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




	SwitchIdleJax := false
	SwitchJaxMoveFw := false
	SwitchJaxMoveBw := false
	SwitchIdle := false
	SwitchHighKick := false
	SwitchMoveFw := false
	SwitchMoveBw := false
	SwitchSzJumpedUp := false
	JaxWhaped := false





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
		idlejx.Draw(win, pixel.IM.Scaled(rv, 1.5).Moved(win.GetPos().Add(rv)))
		idle.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))
		win.Clear(colornames.Black)
	}

	if win.JustPressed(pixelgl.Key7) {
		fmt.Println("subziraColliderHorizontal", player1col.Horizontal)
		fmt.Println("subziraColliderVertical", player1col.Vertical)
		fmt.Println("JaxColVertical", player2col.Vertical)
		//fmt.Println("JaxBox", player2col.Box)
		//fmt.Println("SZbox", player1col.Box)
		fmt.Println("subziraPos", lv)
		fmt.Println("JaxPos", rv)
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
		idlejx.Draw(win, pixel.IM.Scaled(rv, 1.5).Moved(win.GetPos().Add(rv)))
	}
	if JaxWhaped == true {
		hitjax.Draw(win, pixel.IM.Scaled(rv, 1.5).Moved(win.GetPos().Add(rv)))
	}
	if JaxWhaped == true && hitjax.CurrentSpriteIndex == 5 {
		JaxWhaped = false
	}
	if SwitchHighKick == true {
		kick.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))
		player1col.ExpandLineCollider(player1col.Horizontal)
		hitjax.CurrentSpriteIndex = 0

		if boxCollision && kick.CurrentSpriteIndex > 6 {
			JaxWhaped = true
			rv.X += 200.0
			rvMax.X += 200.0
			player2col.Horizontal.A.X += 200.0
			player2col.Horizontal.B.X += 200.0
			player2col.Vertical.A.X += 200.0
			player2col.Vertical.B.X += 200.0
			player2col.Box.Max.X += 200.0
			player2col.Box.Min.X += 200.0
		}
	}
	if SwitchIdle == true {
		idle.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))

	}

	if SwitchSzJumpedUp == true {
		jumpUp.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))
		if !colright && !pl1pl2 && jumpUp.CurrentSpriteIndex < 4 {
			lv.Y += 1.5
			player1col.Horizontal.B.Y += 1.5
			player1col.Vertical.B.Y += 1.5
		}
		if jumpUp.CurrentSpriteIndex == 4 {
			SwitchSzJumpedUp = false
			SwitchIdle = true
		}
		if SwitchSzJumpedUp != true && jumpUp.CurrentSpriteIndex == 4 {
			lv.Y = 100.0
			player1col.Horizontal.B.Y = 100.0
			player1col.Vertical.B.Y = 200.0
		}

	}

	if SwitchMoveFw == true {

		movefw.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))
		if !colright && !pl1pl2 {
			lv.X += 0.6
			player1col.Horizontal.A.X += 0.6
			player1col.Horizontal.B.X += 0.6
			player1col.Vertical.A.X += 0.6
			player1col.Vertical.B.X += 0.6
			player1col.Box.Max.X += 0.6
			player1col.Box.Min.X += 0.6
		}
	}

	if SwitchMoveBw == true {

		movebw.Draw(win, pixel.IM.Scaled(lv, 1.5).Moved(win.GetPos().Add(lv)))
		if !colleft {
			lv.X -= 0.6
			player1col.Horizontal.A.X -= 0.6
			player1col.Horizontal.B.X -= 0.6
			player1col.Vertical.A.X -= 0.6
			player1col.Vertical.B.X -= 0.6
			player1col.Box.Max.X -= 0.6
			player1col.Box.Min.X -= 0.6
		}
	}
	if SwitchJaxMoveFw == true {

		movejaxfw.Draw(win, pixel.IM.Scaled(rv, 1.5).Moved(win.GetPos().Add(rv)))
		if !pl2colleft && !pl2pl1 {
			rv.X -= 0.6
			player2col.Horizontal.A.X -= 0.6
			player2col.Horizontal.B.X -= 0.6
			player2col.Vertical.A.X -= 0.6
			player2col.Vertical.B.X -= 0.6
			player2col.Box.Max.X -= 0.6
			player2col.Box.Min.X -= 0.6
		}
	}

	if SwitchJaxMoveBw == true {
		movejaxbw.Draw(win, pixel.IM.Scaled(rv, 1.5).Moved(win.GetPos().Add(rv)))
		if !pl2collright {
			rv.X += 0.6
			player2col.Horizontal.A.X += 0.6
			player2col.Horizontal.B.X += 0.6
			player2col.Vertical.A.X += 0.6
			player2col.Vertical.B.X += 0.6
			player2col.Box.Max.X += 0.6
			player2col.Box.Min.X += 0.6
		}
	}
}
*/
