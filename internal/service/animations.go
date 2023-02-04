package service

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/misnaged/spriteplus"
	"image"
	"os"
)

// IAnim is an interface for each MK Anim
type IAnim interface {
	BuildPics() []pixel.Picture
	BuildSprites() []*pixel.Sprite
	BuildAnimation(frames int) (anim *spriteplus.Animation, err error)
}

func NewAnim(pics []pixel.Picture, sprites []*pixel.Sprite, anim *spriteplus.Animation, pathpics []string) IAnim {
	return &Anim{
		Pics:       pics,
		Sprites:    sprites,
		AnimSource: anim,
		Pathpics:   pathpics,
	}
}

// Anim is a basic structure for MK's Anim
// All Anims should `inherit` this struct
type Anim struct {
	Pics       []pixel.Picture
	Sprites    []*pixel.Sprite
	AnimSource *spriteplus.Animation
	Pathpics   []string
}

func (f *Anim) GatherPics(pics []string) (ppics []string) {
	pics = ppics
	ppics = f.Pathpics
	return
}

// BuildPics method is used to collect picture paths
// and combine them into []pixel.Picture slice
// to be used in BuildSprites method
func (f *Anim) BuildPics() (pics []pixel.Picture) {
	for _, p := range f.Pathpics {
		pic, err := LoadPicture(p)
		if err != nil {
			panic(err)
		}
		pics = append(pics, pic)
	}
	return
}

// BuildSprites method is gathering the data([]pixel.Picture) from Pics method
// and `building` []*pixel.Sprite slice
// to be finally used in our BuildAnimation method
func (f *Anim) BuildSprites() (sprites []*pixel.Sprite) {
	sprites = f.Sprites
	for _, p := range f.BuildPics() {
		sprite := pixel.NewSprite(p, p.Bounds())
		sprites = append(sprites, sprite)
	}
	return
}

// LoadPicture load anim with path given
func LoadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// BuildAnimation method gets sprites array and builds proper animation
func (f *Anim) BuildAnimation(frames int) (anim *spriteplus.Animation, err error) {
	anim = f.AnimSource
	anim, err = spriteplus.MakeAnimation(f.BuildSprites(), frames)
	if err != nil {
		return nil, fmt.Errorf("animation failed due to: %v", err)
	}
	return anim, nil
}
