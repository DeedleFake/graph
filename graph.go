package main

import (
	"github.com/DeedleFake/sdl"
	"image"
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

	g := &Graph{
		d: d,

		Bounds:    Rect{0, 0, 10, 10},
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
	g.d.Line(image.Pt(10, 10), image.Pt(100, 300))

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

func main() {
	g, err := NewGraph()
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Pause()
}
