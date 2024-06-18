package entity

import (
	"github.com/aykevl/tinygl/gfx"
	"math"
	"tinygo.org/x/drivers/pixel"
)

// Entity represents a game entity. Usable for dynamic and static objects.
type Entity struct {
	PosX          float32
	PosY          float32
	Width         float32
	Height        float32
	Image         pixel.Image[pixel.RGB565BE]
	ScreenElement *gfx.Rect[pixel.RGB565BE]
}

var (
	black = pixel.NewRGB565BE(0, 0, 0)
)

func NewEntity(posX, posY, width, height float32) *Entity {
	intPosX := int(math.Floor(float64(posX)))
	intPosY := int(math.Floor(float64(posY)))
	intWidth := int(math.Floor(float64(width)))
	intHeight := int(math.Floor(float64(height)))
	return &Entity{
		PosX:          posX,
		PosY:          posY,
		Width:         width,
		Height:        height,
		ScreenElement: gfx.NewRect(black, intPosX, intPosY, intWidth, intHeight),
	}
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
	e.ScreenElement.Move(int(e.PosX), int(e.PosY))
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
	InitialJumpSpeed   = -20
	Gravitation        = -10
	PlayerMinYPosition = 160
)

func NewPlayer() *PlayerEntity {
	return &PlayerEntity{
		Entity: NewEntity(0, PlayerMinYPosition, 48, 96),
	}
}

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

	e.ScreenElement.Move(int(e.PosX), int(e.PosY))
}
