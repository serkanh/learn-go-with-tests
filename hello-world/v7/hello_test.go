package main

import "testing"

func HelloTest(t *testing.T) {
	assertFunc := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
	t.Run("If name is not passed as an argument, default to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertFunc(got, want)
	})
}
