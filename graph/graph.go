package graph

import (
	"image"
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
	offX := int(r.Min.X * float64(g.d.Width()) / r.Dx())
	offY := int(r.Min.Y * float64(g.d.Height()) / r.Dy())

	last := image.Pt(-1, -1)
	for x := r.Min.X; x < r.Max.X+g.Precision; x += g.Precision {
		y := f(x)

		sx := int(x*float64(g.d.Width())/r.Dx()) - offX
		sy := g.d.Height() - (int(y*float64(g.d.Height())/r.Dy()) - offY)

		if last.X >= 0 {
			err := g.d.Line(last, image.Pt(sx, sy))
			if err != nil {
				return err
			}
		}

		last = image.Pt(sx, sy)
	}

	return nil
}
