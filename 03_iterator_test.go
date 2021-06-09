package main

import "testing"

func TestRepeat(t *testing.T) {
	input := "c"
	result := Repeat(input)
	expected := "ccccc"

	if result != expected {
		t.Errorf("expected: %s, result: %s, input: %s", expected, result, input)
	}
}
