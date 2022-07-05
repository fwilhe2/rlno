package main

import (
	"testing"
)

func TestParseValueString(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"101 Tsd. €", 101000},
		{"2,13 Mio. €", 2130000},
	}

	for _, tt := range tests {
		actual := parseValueString(tt.input)

		if actual != tt.want {
			t.Errorf("got %d, want %d", actual, tt.want)
		}
	}
}
