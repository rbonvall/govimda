package buffer

import (
	"io/ioutil"
	"os"
	"strings"
)

type Line struct {
	data string
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
	newLine := &Line{data: data, next: nil}
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
