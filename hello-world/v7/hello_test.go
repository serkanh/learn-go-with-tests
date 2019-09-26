package main
import "testing"
func HelloTest(t *testing.T) {
	const assertFunc = func(t *testing.T, got, want){
		if got != want{
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
}
