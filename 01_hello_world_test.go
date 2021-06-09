package main

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Walk Hello for people", func(t *testing.T) {
		params := "Elton"
		result := HelloWorld(params)
		expected := "Hello, Elton"

		if result != expected {
			t.Errorf("result '%s', expected '%s', params: %s", result, expected, params)
		}
	})

	t.Run("says 'Hello world' when an empty string", func(t *testing.T) {
		result := HelloWorld("")
		expected := "Hello, world"

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}
