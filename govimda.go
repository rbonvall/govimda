package main

import (
	"screen"
)


func main() {
	screen.Init()

	s := screen.New()
	s.Draw()

	ch := make(chan string)
	go screen.WaitForInput(ch)
	for {
		message := <-ch
		if message == "quit" {
			break
		}
	}

	screen.Close()
}
