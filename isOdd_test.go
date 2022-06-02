package isOdd

import "testing"

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func TestOdd_1(t *testing.T) {
	if !IsOdd(3) {
		t.Fatal("Test Failed")
	}
}

func TestOdd_2(t *testing.T) {
	if !IsOdd(MaxInt) {
		t.Fatal("Test Failed")
	}
}

func TestNotOdd_1(t *testing.T) {
	if IsOdd(4) {
		t.Fatal("Test Failed")
	}
}

func TestNotOdd_2(t *testing.T) {
	if IsOdd(MinInt) {
		t.Fatal("Test Failed")
	}
}
