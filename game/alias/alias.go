package alias

import (
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/gfx"
	"tinygo.org/x/drivers/pixel"
)

type Canvas struct {
	*gfx.Canvas[pixel.RGB565BE]
}

type Screen struct {
	*tinygl.Screen[pixel.RGB565BE]
}

type Rect struct {
	*gfx.Rect[pixel.RGB565BE]
}

type Text struct {
	*tinygl.Text[pixel.RGB565BE]
}

type Color struct {
	pixel.RGB565BE
}

type Image struct {
	*gfx.Image[pixel.RGB565BE]
}
