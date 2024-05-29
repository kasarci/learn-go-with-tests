package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("test","")
		want := "Hello, test!"
	
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World!' when an empty string is supplied.", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {
		got := Hello("test", "Spanish")
		want := "Hola, test!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {
		got := Hello("test", "French")
		want := "Bonjour, test!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
