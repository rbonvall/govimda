package screen

import (
	"buffer"
	"github.com/nsf/termbox-go"
)

type Viewport struct {
	X, Y, Width, Height int
	Buffer *buffer.T
	I, J int
}

func NewViewport(x, y, w, h int) *Viewport {
	return &Viewport{
		X: x, Y: y,
		Width: w,
		Height: h,
	}

}

func (v *Viewport) printStringAt(y, x0 int, s string) {
	fg := termbox.ColorDefault
	bg := termbox.ColorDefault
	for i, c := range s {
		x := x0 + i
		if x > v.X + v.Width {
			return
		}
		termbox.SetCell(x, y, rune(c), fg, bg)
	}
}

func (v *Viewport) Draw() {
	i := v.I
	for y := v.Y; y < v.Height; y++ {
		line := v.Buffer.LineByIndex(i)
		if line != nil {
			v.printStringAt(y, v.X, line.Data)
		}
		i++
	}
	termbox.Flush()
}

func (v *Viewport) Paint(color termbox.Attribute) {
	for y := v.Y; y < v.Height; y++ {
		for x := v.X; x < v.Width; x++ {
			termbox.CellBuffer()[y * v.Width + x].Bg = color

		}
	}
}
