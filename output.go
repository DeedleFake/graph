package graph

import (
	"image"
)

// An Output is a place for a Graph to render to.
type Output interface {
	// Line draws a line from from to to. Any error returned by Line
	// will be returned by the calling Graph method.
	Line(from, to image.Point) error

	// Width returns the width of the output.
	Width() int

	// Height returns the height of the output.
	Height() int
}
