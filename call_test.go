package util

import (
	"testing"
)

func sum(a, b, c int) int {
	return a + b + c
}

func TestCall(t *testing.T) {
	rets := Call(sum, 1, 2, 3)
	if len(rets) != 1 {
		t.Fatal("Failed getting correct result length.")
	}
	if rets[0].(int) != 6 {
		t.Fatal("Failed getting correct result.")
	}
}

func TestCallWithBadArguments(t *testing.T) {
	if !assertPanic(func(){Call(sum, 1, 2)}) {
		t.Fatal("Failed panicking less argument call.")
	}
	if !assertPanic(func(){Call(sum, 1, 2, 3, 4)}) {
		t.Fatal("Failed panicking more argument call.")
	}
	if !assertPanic(func(){Call(sum, 1, 2, "")}) {
		t.Fatal("Failed panicking more argument call.")
	}
}

func assertPanic(f func ()) bool {
	var panick bool
	func () {
		defer func() {
			if err := recover(); err != nil {
				panick = true
			}
		}()
		f()
	}()
	return panick
}
