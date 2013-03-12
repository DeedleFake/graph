package graph

import (
	"image"
	"image/color"
)

// A DotOutput is an output that can have individual pixels set. It is
// designed to be compatible with draw.Image.
type DotOutput interface {
	Bounds() image.Rectangle
	Set(x, y int, c color.Color)
}
