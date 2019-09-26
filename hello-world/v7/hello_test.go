package main

import "testing"

func TestHello(t *testing.T) {
	assertFunc := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}
	}
	t.Run("If no name and lang passed as an argument, language should default to spanish and name should default World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertFunc(got, want)
	})
	t.Run("If only name param is passed it should return by defaulting to English and name passed", func(t *testing.T) {
		got := Hello("", "Serkan")
		want := "Hello Serkan"
		assertFunc(got, want)
	})
	t.Run("If name param and spanish is passed it should return Hola Serkan", func(t *testing.T) {
		got := Hello("spanish", "Serkan")
		want := "Hola Serkan"
		assertFunc(got, want)
	})
}
