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

func solveByGaussElimination(A [][]int, b []int) []int {
	augmentMatrix := makeAugmentedMatrix(A, b)
	afterGauss := gaussElimination(augmentMatrix)
	solution := backSubstitution(afterGauss)
	return solution
}

func makeAugmentedMatrix(A [][]int, b []int) [][]int {
	amountOfUnknownVar := len(A) // this is row
	amountOfColumns := len(A[0]) // this is column
	augMa := make([][]int, amountOfUnknownVar)

	for i := 0; i < amountOfUnknownVar; i++ {
		augMa[i] = make([]int, amountOfColumns+1) //@Amalie hvor lange skal de være
	}

	for i := 0; i < amountOfUnknownVar; i++ {
		for j := 0; j < amountOfColumns; j++ {
			augMa[i][j] = A[i][j]
		}
		augMa[i][amountOfColumns] = b[i]
	}
	return augMa
}

func gaussElimination(augMa [][]int) [][]int {
	n := len(augMa) // n is number of unknowns
	// temp := make([][]float64, len(augMa))

	for i := 0; i < n; i++ {
		if float64(augMa[i][i]) == 0.0 { // should 0 be eps ?
			// throw an error
			fmt.Print("Divison by zero encountered")
			continue
		}
		for j := i + 1; j < n; j++ {
			ratio := float64(augMa[j][i]) / float64(augMa[i][i])
			// fmt.Printf("augMa[j][i]: %d \n", augMa[j][i])
			// fmt.Printf("augMa[i][i]: %d \n", augMa[i][i])
			// fmt.Printf("ratio is: %v \n", ratio)

			for k := 0; k < n+1; k++ {
				augMa[j][k] = int(float64(augMa[j][k]) - ratio*float64(augMa[i][k]))
				// fmt.Printf("augMa[j][k]: %d \n", augMa[j][k])
				// fmt.Printf("augMa[i][k]: %d \n", augMa[i][k])
			}
		}
	}
	// temp[0] = []float64(augMa[0])
	return augMa
}

// https://www.codegrepper.com/code-examples/python/gauss+elimination+python+numpy
func backSubstitution(augMatrix [][]int) []int {
	numberOfRows := len(augMatrix) // trying to get the length of rows
	res := make([]int, numberOfRows)
	res[numberOfRows-1] = (augMatrix[numberOfRows-1][numberOfRows]) / (augMatrix[numberOfRows-1][numberOfRows-1]) // x[n-1] = a[n-1][n] / a[n-1][n-1]
	// fmt.Printf("This is res1: %d \n", res)

	for i := numberOfRows - 2; i > -1; i-- {
		res[i] = augMatrix[i][numberOfRows]
		// fmt.Printf("This is res2: %d \n", res)

		for j := i + 1; j < numberOfRows; j++ {
			res[i] = res[i] - augMatrix[i][j]*res[j]
			// fmt.Printf("This is res3: %d \n", res)
		}
		res[i] = res[i] / augMatrix[i][i]
		// fmt.Printf("This is res4: %d \n", res)
	}
	return res
}
