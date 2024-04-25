package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sebas", "")
		want := "Hello, Sebas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sebas", "Spanish")
		want := "Hola, Sebas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Sebas", "French")
		want := "Bonjour, Sebas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Sebas", "German")
		want := "Hallo, Sebas"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
