package buffer

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewFromFile(test *testing.T) {
	f, err := ioutil.TempFile("", "testnewfromfile")
	if err != nil {
		test.Error("Could not create temp file.")
	}
	defer os.Remove(f.Name())
	lines := []string{
		"Lorem ipsum dolor sit amet,",
		"consectetur adipisicing elit, sed",
		"do eiusmod tempor incididunt ut",
		"labore et dolore magna aliqua.",
	}
	for _, line := range lines {
		f.Write([]byte(line))
		f.Write([]byte("\n"))
	}
	f.Close()
	
	buffer, err := NewFromFile(f.Name())

	var current_line *Line

	// Traverse buffer down from the beginning.
	current_line = buffer.first_line
	for i := 0; i < len(lines); i++ {
		expected := lines[i]
		got := current_line.data
		if (expected != got) {
			test.Logf("Traversing file.")
			test.Errorf(`Line %v. Expected: "%v". Got: "%v"`, i, expected, got)
		}
		current_line = current_line.next
	}

	// Traverse buffer up from the end.
	current_line = buffer.last_line
	for i := len(lines) - 1; i >= 0; i-- {
		expected := lines[i]
		got := current_line.data
		if (expected != got) {
			test.Logf("Traversing file in reverse.")
			test.Errorf(`Line %v. Expected: "%v". Got: "%v"`, i, expected, got)
		}
		current_line = current_line.prev
	}
}
