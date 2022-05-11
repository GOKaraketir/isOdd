package isOdd

import "testing"

func TestOdd(t *testing.T) {
	if !IsOdd(3) {
		t.Fatal("Test Failed")
	}
}

func TestNotOdd(t *testing.T) {
	if IsOdd(4) {
		t.Fatal("Test Failed")
	}
}
