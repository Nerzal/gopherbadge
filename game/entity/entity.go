package entity

import (
	"tinygo.org/x/drivers/pixel"
)

// EnemyEntity extends the Entity by a collision flag.
type EnemyEntity struct {
	*Entity
	DidCollide    bool
	HasBeenScored bool
}

// Entity represents a game entity. Usable for dynamic and static objects.
type Entity struct {
	PosX   float32
	PosY   float32
	Width  float32
	Height float32
	Image  pixel.Image[pixel.RGB565BE]
}

// HasCollision checks for collisions between entities based on position and size
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

// Move pushes the enemy towards the left, based on the time which has passed since the last update and the current speed of the enemy
func (e *EnemyEntity) Move(deltaTime, movementSpeed float32) {
	e.Entity.PosX = e.Entity.PosX - movementSpeed*deltaTime
}

func (e *Entity) ShouldBeCulled() bool {
	return e.PosX+e.Width <= 0
}

func (e *EnemyEntity) HasBeenPassedByPlayer(player *Entity) bool {
	return player.PosX > e.PosX+e.Width
}
