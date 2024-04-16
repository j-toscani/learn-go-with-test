package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to peaople", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMesssage(got, want, t)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMesssage(got, want, t)
	})
}

func assertCorrectMesssage(got string, want string, t testing.TB) {
	t.Helper() // tells go this is a test helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
