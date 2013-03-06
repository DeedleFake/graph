package graph

import (
	"math"
)

type GraphFunc func(x float64) (y float64)

type Graph struct {
	d Output

	Bounds    Rect
	Precision float64
}

func New(d Output) *Graph {
	return &Graph{
		d: d,

		Bounds:    Rt(-5, -5, 10, 10),
		Precision: 1,
	}
}

func (g *Graph) Graph(f GraphFunc) error {
	r := g.Bounds.Canon()
	offX := r.Min.X * float64(g.d.Width()) / r.Dx()
	offY := r.Min.Y * float64(g.d.Height()) / r.Dy()

	last := Point{math.NaN(), math.NaN()}
	for x := r.Min.X; x < r.Max.X+g.Precision; x += g.Precision {
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
