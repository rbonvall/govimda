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

func NewViewport() *Viewport {
	// Maximized for now.

	winWidth, winHeight := termbox.Size()
	return &Viewport{
		X: 0, Y: 0,
		Width: winWidth,
		Height: winHeight - 2,
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
