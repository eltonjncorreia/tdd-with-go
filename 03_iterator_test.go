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

func TestRepeatN(t *testing.T) {
	inputN := 2
	inputChar := "a"
	result := Repeat(inputChar, inputN)
	expected := "aa"

	if expected != result {
		t.Errorf("expected: %s, result %s, inputN: %d, inputChar: %s", expected, result, inputN, inputChar)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 1000)
	}
}
