package screen

import (
	"buffer"
	"github.com/nsf/termbox-go"
	"strings"
)

type Viewport struct {
	X, Y, Width, Height int
	Buffer *buffer.T
	I, J int
	Fg, Bg termbox.Attribute
}

func NewViewport(x, y, w, h int) *Viewport {
	return &Viewport{
		X: x, Y: y,
		Width: w,
		Height: h,
		Fg: termbox.ColorDefault,
		Bg: termbox.ColorDefault,
	}
}

func (v *Viewport) printStringAt(y, x0 int, s string) {
	for i, c := range s {
		x := x0 + i
		if x > v.X + v.Width {
			return
		}
		termbox.SetCell(x, y, rune(c), v.Fg, v.Bg)
	}
}

func fitStringToWidth(s string, w int) string {
	n := len(s)
	if n < w {
		return s + strings.Repeat(" ", w - n)
	}
	return s[:w]
}

func (v *Viewport) Draw() {
	if v.Buffer == nil {
		return
	}
	i := v.I
	for y := v.Y; y < v.Height; y++ {
		line := v.Buffer.LineByIndex(i)
		if line != nil {
			l := strings.Replace(line.Data, "\t", "    ", -1)
			l = fitStringToWidth(l, v.Width)
			v.printStringAt(y, v.X, l)
		}
		i++
	}
	termbox.Flush()
}

func (v *Viewport) Paint(color termbox.Attribute) {
	cb := termbox.CellBuffer()
	screenWidth, _ := termbox.Size()
	for y := v.Y; y < v.Y + v.Height; y++ {
	for x := v.X; x < v.X + v.Width;  x++ {
			cb[screenWidth * y + x].Bg = color
			// time.Sleep(5 * time.Millisecond)
			// termbox.Flush()
		}
	}
	termbox.Flush()
}
