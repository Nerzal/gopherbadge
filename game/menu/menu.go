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
	display       *st7789.DeviceOf[pixel.RGB565BE]
	buttonPressed bool
}

// New initializes a new menu service.
func New(display st7789.DeviceOf[pixel.RGB565BE]) *Service {
	return &Service{
		display: &display,
	}
}

func (s *Service) OnButtonPressed() {
	s.buttonPressed = true
}

func (s *Service) DrawStartMenu() {
	const (
		titleText = "Go forth"
	)

	tinyfont.WriteLine(s.display, &freesans.Regular24pt7b, 80, 50, titleText, color.RGBA{0, 0, 0, 0})

	colors := []color.RGBA{
		color.RGBA{0, 0, 0, 0},
		color.RGBA{255, 0, 0, 0},
		color.RGBA{0, 255, 0, 0},
		color.RGBA{0, 0, 255, 0},
	}

	go func() {
		for {
			if s.buttonPressed {
				return
			}

			for _, textColor := range colors {
				s.animateStartText(textColor)

			}
		}
	}()

	for {
		if s.buttonPressed {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Service) animateStartText(textColor color.RGBA) {
	const startText = "Press A to start!"

	tinyfont.WriteLine(s.display, &freesans.Regular18pt7b, 20, 180, startText, textColor)
	time.Sleep(500 * time.Millisecond)
}
