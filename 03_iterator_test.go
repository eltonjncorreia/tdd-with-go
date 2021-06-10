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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
