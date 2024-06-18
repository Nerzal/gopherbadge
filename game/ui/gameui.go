package ui

import (
	"fmt"
	"image/color"
	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	scoreFormat       = "score: %d"
	scorePositionX    = 8
	scorePositionY    = 16
	distancePositionX = 8
	distancePositionY = 32
	distanceFormat    = "distance: %.1f"
)

func DrawGameUi(display *st7789.DeviceOf[pixel.RGB565BE], score int, distance float32) {
	tinyfont.WriteLine(display, &freesans.Regular9pt7b, scorePositionX, scorePositionY, fmt.Sprintf(scoreFormat, score), color.RGBA{0, 0, 0, 0})
	tinyfont.WriteLine(display, &freesans.Regular9pt7b, distancePositionX, distancePositionY, fmt.Sprintf(distanceFormat, distance), color.RGBA{0, 0, 0, 0})
}
