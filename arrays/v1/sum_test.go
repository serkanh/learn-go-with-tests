package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([5]int{1, 2, 3, 4, 5})
	want := 15
	if got != want {
		t.Errorf("got:%d\nwant:%d", got, want)
	}
}

func ExampleSum() {
	got := Sum([5]int{1, 2, 3, 4, 5})
	fmt.Println(got) //Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([5]int{1, 2, 3, 4, 5})
	}

}
