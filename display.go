package main

import (
	"github.com/DeedleFake/sdl"
	"image"
	"image/color"
	"runtime"
)

type Display struct {
	win *sdl.Window
	ren *sdl.Renderer
}

func NewDisplay(title string, w, h int) (*Display, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	win, ren, err := sdl.CreateWindowAndRenderer(
		w,
		h,
		0,
	)
	if err != nil {
		return nil, err
	}

	win.SetTitle(title)

	d := &Display{
		win: win,
		ren: ren,
	}
	runtime.SetFinalizer(d, (*Display).Close)

	return d, nil
}

func (d *Display) Close() error {
	d.ren.Destroy()
	d.win.Destroy()

	sdl.Quit()

	return nil
}

//func (d *Display) Line(from, to image.Point) error {
//	x1, y1, x2, y2 := float64(from.X), float64(from.Y), float64(to.X), float64(to.Y)
//
//	steep := math.Abs(y2-y1) > math.Abs(x2-x1)
//	if steep {
//		x1, y1 = y1, x1
//		x2, y2 = y2, x2
//	}
//
//	if x1 > x2 {
//		x1, x2 = x2, x1
//		y1, y2 = y2, y1
//	}
//
//	dx := x2 - x1
//	dy := math.Abs(y2 - y1)
//
//	e := dx / 2
//	ys := -1
//	if y1 < y2 {
//		ys = 1
//	}
//
//	y := int(y1)
//	for x := int(x1); x < int(x2); x++ {
//		if steep {
//			d.Set(y, x, color.White)
//		} else {
//			d.Set(x, y, color.White)
//		}
//
//		e -= dy
//		if e < 0 {
//			y += int(ys)
//			e += dx
//		}
//	}
//
//	err := d.UpdateRect(image.Rectangle{from, to})
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (d *Display) Color(c color.Color) error {
	return d.ren.SetDrawColor(RGBA8(c))
}

func (d *Display) Line(from, to image.Point) error {
	return d.ren.DrawLine(from.X, from.Y, to.X, to.Y)
}

func (d *Display) Clear() error {
	return d.ren.Clear()
}

func (d *Display) Flip() {
	d.ren.Present()
}

func RGBA8(c color.Color) (r8, g8, b8, a8 uint8) {
	r, g, b, a := c.RGBA()

	r8 = uint8(r * 255 / 0xFFFF)
	g8 = uint8(g * 255 / 0xFFFF)
	b8 = uint8(b * 255 / 0xFFFF)
	a8 = uint8(a * 255 / 0xFFFF)

	return
}
