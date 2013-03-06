package main

import (
	"./graph"
	"flag"
	"github.com/DeedleFake/sdl"
	"image"
	"image/color"
	"math"
	"time"
)

func function(x float64) float64 {
	return math.Log(x)
}

func drawAxes(d *Display) error {
	err := d.Line(image.Pt(d.Width()/2, 0), image.Pt(d.Width()/2, d.Height()))
	if err != nil {
		return err
	}

	err = d.Line(image.Pt(0, d.Height()/2), image.Pt(d.Width(), d.Height()/2))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var flags struct {
		w, h int
		p    float64
		axes bool
	}
	flag.IntVar(&flags.w, "w", 640, "The width of the screen.")
	flag.IntVar(&flags.h, "h", 480, "The height of the screen.")
	flag.Float64Var(&flags.p, "p", .1, "The precision of the graph.")
	flag.BoolVar(&flags.axes, "axes", true, "Draw axes.")
	flag.Parse()

	d, err := NewDisplay("Graph", flags.w, flags.h)
	if err != nil {
		panic(err)
	}

	err = d.Clear(color.Black)
	if err != nil {
		panic(err)
	}

	if flags.axes {
		err = d.Color(color.White)
		if err != nil {
			panic(err)
		}

		err = drawAxes(d)
		if err != nil {
			panic(err)
		}
	}

	err = d.Color(color.RGBA{255, 0, 0, 0})
	if err != nil {
		panic(err)
	}

	g := graph.New(d)
	if err != nil {
		panic(err)
	}

	g.Precision = flags.p

	err = g.Graph(function)
	if err != nil {
		panic(err)
	}
	d.Flip()

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
