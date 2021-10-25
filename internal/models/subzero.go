package models

import (
	"GoMK/internal/service"
	"fmt"
	"github.com/cebarks/spriteplus"
)


type SubZero struct{
	*service.Fighter
}
func (sz *SubZero) Idle() (idle *spriteplus.Animation, err error){
	sz.Pics = sz.BuildPics()
	sz.Sprites = sz.BuildSprites()
	idle, err =  sz.BuildAnimation(77); if err != nil{
		return nil, fmt.Errorf("idle animation failed with: %v", err)
	}
	return idle, nil
}


