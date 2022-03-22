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

func solveByGaussEliminationTryTwo(A [][]int, b []int) []int {
	augmentMatrix := makeAugmentedMatrix(A, b)
	afterGauss := gaussEliminationpart2(augmentMatrix)
	solution := backSubstitution(afterGauss)

	return solution
}

func makeAugmentedMatrix(A [][]int, b []int) [][]int {
	amountOfRows := len(A)       // this is row
	amountOfColumns := len(A[0]) // this is column
	augMa := make([][]int, amountOfRows)
	fmt.Printf("Amount of rows of A %d\n", amountOfRows)
	fmt.Printf("Amount of Columns of A: %d\n", amountOfColumns)
	fmt.Printf("length of b: %d \n", len(b))

	for i := 0; i < amountOfRows; i++ {
		augMa[i] = make([]int, amountOfColumns+1) //@Amalie hvor lange skal de være
	}

	for i := 0; i < amountOfRows; i++ {
		for j := 0; j < amountOfColumns; j++ {
			augMa[i][j] = A[i][j]
		}
		augMa[i][amountOfColumns] = b[i]
	}
	return augMa
}

func gaussElimination(augMa [][]int) [][]int {
	n := len(augMa[0]) - 1 // n is number of unknowns
	// temp := make([][]float64, len(augMa))

	for i := 0; i < n; i++ {
		if augMa[i][i] == 0 { // should 0 be eps ?
			// throw an error
			fmt.Print("Divison by zero encountered \n")
			continue
		}
		for j := i + 1; j < n; j++ {
			ratio := augMa[j][i] / augMa[i][i]
			//newBit := 1
			//if augMa[j][i] == augMa[i][i] {
			// newBit = 0
			// }

			// fmt.Printf("augMa[j][i]: %d \n", augMa[j][i])
			// fmt.Printf("augMa[i][i]: %d \n", augMa[i][i])
			// fmt.Printf("ratio is: %v \n", ratio)

			for k := 0; k < n+1; k++ {
				augMa[j][k] = augMa[j][k] - ratio*augMa[i][k]
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
	numberOfUnknownVars := len(augMatrix[0]) - 1 // trying to get the length of coulumns
	res := make([]int, numberOfUnknownVars)
	res[numberOfUnknownVars-1] = augMatrix[numberOfUnknownVars - 1][numberOfUnknownVars] // either 0 or 1
	fmt.Printf("This is augMatrix: %d \n", augMatrix)

	fmt.Printf("Amount of row: %d\n", len(augMatrix))
	fmt.Printf("agMa[20][19]: %d\n", augMatrix[200][18])

	for i := numberOfUnknownVars - 2; i > -1; i-- { // looks at every row not all zero
		if augMatrix[i][numberOfUnknownVars] == 1 {
			res[i] = res[i+1] // this is the coefficient next to
		}
		
		// fmt.Printf("first res[i] just set to: %d\n", res[i])
		// fmt.Printf("This is res2: %d \n", res)

		for j := i + 1; j < numberOfUnknownVars; j++ {
			if augMatrix[i][j] == 1 {
				res[i] = res[i] ^ res[j] // this is the coefficient next to
			}
					
			// fmt.Printf("second res[i] just set to: %d \n", res[i])
			// fmt.Printf("augMatrix[i][j] is: %d\n", augMatrix[i][j])
			// fmt.Printf("(augMatrix[i][j] mod 2) is: %v\n", modu)

		}
		
		res[i] = res[i] ^ augMatrix[i][numberOfUnknownVars]
		// fmt.Printf("third res[%d] just set to: %d\n", i, res[i])
		// fmt.Printf("This is res4: %d \n", res)
	}
	return res
}

// https://stackoverflow.com/questions/11483925/how-to-implementing-gaussian-elimination-for-binary-equations
func gaussEliminationpart2(augMa [][]int) [][]int {
	n := len(augMa[0]) - 1 // n is number of unknowns
	// temp := make([][]float64, len(augMa))

	for K := 0; K < n; K++ {
		// Ensure we have non-zero entry in augma[k,k]
		if augMa[K][K] == 0 {
			for i := K+1; i <= n; i++ {
				if augMa[i][K] != 0 {
					augCopyTo := make([]int, len(augMa[K]))
					augCopyFrom := make([]int, len(augMa[i]))
					copy(augCopyTo, augMa[K])
					copy(augCopyFrom, augMa[i])
					augMa[K] = augCopyFrom
					augMa[i] = augCopyTo
					break
				}
			}

			if augMa[K][K] == 0 {
				// Then we have a zero column
				// This means we have no solutions or multiple solutinos
				fmt.Printf("Zero column, so multiple or no solution")
				continue
			}
		}

		for I := K + 1; I <= n; I++ {
			if augMa[I][K] == 1 {
				for J := 0; I <= n; I++ {
					augMa[I][J] = augMa[I][J] ^ augMa[K][J]
				}
			}
		}

	}

	// nom_row := len(augMa)
	// print("nom of row: %d \n", nom_row)
	// nom_col := len(augMa[0])
	// res := make([]int, nom_row)
	// for i := 0; i < nom_row ; i++ {
	// 	res[i] = augMa[i][nom_col-1]
	// }

	return augMa
}
