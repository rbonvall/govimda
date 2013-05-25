package screen

import (
	"testing"
)

func TestFitStringToWidth(test *testing.T) {
	type case_ struct {
		s string
		w int
		expected string
	}
	cases := []case_{
		{"donkey",  0, ""},
		{"donkey",  4, "donk"},
		{"donkey",  6, "donkey"},
		{"donkey",  8, "donkey  "},
		{"donkey", 10, "donkey    "},
	}
	for _, c := range cases {
		got := fitStringToWidth(c.s, c.w)
		if got != c.expected {
			test.Errorf(`Expected: "%v". Got: "%v"`, c.expected, got)
		}
	}
}

