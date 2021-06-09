package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	n1, n2 := 2, 2
	sum := Add(n1, n2)
	expected := 4

	if sum != expected {
		t.Errorf("sum: %d, expected: %d, input one: %d, input two %d", sum, expected, n1, n2)
	}
}

func TestSub(t *testing.T) {
	x, y := 4, 2
	sub := Sub(x, y)
	expected := 2

	if expected != sub {
		t.Errorf("expected: %d, sub: %d, x: %d, y: %d", expected, sub, x, y)
	}
}
