package main

import "testing"

func TestRepeat(t *testing.T) {
	input := "c"
	result := Repeat(0, input)
	expected := "ccccc"

	if result != expected {
		t.Errorf("expected: %s, result: %s, input: %s", expected, result, input)
	}
}

func TestRepeatN(t *testing.T) {
	inputN := 2
	inputChar := "a"
	result := Repeat(inputN, inputChar)
	expected := "aa"

	if expected != result {
		t.Errorf("expected: %s, result %s, inputN: %d, inputChar: %s", expected, result, inputN, inputChar)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat(0, "a")
	}
}
