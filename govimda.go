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

	go screen.WaitForInput(e.Command)
	e.MainLoop()
	screen.Close()
}
