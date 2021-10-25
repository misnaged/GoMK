package service

import (
	"github.com/cebarks/spriteplus"
	"github.com/faiface/pixel"
	"image"
	"os"
)

// TODO IMPLEMENT INTERFACE
/*
// IFigther is an interface for each MK fighter
type IFigther interface {
	BuildPics(pics []pixel.Picture) ([]pixel.Picture)
	BuildSprites(sprites []*pixel.Sprite) ([]*pixel.Sprite)
	BuildAnimation(frames int) (anim *spriteplus.Animation, err error)
}
*/

// Fighter is a basic structure for MK's fighter
// All fighters should `inherit` this struct
type Fighter struct {
	Pics     []pixel.Picture
	Sprites  []*pixel.Sprite
	Anim     *spriteplus.Animation
	Pathpics []string
}

// BuildPics method is used to collect picture paths
// and combine them into []pixel.Picture slice
// to be used in BuildSprites method
func (f *Fighter) BuildPics() (pics []pixel.Picture) {
	for _, p := range f.Pathpics {
		pic, err := LoadPicture(p)
		if err != nil {
			panic(err)
		}
		pics = append(pics, pic)
	}
	return pics
}

// BuildSprites method is gathering the data([]pixel.Picture) from Pics method
// and `building` []*pixel.Sprite slice
// to be finally used in our BuildAnimation method
func (f *Fighter) BuildSprites() (sprites []*pixel.Sprite) {
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
func (f *Fighter) BuildAnimation(frames int) (anim *spriteplus.Animation, err error) {
	anim = f.Anim
	anim, err = spriteplus.MakeAnimation(f.BuildSprites(), frames)
	return anim, nil
}
