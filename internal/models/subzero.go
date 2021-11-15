package models

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/cebarks/spriteplus"
)

type SubZero struct {
	anim service.Anim
}

var highkickpics = []string{"sz/highkick/01.png", "sz/highkick/02.png", "sz/highkick/03.png", "sz/highkick/04.png", "sz/highkick/05.png", "sz/highkick/06.png", "sz/highkick/07.png", "sz/highkick/08.png", "sz/highkick/09.png", "sz/highkick/10.png", "sz/highkick/11.png", "sz/highkick/12.png", "sz/highkick/13.png"}
var idlepics = []string{"sz/idle/02.png", "sz/idle/03.png", "sz/idle/04.png", "sz/idle/05.png", "sz/idle/06.png", "sz/idle/07.png", "sz/idle/08.png", "sz/idle/09.png", "sz/idle/10.png", "sz/idle/11.png", "sz/idle/12.png"}
var movepics = []string{"sz/move/01.png", "sz/move/02.png", "sz/move/03.png", "sz/move/04.png", "sz/move/05.png", "sz/move/06.png", "sz/move/07.png", "sz/move/08.png", "sz/move/09.png"}

func (sz *SubZero) Idle() (idle *spriteplus.Animation, err error) {
	idle, err = sz.anim.BuildAnimation(77)
	if err != nil {
		return nil, fmt.Errorf("animation Idle failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), idle, idlepics)

	return idle, nil
}

func (sz *SubZero) HighKick() (hgkick *spriteplus.Animation, err error) {
	hgkick, err = sz.anim.BuildAnimation(45)
	if err != nil {
		return nil, fmt.Errorf("animation HighKick failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), hgkick, highkickpics)

	return hgkick, nil
}
func (sz *SubZero) Move() (move *spriteplus.Animation, err error) {
	move, err = sz.anim.BuildAnimation(40)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(sz.anim.BuildPics(), sz.anim.BuildSprites(), move, movepics)

	return move, nil
}
