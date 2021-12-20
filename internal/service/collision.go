package service

import (
	"github.com/faiface/pixel"
	"time"
)

type Collider struct {
	Vertical   pixel.Line
	Horizontal pixel.Line
	Box        pixel.Rect
}

// ExpandLineCollider is used  when we need to perform some action that
// requires collision from one object to anot, such as hit, kick etc...
func (c *Collider) ExpandLineCollider(h pixel.Line) {
	for i := 0; i <= 2; i++ {
		if i <= 0 {
			h.B.X += 50.0
			time.Sleep(1 * time.Millisecond)
		} else if i >= 2 {
			h.B.X -= 50.0
			time.Sleep(1 * time.Millisecond)

		}
	}

}

// DetectBoxCollision is a trigger which reacts when 1 rect shape intersect another
func (c *Collider) DetectBoxCollision(collider, collidee pixel.Rect) (collided bool) {
	collider = c.Box

	ok := collider.Intersects(collidee)
	if ok {
		collided = true
	}
	return collided
}

// DetectLineCollision is a trigger which reacts when 1 line intersect another
func (c *Collider) DetectLineCollision(collider, collidee pixel.Line) (collided bool) {
	collider = c.Horizontal

	_, ok := collider.Intersect(collidee)
	if ok {
		collided = true
	}
	return collided
}
