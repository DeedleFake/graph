package graph

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

// A BoolFunc returns the color that should be drawn for the given
// point on the graph.
type BoolFunc func(x, y float64) color.Color

// A DotGraph draws graphs that require the ability to set dots
// instead of lines.
type DotGraph struct {
	d DotOutput

	Bounds    Rect
	Precision float64
}

// NewDot returns a new DotGraph that draws to the given DotOutput.
func NewDot(d DotOutput) *DotGraph {
	return &DotGraph{
		d: d,

		Bounds:    Rt(-5, -5, 10, 10),
		Precision: 1,
	}
}

func (g *DotGraph) set(x, y float64, gb Rect, ob image.Rectangle, off Point, prec float64, c color.Color) {
	from := Point{x, y}.GraphToOutputNoOffset(gb, ob)
	from.X -= off.X
	from.Y = float64(ob.Dy()) - ((float64(ob.Dy()) - from.Y) - off.Y)

	to := Point{prec, prec}.GraphToOutputNoOffset(gb, ob)
	to.X = (to.X - off.X) + from.X
	to.Y = (float64(ob.Dy()) - ((float64(ob.Dy()) - to.Y) - off.Y)) + from.Y

	for y := int(from.Y); y < int(to.Y); y++ {
		for x := int(from.X); x < int(to.X); x++ {
			g.d.Set(x, y, c)
		}
	}
}

// Bool draws a graph to the associated output using the given
// BoolFunc.
//
// BUG: This doesn't work. It draws weird, random squares.
func (g *DotGraph) Bool(f BoolFunc) {
	gb := g.Bounds.Canon()
	ob := g.d.Bounds().Canon()

	off := Point{
		X: (gb.Min.X * float64(ob.Dx()) / gb.Dx()) - float64(ob.Min.X),
		Y: (gb.Min.Y * float64(ob.Dy()) / gb.Dy()) - float64(ob.Min.Y),
	}

	p := math.Abs(g.Precision)
	if p == 0 {
		p = 1
	}

	var wg sync.WaitGroup
	for y := gb.Min.Y; y < gb.Max.Y; y += p {
		for x := gb.Min.X; x < gb.Max.X; x += p {
			if c := f(x, y); c != nil {
				for runtime.NumGoroutine() > 1000 {
					runtime.Gosched()
				}

				wg.Add(1)
				go func(x, y float64, c color.Color) {
					defer wg.Done()

					g.set(x, y, gb, ob, off, p, c)
				}(x, y, c)
			}
		}
	}
}
