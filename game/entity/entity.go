package entity

import "tinygo.org/x/drivers/pixel"

// Entity represents a game entity. Usable for dynamic and static objects.
type Entity struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
	Image  pixel.Image[pixel.RGB565BE]
}

// Entity HasCollision checks
func (e *Entity) HasCollision(e2 *Entity) bool {
	if e.PosX+e.Width < e2.PosX {
		return false
	}

	if e.PosY > e2.PosY {
		return false
	}

	if e.PosY+e.Height < e2.PosY {
		return false
	}

	return true
}
