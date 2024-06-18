package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/conejoninja/gopherbadge/game/alias"
	"github.com/conejoninja/gopherbadge/game/entity"
	"github.com/conejoninja/gopherbadge/game/menu"
	"tinygo.org/x/drivers/pixel"

	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/gfx"
)

// Global Game Infos
var (
	score              = 0
	highScore          = 0
	lives              = 3
	gameState          = StartState
	deltaTime          = float32(0.0)
	lastDeltaTimestamp = time.Now()
	buttonPressed      = false

	player = entity.NewPlayer()

	menuService     *menu.Service
	backgroundColor = color.RGBA{255, 255, 255, 255}
	// white           = color.RGBA{0, 0, 0, 0}

	speaker machine.Pin
	bzrPin  machine.Pin

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
	canvas, screen, btnA := initialize[pixel.RGB565BE]()

	menuService = menu.New(canvas, screen)
	gameLoop(canvas, screen, btnA)
}

func gameLoop(canvas alias.Canvas, screen alias.Screen, btnA machine.Pin) {
	for {
		switch gameState {
		case StartState:
			menuService.DrawStartMenu()
			startGame()
			// canvas.FillScreen(color.RGBA{255, 255, 255, 255})
		case InGameState:
			now := time.Now()
			deltaTime = float32(now.Sub(lastDeltaTimestamp).Seconds())
			// canvas.FillScreen(color.RGBA{255, 255, 255, 255})

			isGameOver := update(btnA, deltaTime)
			screen.Update()

			if isGameOver {
				gameState = GameOverState
			}

			lastDeltaTimestamp = now
		case GameOverState:
			if score > highScore {
				highScore = score
			}

			menuService.DrawGameOverMenu()
			restart(btnA)
		}

	}
}

func update(btnA machine.Pin, deltaTime float32) bool {
	// TODO move world unit movement speed based to the left

	if buttonPressed {
		player.Jump()
	}

	player.Move(deltaTime)
	println("PlayerX: ", player.PosX, "PlayerY:", player.PosY)

	cullingOffset := -1

	for idx, enemy := range enemies {
		enemy.Move(deltaTime, MovementSpeed)

		if !enemy.HasBeenScored && enemy.HasBeenPassedByPlayer(player.Entity) {
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
			score--
			if score <= 0 {
				return true
			}
		}
	}

	enemies = enemies[cullingOffset+1:]

	return false
}

func updateScore(scoredPoints int) {
	lives += scoredPoints
	// TODO implement effects when certain milestones have been passed?
}

func startGame() {
	for {
		if buttonPressed {
			gameState = InGameState
			break
		}

		time.Sleep(50 * time.Millisecond)
	}

	lives = 3
	score = 0

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

func initialize[T pixel.Color]() (alias.Canvas, alias.Screen, machine.Pin) {
	// Setup the screen pins
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	board.Buttons.Configure()
	display := board.Display.Configure()
	board.Display.SetBrightness(board.Display.MaxBrightness())
	canvas, screen := initUi(display)

	// get and configure buttons on the board
	btnA := machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnA.SetInterrupt(machine.PinToggle, ButtonStateChanged)

	bzrPin = machine.SPEAKER
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	speaker = machine.SPEAKER_ENABLE
	speaker.Configure(machine.PinConfig{Mode: machine.PinOutput})
	speaker.High()

	tone(1045)

	tone(800)

	return alias.Canvas{Canvas: canvas}, alias.Screen{Screen: screen}, btnA
}

func tone(tone int) {
	for i := 0; i < 30; i++ {
		bzrPin.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)

		bzrPin.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}

func initUi[T pixel.Color](display board.Displayer[T]) (*gfx.Canvas[T], *tinygl.Screen[T]) {
	buf := pixel.NewImage[T](int(240), int(320)/10)
	screen := tinygl.NewScreen[T](display, buf, board.Display.PPI())

	var (
		// black = pixel.NewColor[T](0x00, 0x00, 0x00)
		white = pixel.NewColor[T](0xff, 0xff, 0xff)
	)

	canvas := gfx.NewCanvas(white, 96, 96)
	canvas.SetGrowable(0, 1)

	screen.SetChild(canvas)

	screen.Update()

	return canvas, screen

}
