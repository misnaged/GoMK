package service

import "github.com/faiface/pixel"

type Geometry struct {
	lv, rv       pixel.Vec
	lRect, rRect pixel.Rect
}

// GetLeftVector represents stats used to identify
// an `anchor` position for the Left Player
func (g *Geometry) GetLeftVector() pixel.Vec {
	g.lv = pixel.Vec{X: 400, Y: 100}
	return g.lv
}

// GetRightVector represents stats used to identify
// an `anchor` position for the Right Player
func (g *Geometry) GetRightVector() pixel.Vec {
	g.rv = pixel.Vec{X: 2000, Y: 100}
	return g.rv
}

func (g *Geometry) GetLeftRect() pixel.Rect {
	lvMin := pixel.Vec{X: 200, Y: 100}
	lvMax := pixel.Vec{X: 800, Y: 200}

	g.lRect = pixel.Rect{
		Min: lvMin,
		Max: lvMax,
	}
	return g.lRect
}

func (g *Geometry) GetRightRect() pixel.Rect {
	rvMin := pixel.Vec{X: 2200, Y: 100}
	rvMax := pixel.Vec{X: 1900, Y: 300}
	g.rRect = pixel.Rect{
		Min: rvMin,
		Max: rvMax,
	}
	return g.rRect
}
