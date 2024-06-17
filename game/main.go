package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/conejoninja/gopherbadge/game/entity"
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
	buttonPressed      = false

	player = entity.Entity{
		PosX:   0,
		PosY:   0,
		Width:  10,
		Height: 10,
		// Image:  pixel.NewImage(pixel.RGB565BE, 10, 10),
	}

	backgroundColor = color.RGBA{255, 255, 255, 255}
	white           = color.RGBA{0, 0, 0, 0}

	enemies = []*entity.Entity{}
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
		switch gameState {
		case StartState:
			drawStartMenu(display)
			startGame(btnA)
		case InGameState:
			now := time.Now()
			deltaTime = now.Sub(lastDeltaTimestamp).Seconds()

			isGameOver := update(btnA, deltaTime)
			draw(display)

			if isGameOver {
				gameState = GameOverState
			}

			lastDeltaTimestamp = now
		case GameOverState:
			drawGameOverMenu(display)
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

func update(btnA machine.Pin, deltaTime float64) bool {

	// TODO move world unit movement speed based to the left

	for _, entity := range enemies {
		if !player.HasCollision(entity) {
			continue
		}

		lives--
		if lives <= 0 {
			return false
		}
	}
	// TODO check collision
	// If collision check Lives
	// if no lives left set game state
	// return true if game is over
	return false

}

func startGame(btnA machine.Pin) {
	if buttonPressed {
		gameState = InGameState
	}
	lastDeltaTimestamp = time.Now()
}

func restart(btnA machine.Pin) {
	if buttonPressed {
		gameState = StartState
	}
}

func drawStartMenu(display st7789.DeviceOf[pixel.RGB565BE]) {
	// Draw Title
	//	tinyfont.WriteLine(&display, &freesans.Regular12pt7b, 20, 180, "MOVE the Gopher to see", backgroundColor)
}

func drawGameOverMenu(display st7789.DeviceOf[pixel.RGB565BE]) {

}

func ButtonStateChanged(btnA machine.Pin) {
	buttonPressed = btnA.Get()
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
	btnA.SetInterrupt(1, ButtonStateChanged)

	return display, btnA
}
