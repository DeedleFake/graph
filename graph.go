package main

import (
	"github.com/DeedleFake/sdl"
	"image"
	"image/color"
	"math"
	"runtime"
	"time"
)

type GraphFunc func(x float64) (y float64)

type Graph struct {
	d *Display

	Bounds    Rect
	Precision float64
}

func NewGraph() (*Graph, error) {
	d, err := NewDisplay("Graph", 640, 480)
	if err != nil {
		return nil, err
	}

	err = d.Color(color.Black)
	if err != nil {
		return nil, err
	}

	err = d.Clear()
	if err != nil {
		return nil, err
	}

	err = d.Color(color.White)
	if err != nil {
		return nil, err
	}

	g := &Graph{
		d: d,

		Bounds:    Rt(-5, -5, 10, 10),
		Precision: 1,
	}
	runtime.SetFinalizer(g, (*Graph).Close)

	return g, nil
}

func (g *Graph) Close() error {
	err := g.d.Close()
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) Pause() {
	fps := time.NewTicker(time.Second / 60)
	for fps != nil {
		var ev sdl.Event
		for sdl.PollEvent(&ev) {
			switch ev.(type) {
			case *sdl.QuitEvent:
				fps.Stop()
				fps = nil
			}
		}

		if fps != nil {
			<-fps.C
		}
	}
}

func (g *Graph) Graph(f GraphFunc) error {
	r := g.Bounds.Canon()
	offX := int(r.Min.X * 640 / r.Dx())
	offY := int(r.Min.Y * 480 / r.Dy())

	last := image.Pt(-1, -1)
	for x := r.Min.X; x < r.Max.X+g.Precision; x += g.Precision {
		y := f(x)

		sx := int(x*640/r.Dx()) - offX
		sy := 480 - (int(y*480/r.Dy()) - offY)

		if last.X >= 0 {
			err := g.d.Line(last, image.Pt(sx, sy))
			if err != nil {
				return err
			}
		}

		last = image.Pt(sx, sy)
	}

	g.d.Flip()

	return nil
}

func main() {
	g, err := NewGraph()
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Precision = .01

	err = g.Graph(math.Sin)
	if err != nil {
		panic(err)
	}

	g.Pause()
}
