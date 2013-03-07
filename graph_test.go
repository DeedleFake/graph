package graph_test

import (
	"github.com/DeedleFake/graph"
	"image"
	"image/color"
	"math"
)

func ExampleNew() {
	img := image.NewRGBA(image.Rect(0, 0, 320, 240))
	imgout := graph.ImageOutput{img, color.Black}

	g := graph.New(imgout)
	err := g.Cart(func(x float64) float64 {
		if x < 0 {
			return math.Sin(x)
		}

		return math.Cos(x)
	})
	if err != nil {
		println(err.Error())
	}
}
