package entity

import (
	_ "embed"
)

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

	e.Image.Move(int(e.PosX), int(e.PosY))

}
