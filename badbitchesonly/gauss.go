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

func solveByGaussEliminationTryTwo(A [][]int, b []int) []int {
	augmentMatrix := makeAugmentedMatrix(A, b)
	afterGauss := gaussEliminationPart2(augmentMatrix)
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

// https://stackoverflow.com/questions/11483925/how-to-implementing-gaussian-elimination-for-binary-equations
func gaussEliminationPart2(augMa [][]int) [][]int {
	noUnknownVars := len(augMa[0]) - 1 // n is number of unknowns //burde nok køre igennem dem alle sammen
	noEquations := len(augMa)
	// temp := make([][]float64, len(augMa))

	for i := 0; i < noUnknownVars; i++ {

		s := i
		if augMa[i][i] == 0 {
			// Håndter at den er 0, dvs.
			for r := i + 1; r < noEquations; r++ {
				// find en hvor den er 1 og byt rækker
				if augMa[r][i] == 1 {
					augCopyTo := make([]int, len(augMa[r]))
					augCopyFrom := make([]int, len(augMa[i]))
					copy(augCopyTo, augMa[r])
					copy(augCopyFrom, augMa[i])
					augMa[r] = augCopyFrom
					augMa[i] = augCopyTo
					s = r
					break
				}
			}
		}

		// xor alle ræker efter r, hvor der står 1
		// TODO: tjek om der skal laves check for at vi ikke er i sidste række når vi starter forloppet
		sliceCopy := make([]int, len(augMa[i]))
		copy(sliceCopy, augMa[i])
		// fmt.Printf("Row %d, should be 1 in index %d: \n %d \n", s, i, sliceCopy)

		for p := s + 1; p < noEquations; p++ {
			if augMa[p][i] == 1 {
				augAfterxor := make([]int, len(augMa[p]))
				for j := 0; j <= noUnknownVars; j++ {
					augAfterxor[j] = augMa[p][j] ^ sliceCopy[j]
				}
				augMa[p] = augAfterxor
			}
		}

	}

	return augMa
}

// https://www.codegrepper.com/code-examples/python/gauss+elimination+python+numpy
func backSubstitution(augMatrix [][]int) []int {
	noUnknownVars := len(augMatrix[0]) - 1 // n is number of unknowns //burde nok køre igennem dem alle sammen
	res := make([]int, noUnknownVars)
	res[noUnknownVars-1] = augMatrix[noUnknownVars-1][noUnknownVars] // either 0 or 1
	// for n := 0; n < 19; n++ {
	// 	fmt.Printf("This is augMatrix: %d \n", augMatrix[n])
	// }

	for i := noUnknownVars - 2; i >= 0; i-- { // looks at every row not all zero
		res[i] = augMatrix[i][noUnknownVars]

		for j := i + 1; j < noUnknownVars; j++ {
			if augMatrix[i][j] == 1 {
				res[i] = res[i] ^ res[j]
			}
		}

	}
	return res
}

// taken from GaussEliminationPart2
// for k := 0; k < noUnknownVars; k++ {
// 	// Ensure we have non-zero entry in augma[k,k]
// 	if augMa[k][k] == 0 {
// 		for j := k + 1; j < noEquations; j++ {
// 			if augMa[j][k] != 0 {
// 				augCopyTo := make([]int, len(augMa[k]))
// 				augCopyFrom := make([]int, len(augMa[j]))
// 				copy(augCopyTo, augMa[k])
// 				copy(augCopyFrom, augMa[j])
// 				augMa[k] = augCopyFrom
// 				augMa[j] = augCopyTo
// 				break
// 			}
// 		}
// 		if augMa[k][k] == 0 {
// 			// Then we have a zero column
// 			// This means we have no solutions or multiple solutinos
// 			fmt.Printf("Zero column, so multiple or no solution")
// 			continue
// 		}
// 		for I := k + 1; I <= n; I++ {
// 			if augMa[I][k] == 1 {
// 				for J := 0; I < n; I++ {
// 					augMa[I][J] = augMa[I][J] ^ augMa[k][J]
// 				}
// 			}
// 		}
// 	}
// }

// nom_row := len(augMa)
// print("nom of row: %d \n", nom_row)
// nom_col := len(augMa[0])
// res := make([]int, nom_row)
// for i := 0; i < nom_row ; i++ {
// 	res[i] = augMa[i][nom_col-1]
// }

// taken from backsubstitution
// if augMatrix[i][noUnknownVars] == 1 {
// 	res[i] = res[i+1] // this is the coefficient next to
// }

// fmt.Printf("first res[i] just set to: %d\n", res[i])
// fmt.Printf("This is res2: %d \n", res)

// for j := i + 1; j < noUnknownVars; j++ {
// 	if augMatrix[i][j] == 1 {
// 		res[i] = res[i] ^ res[j] // this is the coefficient next to
// 	}

// 	// fmt.Printf("second res[i] just set to: %d \n", res[i])
// 	// fmt.Printf("augMatrix[i][j] is: %d\n", augMatrix[i][j])
// 	// fmt.Printf("(augMatrix[i][j] mod 2) is: %v\n", modu)

// }

// fmt.Printf("third res[%d] just set to: %d\n", i, res[i])
// fmt.Printf("This is res4: %d \n", res)

// Not in use
// func gaussElimination(augMa [][]int) [][]int {
// 	n := len(augMa[0]) - 1 // n is number of unknowns
// 	// temp := make([][]float64, len(augMa))

// 	for i := 0; i < n; i++ {
// 		if augMa[i][i] == 0 { // should 0 be eps ?
// 			// throw an error
// 			fmt.Print("Divison by zero encountered \n")
// 			continue
// 		}
// 		for j := i + 1; j < n; j++ {
// 			ratio := augMa[j][i] / augMa[i][i]
// 			//newBit := 1
// 			//if augMa[j][i] == augMa[i][i] {
// 			// newBit = 0
// 			// }

// 			// fmt.Printf("augMa[j][i]: %d \n", augMa[j][i])
// 			// fmt.Printf("augMa[i][i]: %d \n", augMa[i][i])
// 			// fmt.Printf("ratio is: %v \n", ratio)

// 			for k := 0; k < n+1; k++ {
// 				augMa[j][k] = augMa[j][k] - ratio*augMa[i][k]
// 				// fmt.Printf("augMa[j][k]: %d \n", augMa[j][k])
// 				// fmt.Printf("augMa[i][k]: %d \n", augMa[i][k])
// 			}
// 		}
// 	}
// 	// temp[0] = []float64(augMa[0])
// 	return augMa
// }

// func solveByGaussElimination(A [][]int, b []int) []int {
// 	augmentMatrix := makeAugmentedMatrix(A, b)
// 	afterGauss := gaussElimination(augmentMatrix)
// 	solution := backSubstitution(afterGauss)
// 	return solution
// }
