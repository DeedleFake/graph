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
	offX := r.Min.X * float64(g.d.Width()) / r.Dx()
	offY := r.Min.Y * float64(g.d.Height()) / r.Dy()

	p := math.Abs(g.Precision)

	last := Point{math.NaN(), math.NaN()}
	for x := r.Min.X; x < r.Max.X+p; x += p {
		y := f(x)

		var to Point
		to.X = (x * float64(g.d.Width()) / r.Dx()) - offX
		to.Y = float64(g.d.Height()) - ((y * float64(g.d.Height()) / r.Dy()) - offY)

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
