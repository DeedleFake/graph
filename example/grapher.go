package main

import (
	".."
	"flag"
	"github.com/DeedleFake/sdl"
	"image"
	"image/color"
	"math"
	"time"
)

func next(a float64) func(float64) float64 {
	return func(x float64) float64 {
		return math.Sin(x * a)
	}
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

		animFrom, animTo float64
		animSpeed        float64
	}
	flag.IntVar(&flags.w, "w", 640, "The width of the screen.")
	flag.IntVar(&flags.h, "h", 480, "The height of the screen.")
	flag.Float64Var(&flags.p, "p", .1, "The precision of the graph.")
	flag.BoolVar(&flags.axes, "axes", true, "Draw axes.")
	flag.Float64Var(&flags.animFrom, "anim.from", -1, "What number to animate from.")
	flag.Float64Var(&flags.animTo, "anim.to", 1, "What number to animate to.")
	flag.Float64Var(&flags.animSpeed, "anim.speed", .01, "Speed of animation.")
	flag.Parse()

	d, err := NewDisplay("Graph Example", flags.w, flags.h)
	if err != nil {
		panic(err)
	}

	g := graph.New(d)
	if err != nil {
		panic(err)
	}

	g.Precision = flags.p

	a := flags.animFrom

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

		a += flags.animSpeed
		if (a >= flags.animTo) || (a <= flags.animFrom) {
			flags.animSpeed *= -1
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

		c := color.RGBA{255, 0, 255, 255}
		switch {
		case flags.animSpeed < 0:
			c = color.RGBA{0, 0, 255, 255}
		case flags.animSpeed > 0:
			c = color.RGBA{255, 0, 0, 255}
		}

		err = d.Color(c)
		if err != nil {
			panic(err)
		}

		err = g.Cart(next(a))
		if err != nil {
			panic(err)
		}

		d.Flip()

		if fps != nil {
			<-fps.C
		}
	}
}
