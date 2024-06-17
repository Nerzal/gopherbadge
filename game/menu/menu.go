package menu

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

// Service is used to handle the menu logic and drawing everything.
type Service struct {
	display         *st7789.DeviceOf[pixel.RGB565BE]
	buttonPressed   bool
	startTextColors []color.RGBA
}

// New initializes a new menu service.
func New(display st7789.DeviceOf[pixel.RGB565BE]) *Service {
	colors := []color.RGBA{
		color.RGBA{0, 0, 0, 0},
		color.RGBA{255, 0, 0, 0},
		color.RGBA{0, 255, 0, 0},
		color.RGBA{0, 0, 255, 0},
	}

	return &Service{
		display:         &display,
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

	s.display.FillScreen(color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(s.display, &freesans.Regular24pt7b, 80, 50, titleText, color.RGBA{0, 0, 0, 0})

	go s.handleStartText()
	s.waitForButton()
}

func (s *Service) handleStartText() {
	for {
		if s.buttonPressed {
			return
		}

		for _, textColor := range s.startTextColors {
			s.animateStartText(textColor)
		}
	}
}

func (s *Service) animateStartText(textColor color.RGBA) {
	const startText = "Press A to start!"

	tinyfont.WriteLine(s.display, &freesans.Regular18pt7b, 20, 180, startText, textColor)
	time.Sleep(500 * time.Millisecond)
}

func (s *Service) DrawGameOverMenu() {
	const (
		titleText = "Game Over - Get Go-od!"
	)

	s.display.FillScreen(color.RGBA{255, 255, 255, 255})

	tinyfont.WriteLine(s.display, &freesans.Regular24pt7b, 80, 50, titleText, color.RGBA{0, 0, 0, 0})

	go s.handleStartText()
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
