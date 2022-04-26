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

func solveByGaussEliminationTryTwo(A [][]int, b []int) GaussRes {
	augmentMatrix := makeAugmentedMatrix(A, b)
	afterGauss := gaussEliminationPart2(augmentMatrix)
	if afterGauss.ResType == "Error" {
		// backSubstitution(afterGauss)
		//return afterGauss
		// Do nothing
	} else if afterGauss.ResType == "Valid" {
		GaussResAfterBack := backSubstitution(afterGauss)
		return GaussResAfterBack
	}
	// fmt.Printf("Gauss: %d\n", solution)
	// fmt.Printf("Gauss lenght: %d\n", len(solution))

	// TODO: Handle more than one solution
	return afterGauss
}

type ResType string

type GaussRes struct {
	ResType ResType //Can be "Error" , "EmptyCol", "Idk", "Both", "Multi"
	TempRes [][]int
	ColNo   []int   //index of empty columns
	DepVar  []int   //index of two dependent variables
	Solved  []int   //After backsubstitution
	Multi   [][]int //Multiple solutions
}

const (
	Request ResType = ""
	Error
	EmptyCol
	IDK
	Both
	Multi
	Valid
)

func makeAugmentedMatrix(A [][]int, b []int) [][]int {
	amountOfRows := len(A)    // this is row
	amountOfVars := len(A[0]) // this is column
	augMa := make([][]int, amountOfRows)
	fmt.Printf("Amount of rows of A %d\n", amountOfRows)
	fmt.Printf("Amount of vars of A: %d\n", amountOfVars)
	fmt.Printf("length of b: %d \n", len(b))

	for i := 0; i < amountOfRows; i++ {
		augMa[i] = make([]int, amountOfVars+1) //@Amalie hvor lange skal de være
	}

	for i := 0; i < amountOfRows; i++ {
		for j := 0; j < amountOfVars; j++ {
			augMa[i][j] = A[i][j]
		}
		augMa[i][amountOfVars] = b[i]
	}

	return augMa
}

// https://stackoverflow.com/questions/11483925/how-to-implementing-gaussian-elimination-for-binary-equations
func gaussEliminationPart2(augMa [][]int) GaussRes {
	// Initialize GaussStruct
	res := GaussRes{
		ResType: Valid,
		TempRes: make([][]int, len(augMa)),
		ColNo:   make([]int, 0),
		DepVar:  make([]int, 0),
	}

	noUnknownVars := len(augMa[0]) - 2 // n is number of unknowns
	noEquations := len(augMa)
	fmt.Printf("len of unknown variable %d \n", noUnknownVars)
	fmt.Printf("len of equations %d \n", noEquations)

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

				// To situations:
				//		First: nulsøjle ergo fri variabel
				//		Second: Der er to variabler afhængige af hinanden (counter positive)
				if r == noEquations-1 {
					allZero := true
					dependent := make([]int, 0)
					emptyCol := make([]int, 0)
					for j := 0; j < i; j++ {
						if augMa[j][i] == 1 {
							allZero = false
							dependent = append(dependent, j)
							dependent = append(dependent, i)
							if res.ResType == EmptyCol {
								res.ResType = Both
							} else {
								res.ResType = IDK
							}
						}
					}
					if allZero {
						if res.ResType == IDK {
							res.ResType = Both
						} else {
							res.ResType = EmptyCol
						}
						emptyCol = append(emptyCol, i)
					}
					res.DepVar = append(res.DepVar, dependent...)
					res.ColNo = append(res.ColNo, emptyCol...)
				}

			}
		}

		// xor alle ræker efter r, hvor der står 1
		sliceCopy := make([]int, len(augMa[i]))
		copy(sliceCopy, augMa[i])
		//fmt.Printf("Row %d, should be 1 in index %d: \n %d \n", s, i, sliceCopy)

		noCol := len(augMa[0])
		for p := s + 1; p < noEquations; p++ {
			// fmt.Printf("j is %d \n", p)
			if augMa[p][i] == 1 {
				augAfterxor := make([]int, len(augMa[p]))
				for j := 0; j < noCol; j++ {
					// fmt.Printf("j is %d, ", j)
					augAfterxor[j] = augMa[p][j] ^ sliceCopy[j]
				}
				augMa[p] = augAfterxor
			}
		}

	}

	// Check if entry in bit column or result column is 1 and return error
	bitIndex := noUnknownVars
	resIndex := bitIndex + 1
	for q := noUnknownVars; q < noEquations; q++ {
		if augMa[q][bitIndex] == 1 {
			if augMa[q][resIndex] != 1 {
				res.ResType = Error
				res.TempRes = augMa
				return res
			}
		} else if augMa[q][resIndex] == 1 {
			res.ResType = Error
			return res
		}
	}

	res.TempRes = augMa

	return res
}

// https://www.codegrepper.com/code-examples/python/gauss+elimination+python+numpy
func backSubstitution(gaussRes GaussRes) GaussRes {
	augMatrix := gaussRes.TempRes
	lenghty := len(augMatrix[0])
	noUnknownVars := lenghty - 2 // n is number of unknowns
	lastCol := lenghty - 1
	bitCol := lenghty - 2
	res := make([]int, noUnknownVars)
	//fmt.Printf("lenghty of augma: %d, noOfUnknownvars: %d, last col at: augMa[%d], bit entry at: augMa[%d] \n", lenghty, noUnknownVars, lastCol, bitCol)

	// printmatrix(augMatrix[:21])
	// prints(augMatrix[noUnknownVars-1], "augma[-1]")

	// start from the last variable = aug[x_n][k_n]
	res[noUnknownVars-1] = augMatrix[noUnknownVars-1][lastCol] ^ augMatrix[noUnknownVars-1][bitCol]

	for i := noUnknownVars - 2; i >= 0; i-- { // looks at every row not all zero
		if gaussRes.ResType == EmptyCol || gaussRes.ResType == Both {
			if contains(gaussRes.ColNo, i) {
				continue //if we are in an free collum then we skip iteration
				//i.e. we need to have a res with both 0 and 1 for this variable
			}
		}
		if gaussRes.ResType == IDK || gaussRes.ResType == Both {
			if contains(gaussRes.DepVar, i) { //dependent vars
				// TODO make helping function
				continue
			}
		}
		res[i] = augMatrix[i][lastCol]
		// prints(augMatrix[i], "")
		// fmt.Printf("res[i] is %d", res[i])
		// prints(res, "res")
		for j := i + 1; j < noUnknownVars; j++ {
			// fmt.Printf("i is %d, j is %d \n", i, j)
			if augMatrix[i][j] == 1 {
				res[i] = res[i] ^ res[j]
			}
		}
		res[i] = res[i] ^ augMatrix[i][bitCol]
		// fmt.Printf("res[i] is %d", res[i])
	}

	// TODO Make helping function to handle empty columns
	gaussRes.Solved = res

	return gaussRes
}

// Taken from https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
