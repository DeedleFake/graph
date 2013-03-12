package main

import (
	"fmt"
	"github.com/DeedleFake/graph"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <out.png>\n", os.Args[0])
		os.Exit(2)
	}

	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	g := graph.NewDot(img)

	fmt.Println("Generating graph...")
	g.Bool(func(x, y float64) color.Color {
		switch {
		case x > 2:
			return color.Black
		case x < -2:
			return color.White
		}

		return nil
	})

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
