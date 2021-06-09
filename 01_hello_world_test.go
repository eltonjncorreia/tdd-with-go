package main

import "testing"

func TestHelloWorld(t *testing.T) {
	params := "Elton"
	result := HelloWorld(params)
	expected := "Hello, Elton"

	if result != expected {
		t.Errorf("result %s, expected %s, input %s", result, expected, params)
	} // t.Fatal("not implemented")
}
