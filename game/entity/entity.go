package entity

import (
	"math"

	"github.com/aykevl/tinygl/gfx"
	"github.com/conejoninja/gopherbadge/game/alias"
	"tinygo.org/x/drivers/pixel"
)

// Entity represents a game entity. Usable for dynamic and static objects.
type Entity struct {
	PosX          float32
	PosY          float32
	Width         float32
	Height        float32
	Image         alias.Image
	ScreenElement alias.Rect
}

var (
	black = pixel.NewRGB565BE(0, 0, 0)
)

func NewEntity(posX, posY, width, height float32, asset string) *Entity {
	println("creating entity")
	intPosX := int(math.Floor(float64(posX)))
	intPosY := int(math.Floor(float64(posY)))
	intWidth := int(math.Floor(float64(width)))
	intHeight := int(math.Floor(float64(height)))

	// image, err := image.NewQOI[pixel.RGB565BE](asset)
	// if err != nil {
	// 	println(err)
	// } else {
	// 	println("entity qoi successfully initialized")
	// }

	return &Entity{
		PosX:          posX,
		PosY:          posY,
		Width:         width,
		Height:        height,
		ScreenElement: alias.Rect{Rect: gfx.NewRect(black, intPosX, intPosY, intWidth, intHeight)},
		// Image:         alias.Image{Image: gfx.NewImage[pixel.RGB565BE](image, intPosX, intPosY)},
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

const (
	EnemySpawnDistance = 60
)

// EnemyEntity extends the Entity by a collision flag.
type EnemyEntity struct {
	*Entity
	DidCollide    bool
	HasBeenScored bool
}

// Move pushes the enemy towards the left, based on the time which has passed since the last update and the current speed of the enemy
func (e *EnemyEntity) Move(deltaTime, movementSpeed float32) {
	e.Entity.PosX = e.Entity.PosX - movementSpeed*deltaTime

	if e.ScreenElement.Rect != nil {
		e.ScreenElement.Move(int(e.PosX), int(e.PosY))
	}

	if e.Image.Image != nil {
		e.Image.Move(int(e.PosX), int(e.PosY))
	}
}

// HasBeenPassedByPlayer returns true, if the entity has fully passed the player
func (e *EnemyEntity) HasBeenPassedByPlayer(player *Entity) bool {
	return player.PosX > e.PosX+e.Width
}
