package menu

import (
	"time"

	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/gfx"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinygl-font/roboto"
)

// Service is used to handle the menu logic and drawing everything.
type Service struct {
	canvas          *gfx.Canvas[pixel.RGB565BE]
	screen          *tinygl.Screen[pixel.RGB565BE]
	buttonPressed   bool
	startTextColors []pixel.RGB565BE
}

// New initializes a new menu service.
func New(canvas *gfx.Canvas[pixel.RGB565BE], screen *tinygl.Screen[pixel.RGB565BE]) *Service {
	colors := []pixel.RGB565BE{
		pixel.NewRGB565BE(0, 0, 0),
		pixel.NewRGB565BE(255, 0, 0),
		pixel.NewRGB565BE(0, 255, 0),
		pixel.NewRGB565BE(0, 0, 255),
	}

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

	go s.handleStartText(otherText)
	s.waitForButton()
}

func (s *Service) handleStartText(text *tinygl.Text[pixel.RGB565BE]) {
	for {
		if s.buttonPressed {
			return
		}

		for _, textColor := range s.startTextColors {
			s.animateStartText(text, textColor)
		}
	}
}

func (s *Service) animateStartText(text *tinygl.Text[pixel.RGB565BE], textColor pixel.RGB565BE) {
	const startText = "Press A to start!"

	text.SetColor(textColor)
	s.screen.Update()
	// tinyfont.WriteLine(s.canvas, &freesans.Regular18pt7b, 20, 180, startText, textColor)
	time.Sleep(500 * time.Millisecond)
}

func (s *Service) DrawGameOverMenu() {
	const (
		titleText = "Game Over - Get Go-od!"
	)

	// s.canvas.FillScreen(color.RGBA{255, 255, 255, 255})

	// tinyfont.WriteLine(s.canvas, &freesans.Regular24pt7b, 80, 50, titleText, color.RGBA{0, 0, 0, 0})

	// go s.handleStartText()
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
