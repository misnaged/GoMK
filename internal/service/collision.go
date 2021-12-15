package service

import (
	"github.com/faiface/pixel"
)

type CrossCollider struct {
	Vertical   pixel.Line
	Horizontal pixel.Line
}

// ExpandLineCollider is used  when we need to perform some action that
// requires collision from one object to anot, such as hit, kick etc...
func (c *CrossCollider) ExpandLineCollider(h pixel.Line) {

	for i := 0; i <= 2; i++ {
		if i <= 0 {
			h.B.X += 50.0
		} else if i >= 2 {
			h.B.X -= 50.0

		}
	}

}

// DetectCollision is a trigger which reacts when 1 line intersect another
func (c *CrossCollider) DetectCollision(collider, collidee pixel.Line) (collided bool) {
	collider = c.Horizontal

	_, ok := collider.Intersect(collidee)
	if ok {
		collided = true
	}
	return
}
