package screen

import (
	"github.com/nsf/termbox-go"
)

func handleKey(comm chan string, ch rune, key termbox.Key, mod termbox.Modifier) {
	if ch == 'q' {
		comm <- "quit"
	}
}

func handleResize(comm chan string, width, height int) {

}

func WaitForInput(comm chan string) {
	for {
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventError:
			panic(event.Err)
		case termbox.EventKey:
			handleKey(comm, event.Ch, event.Key, event.Mod)
		case termbox.EventResize:
			handleResize(comm, event.Width, event.Height)
		}
	}
}
