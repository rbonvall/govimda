package buffer

import (
	"io/ioutil"
	"os"
	"testing"
)

var lines = []string{
		"Lorem ipsum dolor sit amet,",
		"consectetur adipisicing elit, sed",
		"do eiusmod tempor incididunt ut",
		"labore et dolore magna aliqua.",
}

func createTestFile() (name string) {
	f, err := ioutil.TempFile("", "testnewfromfile")
	if err != nil {
		panic("Could not create temp file.")
	}
	defer f.Close()
	for _, line := range lines {
		f.Write([]byte(line))
		f.Write([]byte("\n"))
	}
	return f.Name()
}

func TestNewFromFile(test *testing.T) {
	name := createTestFile()
	defer os.Remove(name)
	
	buffer, _ := NewFromFile(name)

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

func TestLineByIndex(test *testing.T) {
	name := createTestFile()
	defer os.Remove(name)
	buffer, _ := NewFromFile(name)

	for i := 0; i < len(lines); i++ {
		expected := lines[i]
		got := buffer.LineByIndex(i).data
		if expected != got {
			test.Errorf(`Line index %v. Expected: "%v". Got: "%v"`, i, expected, got)
		}
	}
}

func TestLineByIndexInReverse(test *testing.T) {
	name := createTestFile()
	defer os.Remove(name)
	buffer, _ := NewFromFile(name)

	for i := 0; i < len(lines); i++ {
		expected := lines[len(lines) - 1 - i]
		got := buffer.LineByIndexInReverse(i).data
		if expected != got {
			test.Errorf(`Line index %v. Expected: "%v". Got: "%v"`, i, expected, got)
		}
	}
}

func TestLineByNumber(test *testing.T) {
	name := createTestFile()
	defer os.Remove(name)
	buffer, _ := NewFromFile(name)

	nrToLine := map[int]*string {
		-5: nil,
		-3: &lines[1],
		-1: &lines[3],
		 0: nil,
		 1: &lines[0],
		 3: &lines[2],
		 5: nil,
	}
	for n, v := range nrToLine {
		line := buffer.LineByNumber(n)
		switch line {
		case nil:
			if v != nil {
				expected := *v
				test.Errorf(`Line number %v. Expected: "%v". Got: nil`, n, expected)
			}
		default:
			got := line.data
			if v != nil {
				expected := *v
				if expected != got {
					test.Errorf(`Line number %v. Expected: "%v". Got: "%v"`, n, expected, got)
				}
			} else {
				test.Errorf(`Line number %v. Expected: nil. Got: "%v"`, n, got)
			}
		}
	}
}
