package buffer

import (
	"container/list"
	"os"
)

type List struct {
	Buffers list.List
	current *list.Element
}


func NewListFromArgs() *List {
	l := &List{}
	if len(os.Args) == 1 {
		l.Buffers.PushBack(NewEmpty())
	}
	for _, f := range os.Args[1:] {
		info, err := os.Stat(f)
		if err != nil {
			l.Buffers.PushBack(NewEmpty())
		} else {
			if info.IsDir() {
				continue
			}
			b, _ := NewFromFile(f)
			l.Buffers.PushBack(b)
		}
	}

	l.current = l.Buffers.Front()
	return l
}

func (l *List) Current() *T {
	return l.current.Value.(*T)
}
