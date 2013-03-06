package graph

import (
	"image"
)

type Output interface {
	Line(from, to image.Point) error

	Width() int
	Height() int
}
