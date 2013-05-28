package main

import (
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

	screen.Init()
	defer screen.Close()

	s := screen.New()

	// initialize buffers
	s.EditArea.Buffer = currentBuffer
	s.Gutter.Buffer, _ = buffer.NewFromStrings([]string{"1", "2", "3", "4"})
	s.Panel.Buffer, _ = buffer.NewFromStrings([]string{"a", "b", "c", "d", "e"})

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
