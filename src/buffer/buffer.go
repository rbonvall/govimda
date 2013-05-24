package buffer

import (
	"io/ioutil"
	"os"
	"strings"
)

type Line struct {
	Data string
	next *Line
	prev *Line
}

type T struct {
	first_line *Line
	last_line *Line
	path string
	name string
}

func NewEmpty() *T {
	l := &Line{next: nil, prev: nil}
	b := &T{
		first_line: l,
		last_line: l,
	}
	return b
}

func appendLine(buffer *T, data string) {
	newLine := &Line{Data: data, next: nil}
	if buffer.last_line == nil {
		newLine.prev = nil
		buffer.first_line = newLine
		buffer.last_line = newLine
	} else {
		newLine.prev = buffer.last_line
		buffer.last_line.next = newLine
		buffer.last_line = newLine
	}
}

func NewFromFile(fileName string) (buffer *T, err error) {
	file, err := os.Open(fileName)
	if err != nil { return }

	data, err := ioutil.ReadAll(file)
	if err != nil { return }
	file.Close()

	contents := strings.Split(string(data), "\n")
	contents = contents[:len(contents) - 1]
	buffer, err = NewFromStrings(contents)
	return
}

func NewFromStrings(strs []string) (buffer *T, err error) {
		buffer = new(T)
		for _, s := range strs {
				appendLine(buffer, s)
		}
		return
}

// LineByIndex takes a zero-based line index
// and returns the corresponding line.
// Index 0 refers to the last line of the buffer.
// If the line does not exist, it returns nil.
func (b *T) LineByIndex(index int) (line *Line) {
	line = b.first_line
	for i := 0; i < index; i++ {
		if line.next == nil {
			return nil
		}
		line = line.next
	}
	return
}

// LineByIndexInReverse takes a line index
// and returns the corresponding line in reverse order.
// Index 0 refers to the last line of the buffer.
// If the line does not exist, it returns nil.
func (b *T) LineByIndexInReverse(index int) (line *Line) {
	line = b.last_line
	for i := 0; i < index; i++ {
		if line.prev == nil {
			return nil
		}
		line = line.prev
	}
	return
}

// LineByNumber takes a line number and returns the corresponding line.
// Negative indices count starting from the last line.
// Number 1 refers to the first line, number -1 refers to the last line.
// There is no line 0.
// If the line does not exist, it returns nil.
func (b *T) LineByNumber(nrLine int) (line *Line) {
	switch {
	case nrLine > 0:
		line = b.LineByIndex(nrLine - 1)
	case nrLine < 0:
		line = b.LineByIndexInReverse(-nrLine - 1)
	}
	return
}
