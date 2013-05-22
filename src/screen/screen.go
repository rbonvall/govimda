package screen

import (
	"buffer"
	"github.com/nsf/termbox-go"
)

type Viewport struct {
	X, Y, Width, Height int
	Buffer *buffer.Type
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
