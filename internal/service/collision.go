package service

import (
	"fmt"
	"github.com/faiface/pixel"
)

type LineCollider struct {
	Collider pixel.Line
}

// DetectCollision is a trigger which reacts when 1 line intersect another
func (lc *LineCollider) DetectCollision(collider, collidee pixel.Line) (collided bool) {
	collider = lc.Collider

	_, ok := collider.Intersect(collidee)
	if ok {
		collided = true
		fmt.Println(collided)
	}
	return
}

// AttachColliderToParent attaches collider's vector to the given parent
// currently not using
func (lc *LineCollider) AttachColliderToParent(parenta, parentb pixel.Vec) (pixel.Vec, pixel.Vec) {
	parenta = lc.Collider.A
	parentb = lc.Collider.A.Sub(lc.Collider.B)

	return parenta, parentb
}
