package editor

import (
	"buffer"
	"screen"
)

type T struct {
	EditArea *screen.Viewport
	Panel    *screen.Viewport
	Gutter   *screen.Viewport
	CmdLine  *screen.Viewport

	BufferList *buffer.List
	Command chan string
}


func New() *T {
	// Initialize buffers.
	W, H := screen.Size()
	gw := 4
	pw := 20
	ew := W - (gw + pw) - 2
	editor := &T{
		Gutter:   screen.NewViewport(0,           0, gw, H - 2),
		EditArea: screen.NewViewport(gw + 1,      0, ew, H - 2),
		Panel:    screen.NewViewport(gw + ew + 2, 0, pw, H - 2),
		CmdLine:  screen.NewViewport(0, H  -1, W, 1),
	}
	//editor.Gutter.Fg = termbox.ColorYellow

	// Initialize buffers.
	editor.BufferList = buffer.NewListFromArgs()
	editor.EditArea.Buffer = editor.BufferList.Current()
	editor.Gutter.Buffer, _ = buffer.NewFromStrings([]string{"1", "2", "3", "4"})
	// editor.Panel.Buffer, _ = buffer.NewFromStrings([]string{"a", "b", "c", "d", "e"})
	editor.Panel.Buffer = buffer.NewEmpty()

	editor.Command = make(chan string)

	editor.Draw()
	return editor
}

func (e *T) Draw() {
	e.Gutter.Draw()
	e.EditArea.Draw()
	e.Panel.Draw()
	e.CmdLine.Draw()

	W, H := screen.Size()
	Xa, Xb := e.Gutter.Width, W - e.Panel.Width
	screen.DrawVerticalLine(Xa, 0, H - 3)
	screen.DrawVerticalLine(Xb, 0, H - 3)
	screen.DrawHorizontalLine(0, W - 1, H - 2)
	screen.PutCell(Xa, H - 2, '┴')
	screen.PutCell(Xb, H - 2, '┴')

	// paint viewports
	//s.Gutter.Paint(termbox.ColorGreen)
	//s.EditArea.Paint(termbox.ColorBlue)
	//s.Panel.Paint(termbox.ColorRed)
	//s.CmdLine.Paint(termbox.ColorYellow)

	screen.Refresh()
}

func (e *T) MainLoop() {
loop:	for {
		msg := <-e.Command
		if msg == "quit" {
			break loop
		}
	}
}
