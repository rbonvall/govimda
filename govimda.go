package main

import (
	"editor"
	"screen"
)


func main() {
	screen.Init()

	e := editor.New()
	e.Draw()
	screen.Refresh()

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
