package graph

import (
	"image"
	"math"
)

// A Rect represents a rectangle. It is patterned after
// image.Rectangle, but uses float64s instead of ints.
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

// A Point represents a point in Cartesian space. It is patterned
// after image.Point, but uses float64s instead of ints.
type Point struct {
	X, Y float64
}

// IsValid returns true if neither p.X nor p.Y are infinity or NaN.
func (p Point) IsValid() bool {
	return !(math.IsNaN(p.X) || math.IsNaN(p.Y) || math.IsInf(p.X, 0) || math.IsInf(p.Y, 0))
}

// ImagePoint returns the given point converted to an image.Point,
// with possible loss of precision.
func (p Point) ImagePoint() image.Point {
	return image.Pt(int(p.X), int(p.Y))
}

// GraphToOutput converts Graph coordinates to Output coordinates.
func (p Point) GraphToOutput(gb Rect, ob image.Rectangle) Point {
	gb = gb.Canon()
	ob = ob.Canon()

	rDx := float64(ob.Dx()) / gb.Dx()
	rDy := float64(ob.Dy()) / gb.Dy()

	off := Point{
		X: gb.Min.X * rDx,
		Y: gb.Min.Y * rDy,
	}

	return Point{
		X: (p.X * rDx) - off.X,
		Y: (p.Y * rDy) - off.Y,
	}
}

// GraphToOutputNoOffset converts Graph coordinates to Output
// coordinates, but doesn't apply an offset. This is only useful for
// special cases.
func (p Point) GraphToOutputNoOffset(gb Rect, ob image.Rectangle) Point {
	gb = gb.Canon()
	ob = ob.Canon()

	rDx := float64(ob.Dx()) / gb.Dx()
	rDy := float64(ob.Dy()) / gb.Dy()

	return Point{
		X: p.X * rDx,
		Y: float64(ob.Dy()) - (p.Y * rDy),
	}
}

// A Vector represents a set of polar coordinates.
type Vector struct {
	Mag, Dir float64
}

// ToPoint converts the Vector to a Point.
func (v Vector) ToPoint() Point {
	if v.Mag == 0 {
		return Point{0, 0}
	}

	y, x := math.Sincos(v.Dir)

	return Point{
		X: v.Mag * x,
		Y: v.Mag * y,
	}
}
