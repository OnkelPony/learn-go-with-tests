package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(got string, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in french", func(t *testing.T) {
		got := Hello("Brigitte", "French")
		want := "Bonjour, Brigitte"
		assertCorrectMessage(got, want, t)
	})
	t.Run("in Hebrew", func(t *testing.T) {
		got := Hello("Ester", "Hebrew")
		want := "Shalom, Ester"
		assertCorrectMessage(got, want, t)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(got, want, t)
	})
	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(got, want, t)
	})

	t.Run("Say hello to emty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)

	})
}
