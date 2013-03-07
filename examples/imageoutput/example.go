package main

import (
	"fmt"
	"github.com/DeedleFake/graph"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <out.png>\n", os.Args[0])
		os.Exit(2)
	}

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	imgout := graph.ImageOutput{img, color.Black}

	g := graph.New(imgout)
	g.Precision = .001

	fmt.Println("Generating graph...")
	err := g.Cart(math.Sin)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("Saving to %v...\n", os.Args[1])
	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
