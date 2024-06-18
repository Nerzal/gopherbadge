package ui

import (
	"fmt"

	"github.com/aykevl/tinygl"
	"github.com/conejoninja/gopherbadge/game/alias"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinygl-font/roboto"
)

const (
	scoreFormat       = "score: %d"
	scorePositionX    = 8
	scorePositionY    = 16
	distancePositionX = 8
	distancePositionY = 32
	distanceFormat    = "distance: %.1f"
)

type Service struct {
	screen alias.Screen
}

func New(screen alias.Screen) *Service {
	return &Service{
		screen: screen,
	}
}

func (s *Service) DrawGameUi(screen, score int, distance float32) {
	var (
		black = pixel.NewColor[pixel.RGB565BE](0x00, 0x00, 0x00)
		white = pixel.NewColor[pixel.RGB565BE](0xff, 0xff, 0xff)
	)

	scoreText := tinygl.NewText(roboto.Regular16, black, white, fmt.Sprintf(scoreFormat, score))
	otherText := tinygl.NewText(roboto.Regular16, black, white, fmt.Sprintf(distanceFormat, score))
	vbx := tinygl.NewVBox[pixel.RGB565BE](white, scoreText, otherText)

	s.screen.SetChild(vbx)
	s.screen.Update()
}
