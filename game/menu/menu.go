package menu

import (
	"time"

	"github.com/aykevl/tinygl"
	"github.com/conejoninja/gopherbadge/game/alias"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinygl-font/roboto"
)

var (
	colors = []alias.Color{
		{pixel.NewRGB565BE(0, 0, 0)},
		{pixel.NewRGB565BE(255, 0, 0)},
		{pixel.NewRGB565BE(0, 255, 0)},
		{pixel.NewRGB565BE(0, 0, 255)},
	}
)

// Service is used to handle the menu logic and drawing everything.
type Service struct {
	canvas          alias.Canvas
	screen          alias.Screen
	buttonPressed   bool
	startTextColors []alias.Color
}

// New initializes a new menu service.
func New(canvas alias.Canvas, screen alias.Screen) *Service {
	return &Service{
		canvas:          canvas,
		screen:          screen,
		startTextColors: colors,
	}
}

func (s *Service) OnButtonPressed() {
	s.buttonPressed = true
}

func (s *Service) DrawStartMenu() {
	const (
		titleText = "Go forth"
	)

	s.canvas.Clear()

	var (
		black = pixel.NewColor[pixel.RGB565BE](0x00, 0x00, 0x00)
		white = pixel.NewColor[pixel.RGB565BE](0xff, 0xff, 0xff)
	)

	title := tinygl.NewText(roboto.Regular32, black, white, titleText)

	dummyText := tinygl.NewText(roboto.Regular32, black, white, "")
	dummyText.SetGrowable(1, 1)

	otherText := tinygl.NewText(roboto.Regular32, black, white, "Press A to start!")
	vbx := tinygl.NewVBox[pixel.RGB565BE](white, title, dummyText, otherText)

	s.screen.SetChild(vbx)
	s.screen.Update()

	go s.handleStartText(alias.Text{Text: otherText})
	s.waitForButton()
}

func (s *Service) handleStartText(text alias.Text) {
	for {
		if s.buttonPressed {
			return
		}

		for _, textColor := range s.startTextColors {
			s.animateStartText(text, textColor)
		}
	}
}

func (s *Service) animateStartText(text alias.Text, textColor alias.Color) {
	const startText = "Press A to start!"

	text.SetColor(textColor.RGB565BE)
	s.screen.Update()
	time.Sleep(400 * time.Millisecond)
}

func (s *Service) DrawGameOverMenu() {
	const (
		titleText = "Game Over - Get Go-od!"
	)

	s.canvas.Clear()

	var (
		black = pixel.NewColor[pixel.RGB565BE](0x00, 0x00, 0x00)
		white = pixel.NewColor[pixel.RGB565BE](0xff, 0xff, 0xff)
	)

	title := tinygl.NewText(roboto.Regular32, black, white, titleText)

	dummyText := tinygl.NewText(roboto.Regular32, black, white, "")
	dummyText.SetGrowable(1, 1)

	otherText := tinygl.NewText(roboto.Regular32, black, white, "Press A to start!")
	vbx := tinygl.NewVBox[pixel.RGB565BE](white, title, dummyText, otherText)

	s.screen.SetChild(vbx)
	s.screen.Update()

	go s.handleStartText(alias.Text{Text: otherText})
	s.waitForButton()
}

func (s *Service) waitForButton() {
	for {
		if s.buttonPressed {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}
