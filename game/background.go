package main

import (
	"image/color"
	"math"
	"tinygo.org/x/drivers"
	"tinygo.org/x/tinydraw"
)

const (
	displayWidth            = 320
	displayHeight           = 240
	horizonHeightPercentage = 63 // percentage measured from the top
	horizonPadding          = 16 // measured in pixels from the left and right
	backgroundElementColor  = color.RGBA{0, 0, 0, 255}
)

// TODO implement actual background not just a line...
func drawBackground(display drivers.Displayer) {
	horizonX := int16(math.Floor(displayHeight * horizonHeightPercentage / 100.0))
	tinydraw.Line(display, horizonPadding, horizonX, displayWidth-horizonPadding, horizonX, backgroundElementColor)
}
