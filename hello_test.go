package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectmessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectmessage(t, got, want)
	})
	t.Run("say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectmessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectmessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Claudie", "French")
		want := "Bonjour, Claudie"
		assertCorrectmessage(t, got, want)
	})
}
