package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("run with 3 numbers", func(t *testing.T) {
		nums3 := []int{1, 2, 3}
		got := Sum(nums3)
		want := 6
		if got != want {
			t.Errorf("\ngot:%v\nwant:%v", got, want)
		}
	})
	t.Run("run with 4 numbers", func(t *testing.T) {
		nums4 := []int{1, 2, 3, 4}
		got := Sum(nums4)
		want := 10
		if got != want {
			t.Errorf("\n\nTest failed!\ngot:%v\nwant:%v", got, want)
		}
	})
}
