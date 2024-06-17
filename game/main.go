package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

// Global Game Infos
var (
	lives              = 3
	score              = 0
	gameState          = StartState
	deltaTime          = 0.0
	lastDeltaTimestamp = time.Now()

	player = Entity{
		PosX:   0,
		PosY:   0,
		Width:  10,
		Height: 10,
		// Image:  pixel.NewImage(pixel.RGB565BE, 10, 10),
	}

	enemies         = []Entity{}
	backgroundColor = color.RGBA{255, 255, 255, 255}
)

// 3 Game states:
const (
	StartState    int = iota
	InGameState   int = iota
	GameOverState int = iota
)

// Foo Vars
const (
	JumpHeight                = 6
	MovementSpeed             = 4
	MinDistanceBetweenEnemies = 12
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
	// ring := color.RGBA{200, 0, 0, 255}

	// Clear the display to white
	display.FillScreen(backgroundColor)

	// Draw blue circles to represent each of the buttons
	// tinydraw.FilledCircle(&display, 25, 120, 14, circle) // LEFT

	// display.DrawBitmap()

	gameLoop(display, btnA)
}

func gameLoop(display st7789.DeviceOf[pixel.RGB565BE], btnA machine.Pin) {

	for {
		switch state {
		case StartState:
			startGame(btnA)
		case InGameState:
			isGameOver := update(btnA)
			draw(display)

			if isGameOver {
				state = GameOverState
				drawGameOverMenu(display)
			}
		case GameOverState:
			restart(btnA)
		}

	}
}

func draw(display st7789.DeviceOf[pixel.RGB565BE]) {
	// display.DrawFastVLine(0, 420, 8, color.RGBA{255, 0, 0, 0})

	// Draw World
	// Draw Gopher
	// Draw "UI"
}

func update(btnA machine.Pin) bool {

	// TODO move world unit movement speed based to the left

	// TODO check collision
	// If collision check Lives
	// if no lives left set game state

	// return true if game is over
	return false
}

func checkCollision() {

}

func startGame(btnA machine.Pin) {
	state = InGameState

}

func restart(btnA machine.Pin) {
	state = StartState
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
