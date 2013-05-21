package main

import (
	"github.com/nsf/termbox-go"
	"buffer"
	"fmt"
	"os"
	"screen"
)


func main() {
	var currentBuffer *buffer.Type
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
	vp := screen.NewViewport()
	termbox.Close()

	fmt.Println(currentBuffer)
	fmt.Println(vp)
}
