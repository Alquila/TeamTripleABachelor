package main

import (
	"errors"
	"fmt"
	"log"
)

type testCase struct {
	a [][]int
	b []int
	x []int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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

// result from above test case turns out to be correct to this tolerance.
var ε = 0

func gaussmain() {
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

func GaussPartial(a0 [][]int, b0 []int) ([]int, error) {
	m := len(b0)
	a := make([][]int, m)
	for i, ai := range a0 {
		row := make([]int, m+1)
		copy(row, ai)
		row[m] = b0[i]
		a[i] = row
	}
	for k := range a {
		iMax := 0
		max := -1
		for i := k; i < m; i++ {
			row := a[i]
			// compute scale factor s = max abs in row
			s := -1
			for j := k; j < m; j++ {
				x := Abs(row[j])
				if x > s {
					s = x
				}
			}
			// scale the abs used to pick the pivot.
			if abs := Abs(row[k]) / s; abs > max {
				iMax = i
				max = abs
			}
		}
		if a[iMax][k] == 0 {
			return nil, errors.New("singular")
		}
		a[k], a[iMax] = a[iMax], a[k]
		for i := k + 1; i < m; i++ {
			for j := k + 1; j <= m; j++ {
				a[i][j] -= a[k][j] * (a[i][k] / a[k][k])
			}
			a[i][k] = 0
		}
	}
	x := make([]int, m)
	for i := m - 1; i >= 0; i-- {
		x[i] = a[i][m]
		for j := i + 1; j < m; j++ {
			x[i] -= a[i][j] * x[j]
		}
		x[i] /= a[i][i]
	}
	return x, nil
}
