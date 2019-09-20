package main

import "testing"

func TestHello(t *testing.T) {
	assertFunc := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot:%v\nwant:%v", got, want)
		}
	}
	t.Run("If no args are passed, result should be Hello World!", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"
		assertFunc(t, got, want)
	})
	t.Run("If no args are passed, result should be Hello Serkan!", func(t *testing.T) {
		got := Hello("Serkan")
		want := "Hello Serkan!"
		assertFunc(t, got, want)
	})

}
