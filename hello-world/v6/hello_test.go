package main

import "testing"

func TestHello(t *testing.T) {
	assertFunc := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got:%q\nwant:%q", got, want)
		}
	}
	t.Run("Test empty param should return Hello World!", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"
		assertFunc(t, got, want)
	})
	t.Run("Test empty param should return Holla World!", func(t *testing.T) {
		got := Hello("", spanish)
		want := "Hola, World"
		assertFunc(t, got, want)
	})
	t.Run("Test empty param should return Bonjour World!", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, World"
		assertFunc(t, got, want)
	})
	t.Run("Test empty param should return Hola Serkan!", func(t *testing.T) {
		got := Hello("Serkan", spanish)
		want := "Hola, Serkan"
		assertFunc(t, got, want)
	})
	t.Run("Test empty param should return Bonjour Serkan!", func(t *testing.T) {
		got := Hello("Serkan", french)
		want := "Bonjour, Serkan"
		assertFunc(t, got, want)
	})
}
