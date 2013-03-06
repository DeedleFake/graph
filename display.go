package main

import (
	"github.com/DeedleFake/sdl"
	"image"
	"image/color"
	"runtime"
	"sync"
)

type Display struct {
	win    *sdl.Window
	screen *sdl.Surface
}

func NewDisplay(title string, w, h int) (*Display, error) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, err
	}

	win, err := sdl.CreateWindow(title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		w,
		h,
		0,
	)
	if err != nil {
		return nil, err
	}

	screen, err := win.GetSurface()
	if err != nil {
		return nil, err
	}

	d := &Display{
		win:    win,
		screen: screen,
	}
	runtime.SetFinalizer(d, (*Display).Close)

	return d, nil
}

func (d *Display) Close() error {
	d.screen.Free()
	d.win.Destroy()

	sdl.Quit()

	return nil
}

func (d *Display) Line(from, to image.Point) error {
	if to.X < from.X {
		from, to = to, from
	}

	s := float64(to.Y-from.Y) / float64(to.X-from.X)

	var wg sync.WaitGroup
	for x := from.X; x < to.X; x++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()

			y := int(s*float64(x)) + from.X
			d.Set(x, y, color.White)
		}(x)
	}
	wg.Wait()

	err := d.UpdateRect(image.Rectangle{from, to})
	if err != nil {
		return err
	}

	return nil
}

func (d *Display) Set(x, y int, c color.Color) {
	r, g, b, a := c.RGBA()
	sr := uint8(r * 255 / 0xFFFF)
	sg := uint8(g * 255 / 0xFFFF)
	sb := uint8(b * 255 / 0xFFFF)
	sa := uint8(a * 255 / 0xFFFF)

	d.screen.FillRect(&sdl.Rect{X: int32(x), Y: int32(y), W: 1, H: 1}, d.screen.Format.MapRGBA(sr, sg, sb, sa))
}

func (d *Display) UpdateRect(r image.Rectangle) error {
	sr := rectToSDL(r)

	return d.win.UpdateSurfaceRects([]sdl.Rect{*sr})
}

func rectToSDL(r image.Rectangle) *sdl.Rect {
	r = r.Canon()

	return &sdl.Rect{
		X: int32(r.Min.X),
		Y: int32(r.Min.Y),
		W: int32(r.Dx()),
		H: int32(r.Dy()),
	}
}
