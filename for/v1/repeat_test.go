package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func ExampleRepeat() {
	got := Repeat("a")
	fmt.Println(got)
	// Output:  aaaaa
}
