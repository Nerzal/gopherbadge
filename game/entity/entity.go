package entity

import (
	"tinygo.org/x/drivers/pixel"
)

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

// ShouldBeCulled returns true if the entity has run off screen and is no longer relevant to the game
func (e *Entity) ShouldBeCulled() bool {
	return e.PosX+e.Width <= 0
}

// EnemyEntity extends the Entity by a collision flag.
type EnemyEntity struct {
	*Entity
	DidCollide    bool
	HasBeenScored bool
}

// Move pushes the enemy towards the left, based on the time which has passed since the last update and the current speed of the enemy
func (e *EnemyEntity) Move(deltaTime, movementSpeed float32) {
	e.Entity.PosX = e.Entity.PosX - movementSpeed*deltaTime
}

// HasBeenPassedByPlayer returns true, if the entity has fully passed the player
func (e *EnemyEntity) HasBeenPassedByPlayer(player *Entity) bool {
	return player.PosX > e.PosX+e.Width
}

// PlayerEntity extends the Entity to add jumping.
type PlayerEntity struct {
	*Entity
	currentYSpeed float32
}

const (
	InitialJumpSpeed   = 10
	Gravitation        = 5
	PlayerMinYPosition = 10
)

func (e *PlayerEntity) Jump() {
	if e.currentYSpeed != 0 {
		return
	}

	e.currentYSpeed = InitialJumpSpeed
}

func (e *PlayerEntity) Move(deltaTime float32) {
	e.PosY = min(e.PosY+e.currentYSpeed*deltaTime, PlayerMinYPosition)

	if e.PosY == PlayerMinYPosition {
		e.currentYSpeed = 0
	} else {
		e.currentYSpeed = e.currentYSpeed - Gravitation*deltaTime
	}
}
