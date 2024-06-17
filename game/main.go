package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

// Global Game Infos
var (
	lives     = 3
	score     = 0
	gameState = StartState

	player = Entity{
		PosX:   0,
		PosY:   0,
		Width:  10,
		Height: 10,
		// Image:  pixel.NewImage(pixel.RGB565BE, 10, 10),
	}

	enemies = []Entity{}
)

// 3 Game states:
const (
	StartState    int = iota
	InGameState   int = iota
	GameOverState int = iota
)

type Entity struct {
	PosX   int16
	PosY   int16
	Width  int16
	Height int16
	Image  pixel.Image[pixel.RGB565BE]
}

func main() {
	display, btnA := initialize()

	// circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	// ring := color.RGBA{200, 0, 0, 255}

	// Clear the display to white
	display.FillScreen(white)

	// Draw blue circles to represent each of the buttons
	// tinydraw.FilledCircle(&display, 25, 120, 14, circle) // LEFT

	// display.DrawBitmap()

	gameLoop(display, btnA)
}

func gameLoop(display st7789.DeviceOf[pixel.RGB565BE], btnA machine.Pin) {
	for {
		update(btnA)
		draw(display)
	}
}

func draw(display st7789.DeviceOf[pixel.RGB565BE]) {
	// display.DrawFastVLine(0, 420, 8, color.RGBA{255, 0, 0, 0})

	// Draw World
	// Draw Gopher
	// Draw "UI"
}

func update(btnA machine.Pin) {

	// TODO move world 1 unit to the left
	// TODO check collision
	// If collision check Lives
	// if no lives left set game state

}

func checkCollision() {

}

func initialize() (st7789.DeviceOf[pixel.RGB565BE], machine.Pin) {
	// Setup the screen pins
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	// get and configure buttons on the board
	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})

	return display, btnA
}
