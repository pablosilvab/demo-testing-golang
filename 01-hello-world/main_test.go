package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Pablo", "English")
		want := "Hello, Pablo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello World when name is empty", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Pablo", "Spanish")
		want := "Hola, Pablo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pablo", "French")
		want := "Bonjour, Pablo"
		assertCorrectMessage(t, got, want)
	})

}

func TestHelloEnglish(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in English", func(t *testing.T) {
		got := Hello("Pablo", "")
		want := "Hello, Pablo"
		assertCorrectMessage(t, got, want)
	})

}
