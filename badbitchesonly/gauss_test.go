package main

import (
	"fmt"
	"log"
	"testing"
)

func TestPrint42(t *testing.T) {
	plaintext := 42
	fmt.Printf("%d", plaintext)
}

func TestGauss(t *testing.T) {
	var tc = testCase{
		a: [][]int{
			{1, 0, 0, 1, 0, 0},
			{1, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 1},
			{0, 1, 0, 0, 1, 0},
			{0, 1, 1, 1, 0, 0},
			{0, 1, 1, 0, 1, 0}},
		b: []int{1, 0, 0, 0, 0, 0},
		x: []int{1, 1, 0, 1, 0, 1},
	}

	x, err := GaussPartial(tc.a, tc.b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
	for i, xi := range x {
		if Abs(tc.x[i]-xi) > ε {
			log.Println("out of tolerance")
			log.Fatal("expected", tc.x)
		}
	}
}
