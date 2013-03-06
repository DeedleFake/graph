package graph

import (
	"math"
)

type Rect struct {
	Min Point
	Max Point
}

func Rt(x, y, w, h float64) Rect {
	return Rect{
		Min: Point{x, y},
		Max: Point{x + w, y + h},
	}
}

func (r Rect) Canon() Rect {
	r.Min.X, r.Max.X = math.Min(r.Min.X, r.Max.X), math.Max(r.Min.X, r.Max.X)
	r.Min.Y, r.Max.Y = math.Min(r.Min.Y, r.Max.Y), math.Max(r.Min.Y, r.Max.Y)

	return r
}

func (r Rect) Dx() float64 {
	return math.Abs(r.Max.X - r.Min.X)
}

func (r Rect) Dy() float64 {
	return math.Abs(r.Max.Y - r.Min.Y)
}

type Point struct {
	X, Y float64
}
