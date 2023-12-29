package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	subtract := subtract(5, 4)
	if subtract != 1 {
		t.Errorf("Subtract was incorrect, got: %d, want: %d.", subtract, 1)
	}
}

func TestDoMath(t *testing.T) {
	doMath := doMath(4, 3, subtract)
	wantNumber := 1
	if doMath != wantNumber {
		t.Errorf("Domath was incorrect, got: %d, want: %d.", doMath, wantNumber)
	}
}
