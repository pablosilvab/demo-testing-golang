package main

import "testing"

type TestTable struct {
	input1   int
	input2   int
	expected int
}

func TestHello(t *testing.T) {
	t.Run("hello world", func(t *testing.T) {
		got := Hello("World")
		expected := "Hello World"
		if got != expected {
			t.Errorf("Got: %s , Expected: %s ", got, expected)
		}
	})

	t.Run("hello friend", func(t *testing.T) {
		got := Hello("Peter")
		expected := "Hello Peter"
		if got != expected {
			t.Errorf("Got: %s , Expected: %s ", got, expected)
		}
	})
}
