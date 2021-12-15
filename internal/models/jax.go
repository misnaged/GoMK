package models

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/Misnaged/spriteplus"
)

type Jax struct {
	anim service.Anim
}

var idlejax = []string{"jax/idle/1.png", "jax/idle/2.png", "jax/idle/3.png", "jax/idle/4.png", "jax/idle/5.png", "jax/idle/6.png", "jax/idle/7.png", "jax/idle/8.png", "jax/idle/9.png", "jax/idle/10.png", "jax/idle/11.png", "jax/idle/12.png", "jax/idle/13.png"}
var movejaxFW = []string{"jax/move/1.png", "jax/move/2.png", "jax/move/3.png", "jax/move/4.png", "jax/move/5.png", "jax/move/6.png", "jax/move/7.png", "jax/move/8.png", "jax/move/9.png"}
var movejaxBW = []string{"jax/move/9_1.png", "jax/move/10.png", "jax/move/11.png", "jax/move/12.png", "jax/move/13.png", "jax/move/14.png", "jax/move/15.png", "jax/move/16.png", "jax/move/17.png"}
var hitjaxpics = []string{"jax/hit/1.png", "jax/hit/2.png", "jax/hit/3.png", "jax/hit/4.png", "jax/hit/5.png", "jax/hit/6.png"}

func (jax *Jax) Idle() (idle *spriteplus.Animation, err error) {
	jax.anim.Pathpics = idlejax
	idle, err = jax.anim.BuildAnimation(80)
	if err != nil {
		return nil, fmt.Errorf("animation Idle failed due to: %v", err)
	}
	service.NewAnim(jax.anim.BuildPics(), jax.anim.BuildSprites(), idle, idlejax)

	return idle, nil
}

func (jax *Jax) MoveFW() (movefw *spriteplus.Animation, err error) {
	jax.anim.Pathpics = movejaxFW

	movefw, err = jax.anim.BuildAnimation(40)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(jax.anim.BuildPics(), jax.anim.BuildSprites(), movefw, movejaxFW)

	return movefw, nil
}
func (jax *Jax) MoveBW() (movebw *spriteplus.Animation, err error) {
	jax.anim.Pathpics = movejaxBW

	movebw, err = jax.anim.BuildAnimation(40)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(jax.anim.BuildPics(), jax.anim.BuildSprites(), movebw, movejaxBW)

	return movebw, nil
}
func (jax *Jax) HitJax() (hitjax *spriteplus.Animation, err error) {
	jax.anim.Pathpics = hitjaxpics

	hitjax, err = jax.anim.BuildAnimation(30)
	if err != nil {
		return nil, fmt.Errorf("animation Move failed due to: %v", err)
	}
	service.NewAnim(jax.anim.BuildPics(), jax.anim.BuildSprites(), hitjax, hitjaxpics)

	return hitjax, nil
}
