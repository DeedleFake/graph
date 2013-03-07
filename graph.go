// Package graph implements a simple mathematical graph generator.
//
// The main type is Graph, which is capable of rendering to an Output. For convience, a type called ImageOutput is provided which wraps several of draw.Image's methods, allowing one to be used as an Output.
//
// For example:
//
//	img := image.NewRGBA(image.Rect(0, 0, 320, 240))
//	imgout := graph.ImageOutput{img, color.Black}
//
//	g := NewGraph(imgout)
//	g.Cart(math.Sin)
package graph

import (
	"math"
)

// CartFunc represents a Cartesian function.
type CartFunc func(x float64) (y float64)

// A Graph calculates graphs from functions and renders to an Output.
type Graph struct {
	d Output

	// Bounds is the bounding area for the graph. On graphing
	// calculators, it's something like XMin, XMax, YMin, and YMax.
	Bounds Rect

	// Precision is how small each step of the graph should be. The
	// closer to zero it is, the better the resulting graph will look.
	Precision float64
}

// New returns a new Graph that renders to the given Output. The
// Bounds and Precision of the returned Graph default to
// Rt(-5, -5, 10, 10) and .1, respectively.
func New(d Output) *Graph {
	return &Graph{
		d: d,

		Bounds:    Rt(-5, -5, 10, 10),
		Precision: .1,
	}
}

// Cart graphs the given CartFunc, rendering to the associated Output.
// It returns an error, if any.
func (g *Graph) Cart(f CartFunc) error {
	r := g.Bounds.Canon()
	ob := g.d.Bounds().Canon()

	// TODO: Implement a method for point that calculates the offset.
	off := Point{
		X: (r.Min.X * float64(ob.Dx()) / r.Dx()) - float64(ob.Min.X),
		Y: (r.Min.Y * float64(ob.Dy()) / r.Dy()) - float64(ob.Min.Y),
	}

	p := math.Abs(g.Precision)

	last := Point{math.NaN(), math.NaN()}
	for x := r.Min.X; x < r.Max.X+p; x += p {
		y := f(x)

		to := Point{x, y}.GraphToOutputNoOffset(r, ob)
		to.X -= off.X
		to.Y = float64(ob.Dy()) - (to.Y - off.Y)

		if !(math.IsNaN(last.Y) || math.IsInf(last.Y, 0) || math.IsNaN(to.Y) || math.IsInf(to.Y, 0)) {
			err := g.d.Line(last.ImagePoint(), to.ImagePoint())
			if err != nil {
				return err
			}
		}

		last = to
	}

	return nil
}
