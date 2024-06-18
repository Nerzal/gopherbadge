package ui

import (
	"github.com/conejoninja/gopherbadge/game/alias"
)

const (
	scoreFormat       = "score: %d"
	scorePositionX    = 8
	scorePositionY    = 16
	distancePositionX = 8
	distancePositionY = 32
	distanceFormat    = "distance: %.1f"
)

func DrawGameUi(canvas alias.Canvas, score int, distance float32) {
	// tinyfont.WriteLine(canvas, &freesans.Regular9pt7b, scorePositionX, scorePositionY, fmt.Sprintf(scoreFormat, score), color.RGBA{0, 0, 0, 0})
	// tinyfont.WriteLine(canvas, &freesans.Regular9pt7b, distancePositionX, distancePositionY, fmt.Sprintf(distanceFormat, distance), color.RGBA{0, 0, 0, 0})
}
