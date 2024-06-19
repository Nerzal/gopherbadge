package main

import (
	"github.com/aykevl/tinygl/gfx"
	"github.com/conejoninja/gopherbadge/game/alias"
	"tinygo.org/x/drivers/pixel"
)

const (
	displayWidth            = 320
	displayHeight           = 240
	horizonHeightPercentage = 67 // percentage measured from the top
	horizonPadding          = 16 // measured in pixels from the left and right
)

// TODO implement actual background not just a line...
func drawBackground(canvas alias.Canvas) {
	line := gfx.NewLine(pixel.NewRGB565BE(0, 0, 0), horizonPadding, 180, displayWidth-horizonPadding, 180, 2)
	canvas.Add(line)
}
