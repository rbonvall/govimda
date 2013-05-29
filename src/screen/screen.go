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

func New() *T {
	W, H := termbox.Size()
	gw := 4
	pw := 20
	ew := W - (gw + pw) - 2
	s := &T{
		Gutter:   NewViewport(0,           0, gw, H - 2),
		EditArea: NewViewport(gw + 1,      0, ew, H - 2),
		Panel:    NewViewport(gw + ew + 2, 0, pw, H - 2),
		CmdLine:  NewViewport(0, H  -1, W, 1),
	}
	s.Gutter.Fg = termbox.ColorYellow

	// initialize buffers
	s.BufferList = buffer.NewListFromArgs()
	s.EditArea.Buffer = s.BufferList.Current()
	s.Gutter.Buffer, _ = buffer.NewFromStrings([]string{"1", "2", "3", "4"})
	// s.Panel.Buffer, _ = buffer.NewFromStrings([]string{"a", "b", "c", "d", "e"})
	s.Panel.Buffer = buffer.NewEmpty()

	for x := 0; x < W; x++ {
		termbox.SetCell(x, H-2, '─', termbox.ColorDefault, termbox.ColorDefault)
	}
	for y := 0; y < H-2; y++ {
		termbox.SetCell(gw, y, '│', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(gw+ew+1, y, '│', termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(gw, H-2, '┴', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(gw+ew+1, H-2, '┴', termbox.ColorDefault, termbox.ColorDefault)

	return s
}

func (s *T) Draw() {
	s.Gutter.Draw()
	s.EditArea.Draw()
	s.Panel.Draw()
	s.CmdLine.Draw()

	// paint viewports
	//s.Gutter.Paint(termbox.ColorGreen)
	//s.EditArea.Paint(termbox.ColorBlue)
	//s.Panel.Paint(termbox.ColorRed)
	//s.CmdLine.Paint(termbox.ColorYellow)

	termbox.Flush()
}

func Init() {
	termbox.Init()
}

func Close() {
	termbox.Close()
}

