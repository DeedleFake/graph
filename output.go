package graph

import (
	"image"
)

// An Output is a place for a Graph to render to.
type Output interface {
	// Line draws a line from from to to. Any error returned by Line
	// will be returned by the calling Graph method.
	Line(from, to image.Point) error

	// Bounds returns the bounds of the output. It is analogous to
	// image.Bounds().
	Bounds() image.Rectangle
}
