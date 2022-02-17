package main

import (
	_ "fmt"
	"testing"
)

func testPrint(t *testing.T) {
	print("hello world!")
}

func testMajority(t *testing.T) {

	x := majority(0, 0, 0)
	if x != 0 {
		t.Errorf(" x is not 0 but %d", x)
	}

	x = majority(0, 0, 1)
	if x != 0 {
		t.Errorf(" x is not 0 but %d", x)
	}

	x = majority(0, 1, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}

	x = majority(1, 1, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}

	x = majority(1, 0, 1)
	if x != 1 {
		t.Errorf(" x is not 1 but %d", x)
	}
}
