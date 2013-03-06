package main

import (
	"./graph"
	"github.com/DeedleFake/sdl"
	"image/color"
	"math"
	"time"
)

func main() {
	d, err := NewDisplay("Graph", 640, 480)
	if err != nil {
		panic(err)
	}

	err = d.Color(color.Black)
	if err != nil {
		panic(err)
	}

	err = d.Clear()
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

	g.Precision = .01

	err = g.Graph(math.Sin)
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
