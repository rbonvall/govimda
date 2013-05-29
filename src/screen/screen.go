package screen

import (
	"github.com/nsf/termbox-go"
	"buffer"
)

type T struct {
	EditArea *Viewport
	Panel    *Viewport
	Gutter   *Viewport
	CmdLine  *Viewport

	BufferList *buffer.List
}

func DrawVerticalLine(x, y0, y1 int) {
	fg := termbox.ColorDefault
	bg := termbox.ColorDefault

	for y := y0; y <= y1; y++ {
		termbox.SetCell(x, y, '│', fg, bg)
	}
}

func DrawHorizontalLine(x0, x1, y int) {
	fg := termbox.ColorDefault
	bg := termbox.ColorDefault

	for x := x0; x <= x1; x++ {
		termbox.SetCell(x, y, '─', fg, bg)
	}
}

func PutCell(x, y int, c rune) {
	fg := termbox.ColorDefault
	bg := termbox.ColorDefault
	termbox.SetCell(x, y, c, fg, bg)
}

// Termbox wrappers.

func Init() {
	termbox.Init()
}

func Close() {
	termbox.Close()
}

func Size() (int, int) {
	return termbox.Size()
}

func Refresh() {
	termbox.Flush()
}

