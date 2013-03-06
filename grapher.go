package main

import (
	"./graph"
	"flag"
	"github.com/DeedleFake/sdl"
	"image/color"
	"math"
	"time"
)

func function(x float64) float64 {
	return math.Log(x)
}

func main() {
	var flags struct {
		w, h int
		p    float64
	}
	flag.IntVar(&flags.w, "w", 640, "The width of the screen.")
	flag.IntVar(&flags.h, "h", 480, "The height of the screen.")
	flag.Float64Var(&flags.p, "p", .1, "The precision of the graph.")
	flag.Parse()

	d, err := NewDisplay("Graph", flags.w, flags.h)
	if err != nil {
		panic(err)
	}

	err = d.Clear(color.Black)
	if err != nil {
		panic(err)
	}

	err = d.Color(color.White)
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
