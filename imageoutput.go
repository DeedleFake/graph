package graph

import (
	"image"
	"image/color"
	"math"
)

// Image is a minimal version of draw.Image that only contains
// required methods.
type Image interface {
	Bounds() image.Rectangle
	Set(x, y int, c color.Color)
}

// An ImageOutput is an implementation of Output that wraps an Image.
type ImageOutput struct {
	Image

	// C is the color to use for drawing lines.
	C color.Color
}

func (io ImageOutput) Line(from, to image.Point) error {
	x1, y1, x2, y2 := float64(from.X), float64(from.Y), float64(to.X), float64(to.Y)

	steep := math.Abs(y2-y1) > math.Abs(x2-x1)
	if steep {
		x1, y1 = y1, x1
		x2, y2 = y2, x2
	}

	if x1 > x2 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}

	dx := x2 - x1
	dy := math.Abs(y2 - y1)

	e := dx / 2
	ys := -1
	if y1 < y2 {
		ys = 1
	}

	y := int(y1)
	for x := int(x1); x < int(x2); x++ {
		if steep {
			io.Set(y, x, io.C)
		} else {
			io.Set(x, y, io.C)
		}

		e -= dy
		if e < 0 {
			y += int(ys)
			e += dx
		}
	}

	return nil
}
