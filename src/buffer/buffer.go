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

type Type struct {
	first_line *Line
	last_line *Line
	path string
	name string
}

func NewEmpty() *Type {
	l := &Line{next: nil, prev: nil}
	b := &Type{
		first_line: l,
		last_line: l,
	}
	return b
}

func appendLine(buffer *Type, data string) {
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

func NewFromFile(fileName string) (buffer *Type, err error) {
	file, err := os.Open(fileName)
	if err != nil { return }

	data, err := ioutil.ReadAll(file)
	if err != nil { return }
	file.Close()

	contents := strings.Split(string(data), "\n")
	contents = contents[:len(contents) - 1]
	buffer = new(Type)
	for _, lineData := range contents {
		appendLine(buffer, lineData)
	}

	return
}

// LineByIndex takes a zero-based line index
// and returns the corresponding line.
// Index 0 refers to the last line of the buffer.
// If the line does not exist, it returns nil.
func (b *Type) LineByIndex(index int) (line *Line) {
	line = b.first_line
	for i := 0; i < index; i++ {
		line = line.next
	}
	return
}

// LineByIndexInReverse takes a line index
// and returns the corresponding line in reverse order.
// Index 0 refers to the last line of the buffer.
// If the line does not exist, it returns nil.
func (b *Type) LineByIndexInReverse(index int) (line *Line) {
	line = b.last_line
	for i := 0; i < index; i++ {
		line = line.prev
	}
	return
}

// LineByNumber takes a line number and returns the corresponding line.
// Negative indices count starting from the last line.
// Number 1 refers to the first line, number -1 refers to the last line.
// There is no line 0.
// If the line does not exist, it returns nil.
func (b *Type) LineByNumber(nrLine int) (line *Line) {
	switch {
	case nrLine > 0:
		line = b.LineByIndex(nrLine - 1)
	case nrLine < 0:
		line = b.LineByIndexInReverse(-nrLine - 1)
	}
	return
}
