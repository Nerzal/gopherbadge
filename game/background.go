package main

import (
	"image/color"

	"github.com/aykevl/tinygl/gfx"
	"tinygo.org/x/drivers/pixel"
)

const (
	displayWidth            = 320
	displayHeight           = 240
	horizonHeightPercentage = 67 // percentage measured from the top
	horizonPadding          = 16 // measured in pixels from the left and right
)

var (
	backgroundElementColor = color.RGBA{0, 0, 0, 255}
)

// TODO implement actual background not just a line...
func drawBackground(canvas *gfx.Canvas[pixel.RGB565BE]) {
	// horizonX := int16(math.Floor(displayHeight * horizonHeightPercentage / 100.0))
	// tinydraw.Line(canvas, horizonPadding, horizonX, displayWidth-horizonPadding, horizonX, backgroundElementColor)
}
