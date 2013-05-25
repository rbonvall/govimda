package main

import (
	"github.com/nsf/termbox-go"
	"buffer"
	"os"
	"screen"
)


func main() {
	var currentBuffer *buffer.T
	var err error

	if len(os.Args) == 1 {
		currentBuffer = buffer.NewEmpty()
	} else {
		currentBuffer, err =  buffer.NewFromFile(os.Args[1])
		if err != nil {
			currentBuffer = buffer.NewEmpty()
		}
	}

	termbox.Init()
	defer termbox.Close()

	s := screen.New()

	// initialize buffers
	s.EditArea.Buffer = currentBuffer
	s.Gutter.Buffer, _ = buffer.NewFromStrings([]string{"1", "2", "3", "4"})
	s.Panel.Buffer, _ = buffer.NewFromStrings([]string{"a", "b", "c", "d", "e"})

	// paint viewports
	//s.Gutter.Paint(termbox.ColorGreen)
	//s.EditArea.Paint(termbox.ColorBlue)
	//s.Panel.Paint(termbox.ColorRed)
	//s.CmdLine.Paint(termbox.ColorYellow)

	s.Draw()

	ch := make(chan string)
	go screen.WaitForInput(ch)
	for {
		message := <-ch
		if message == "quit" {
			break
		}
	}
}
