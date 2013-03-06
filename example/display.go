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

	c color.Color
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

		c: color.Black,
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

func (d *Display) Color(c color.Color) error {
	old := d.c

	err := d.ren.SetDrawColor(RGBA8(c))
	if err != nil {
		d.c = old
		return err
	}

	d.c = c

	return nil
}

func (d *Display) Line(from, to image.Point) error {
	return d.ren.DrawLine(from.X, from.Y, to.X, to.Y)
}

func (d *Display) Clear(c color.Color) error {
	defer d.Color(d.c)

	err := d.Color(c)
	if err != nil {
		return err
	}

	return d.ren.Clear()
}

func (d *Display) Flip() {
	d.ren.Present()
}

func (d *Display) Width() int {
	w, _ := d.win.GetSize()

	return w
}

func (d *Display) Height() int {
	_, h := d.win.GetSize()

	return h
}

func RGBA8(c color.Color) (r8, g8, b8, a8 uint8) {
	r, g, b, a := c.RGBA()

	r8 = uint8(r * 255 / 0xFFFF)
	g8 = uint8(g * 255 / 0xFFFF)
	b8 = uint8(b * 255 / 0xFFFF)
	a8 = uint8(a * 255 / 0xFFFF)

	return
}
