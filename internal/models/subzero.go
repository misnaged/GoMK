package models

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/Misnaged/spriteplus"
)

type SubZero struct {
	anim service.Anim
}

var highkickpics = []string{"sz/highkick/01.png", "sz/highkick/02.png", "sz/highkick/03.png", "sz/highkick/04.png", "sz/highkick/05.png", "sz/highkick/06.png", "sz/highkick/07.png", "sz/highkick/08.png", "sz/highkick/09.png", "sz/highkick/10.png", "sz/highkick/11.png", "sz/highkick/12.png", "sz/highkick/13.png"}
var idlepics = []string{"sz/idle/02.png", "sz/idle/03.png", "sz/idle/04.png", "sz/idle/05.png", "sz/idle/06.png", "sz/idle/07.png", "sz/idle/08.png", "sz/idle/09.png", "sz/idle/10.png", "sz/idle/11.png", "sz/idle/12.png"}
var movefwpics = []string{"sz/move/01.png", "sz/move/02.png", "sz/move/03.png", "sz/move/04.png", "sz/move/05.png", "sz/move/06.png", "sz/move/07.png", "sz/move/08.png", "sz/move/09.png"}
var movebwpics = []string{"sz/move/09_1.png", "sz/move/10.png", "sz/move/11.png", "sz/move/12.png", "sz/move/13.png", "sz/move/14.png", "sz/move/15.png", "sz/move/16.png", "sz/move/17.png"}
var jumpupPics = []string{"sz/jumpup/1.png", "sz/jumpup/2.png", "sz/jumpup/3.png", "sz/jumpup/4.png", "sz/jumpup/5.png"}

func (sz *SubZero) Idle() (idle *spriteplus.Animation, err error) {
	sz.anim.Pathpics = idlepics
	idle, err = sz.anim.BuildAnimation(70)
	if err != nil {
		return nil, fmt.Errorf("animation Idle failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), idle, idlepics)

	return idle, nil
}

func (sz *SubZero) JumpUp() (jumpUp *spriteplus.Animation, err error) {
	sz.anim.Pathpics = jumpupPics
	jumpUp, err = sz.anim.BuildAnimation(65)
	if err != nil {
		return nil, fmt.Errorf("animation Idle failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), jumpUp, jumpupPics)

	return jumpUp, nil
}

func (sz *SubZero) HighKick() (hgkick *spriteplus.Animation, err error) {
	sz.anim.Pathpics = highkickpics

	hgkick, err = sz.anim.BuildAnimation(15)
	if err != nil {
		return nil, fmt.Errorf("animation HighKick failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), hgkick, highkickpics)

	return hgkick, nil
}
func (sz *SubZero) MoveFW() (movefw *spriteplus.Animation, err error) {
	sz.anim.Pathpics = movefwpics

	movefw, err = sz.anim.BuildAnimation(40)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), movefw, movefwpics)

	return movefw, nil
}
func (sz *SubZero) MoveBW() (movebw *spriteplus.Animation, err error) {
	sz.anim.Pathpics = movebwpics

	movebw, err = sz.anim.BuildAnimation(40)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), movebw, movebwpics)

	return movebw, nil
}
