package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/st7789"

	"tinygo.org/x/tinydraw"
)

func main() {

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
	btnB := machine.BUTTON_B
	btnUp := machine.BUTTON_UP
	btnLeft := machine.BUTTON_LEFT
	btnDown := machine.BUTTON_DOWN
	btnRight := machine.BUTTON_RIGHT
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnUp.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnLeft.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnDown.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnRight.Configure(machine.PinConfig{Mode: machine.PinInput})

	display.FillScreen(color.RGBA{255, 255, 255, 255})

	circle := color.RGBA{0, 100, 250, 255}
	white := color.RGBA{255, 255, 255, 255}
	// ring := color.RGBA{200, 0, 0, 255}

	// Clear the display to white
	display.FillScreen(white)

	// Draw blue circles to represent each of the buttons
	tinydraw.FilledCircle(&display, 25, 120, 14, circle) // LEFT

	// display.DrawBitmap()

	for {
		update()
		draw()
	}

}

func draw() {

}

func update() {

}
