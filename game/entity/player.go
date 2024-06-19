package entity

import (
	_ "embed"

	"github.com/conejoninja/gopherbadge/game/assets"
)

// PlayerEntity extends the Entity to add jumping.
type PlayerEntity struct {
	*Entity
	currentYSpeed float32
}

const (
	InitialJumpSpeed   = -200
	Gravitation        = -180
	PlayerMinYPosition = 120
)

func NewPlayer() *PlayerEntity {
	println("creating player entity")
	return &PlayerEntity{
		Entity: NewEntity(16, PlayerMinYPosition, 48, 96, assets.PlayerSprite1),
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

	if e.ScreenElement.Rect != nil {
		e.ScreenElement.Move(int(e.PosX), int(e.PosY))
	}

	if e.Image.Image != nil {
		e.Image.Move(int(e.PosX), int(e.PosY))
	}

}
