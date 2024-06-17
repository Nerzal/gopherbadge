package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/conejoninja/gopherbadge/game/entity"
	"github.com/conejoninja/gopherbadge/game/menu"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

// Global Game Infos
var (
	lives              = 3
	score              = 0
	gameState          = StartState
	deltaTime          = float32(0.0)
	lastDeltaTimestamp = time.Now()
	buttonPressed      = false

	player = entity.Entity{
		PosX:   0,
		PosY:   0,
		Width:  10,
		Height: 10,
		// Image:  pixel.NewImage(pixel.RGB565BE, 10, 10),
	}

	menuService     *menu.Service
	backgroundColor = color.RGBA{255, 255, 255, 255}
	white           = color.RGBA{0, 0, 0, 0}

	enemies           = []*entity.EnemyEntity{}
	currentEnemyScore = initialEnemyScore
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
	initialEnemyScore         = 100
)

func main() {
	display, btnA := initialize()
	display.FillScreen(backgroundColor)

	menuService = menu.New(display)
	gameLoop(display, btnA)
}

func gameLoop(display st7789.DeviceOf[pixel.RGB565BE], btnA machine.Pin) {
	for {
		switch gameState {
		case StartState:
			menuService.DrawStartMenu()
			startGame(btnA)
		case InGameState:
			now := time.Now()
			deltaTime = float32(now.Sub(lastDeltaTimestamp).Seconds())

			isGameOver := update(btnA, deltaTime)
			draw(display)

			if isGameOver {
				gameState = GameOverState
			}

			lastDeltaTimestamp = now
		case GameOverState:
			menuService.DrawGameOverMenu()
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

func update(btnA machine.Pin, deltaTime float32) bool {
	// TODO move world unit movement speed based to the left

	cullingOffset := -1

	for idx, enemy := range enemies {
		enemy.Move(deltaTime, MovementSpeed)

		if enemy.HasBeenPassedByPlayer(&player) && !enemy.HasBeenScored {
			updateScore(currentEnemyScore)
			enemy.HasBeenScored = true
		}

		if enemy.ShouldBeCulled() {
			cullingOffset = idx
			continue
		}

		if enemy.DidCollide {
			continue
		}

		if player.HasCollision(enemy.Entity) {
			lives--
			if lives <= 0 {
				return true
			}
		}
	}

	enemies = enemies[cullingOffset+1:]

	return false
}

func updateScore(scoredPoints int) {
	score += scoredPoints
	// TODO implement effects when certain milestones have been passed?
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

func ButtonStateChanged(btnA machine.Pin) {
	buttonPressed = !buttonPressed
	if buttonPressed {
		menuService.OnButtonPressed()
	}
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
	btnA.SetInterrupt(machine.PinToggle, ButtonStateChanged)

	return display, btnA
}
