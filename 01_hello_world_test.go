package main

import "testing"

func TestHelloWorld(t *testing.T) {
	checkMessage := func(t *testing.T, result, expected, params string) {
		t.Helper()

		if result != expected {
			t.Errorf("result '%s', expected '%s', params: %s", result, expected, params)
		}
	}

	t.Run("Walk Hello for people", func(t *testing.T) {
		params := "Elton"
		result := HelloWorld(params)
		expected := "Hello, Elton"
		checkMessage(t, result, expected, params)
	})

	t.Run("says 'Hello world' when an empty string", func(t *testing.T) {
		result := HelloWorld("")
		expected := "Hello, world"
		checkMessage(t, result, expected, "")
	})
}
