package main

import (
	"math"
)

type GaussRes struct {
	ResType string
	TempRes [][]int
	ColNo   []int   // index of empty columns
	DepCol  []int   // index of dependent columns
	Solved  []int   // After backsubstitution
	Multi   [][]int // Multiple solutions
}

const (
	Error    string = "Error"
	Valid    string = "Valid"
	Multi    string = "Multi"
	EmptyCol string = "EmptyCol"
	DepVar   string = "DepVar"
	Both     string = "Both"
)

// SolveByGaussElimination
// Takes a matrix A and a vector b and solves Ax = b for x
// The function returns a GaussRes struct with the solution.
// If no valid solution existed GaussRes.ResType returns Error.
func SolveByGaussElimination(A [][]int, b []int) GaussRes {
	augmentMatrix := MakeAugmentedMatrix(A, b)
	afterGauss := GaussElimination(augmentMatrix)
	if afterGauss.ResType == "Error" {
		// Do nothing
	} else {
		GaussResAfterBack := BackSubstitution(afterGauss)
		return GaussResAfterBack
	}
	return afterGauss
}

// MakeAugmentedMatrix
// Takes a matrix A and a vector b as input and
// returns the composed matrix of the two.
func MakeAugmentedMatrix(A [][]int, b []int) [][]int {
	amountOfRows := len(A)    // this is row
	amountOfVars := len(A[0]) // this is column
	augMa := make([][]int, amountOfRows)

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

// GaussElimination
// Takes an augemnted matrix as input and returns a matrix in echelon form.
func GaussElimination(augMa [][]int) GaussRes {
	// Initialize GaussStruct
	res := GaussRes{
		ResType: Valid,
		TempRes: make([][]int, len(augMa)),
		ColNo:   make([]int, 0),
		DepCol:  make([]int, 0),
	}

	noUnknownVars := len(augMa[0]) - 2 // n is number of unknowns
	noEquations := len(augMa)
	freeVar := make([]int, 0)
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
					for j := 0; j < i; j++ {
						if augMa[j][i] == 1 {
							allZero = false
							if res.ResType == EmptyCol {
								res.ResType = Both
							} else {
								res.ResType = DepVar
							}
							freeVar = append(freeVar, i)
							res.DepCol = append(res.DepCol, i)
						}
					}
					if allZero {
						res.ResType = EmptyCol
						freeVar = append(freeVar, i)
						res.ColNo = append(res.ColNo, i)
					}
				}
			}
		}

		// XOR alle ræker efter r, hvor der står 1
		sliceCopy := make([]int, len(augMa[i]))
		copy(sliceCopy, augMa[i])

		noCol := len(augMa[0])
		for p := s + 1; p < noEquations; p++ {
			if augMa[p][i] == 1 {
				augAfterxor := make([]int, len(augMa[p]))
				for j := 0; j < noCol; j++ {
					augAfterxor[j] = augMa[p][j] ^ sliceCopy[j]
				}
				augMa[p] = augAfterxor
			}
		}
		if len(freeVar) > 0 {
			for _, index := range freeVar {
				if augMa[index][i] == 1 {
					augAfterxor := make([]int, len(augMa[index]))
					for j := 0; j < noCol; j++ {
						augAfterxor[j] = augMa[index][j] ^ sliceCopy[j]
					}
					augMa[index] = augAfterxor
				}
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
	if len(freeVar) > 0 {
		for _, index := range freeVar {
			if augMa[index][bitIndex] == 1 {
				if augMa[index][resIndex] != 1 {
					res.ResType = Error
					res.TempRes = augMa
					return res
				}
			} else if augMa[index][resIndex] == 1 {
				res.ResType = Error
				return res
			}
		}
	}

	res.TempRes = augMa

	return res
}

// BackSubstitution
// receives a GaussRes as input and performs back substitution.
// Returns a GaussRes struct. If back substitution was not possible as the system was not valid
// GaussRes.ResType is set to Error.
// Otherwise, GaussRess is updated and returned with the found R1, R2 and R3 values.
func BackSubstitution(gaussRes GaussRes) GaussRes {
	augMatrix := gaussRes.TempRes
	lengthy := len(augMatrix[0])
	noUnknownVars := lengthy - 2 // n is number of unknowns
	lastCol := lengthy - 1
	bitCol := lengthy - 2
	gaussRes = HandleMulti(gaussRes)
	noMulti := len(gaussRes.Multi)

	for m := 0; m < noMulti; m++ {
		res := make([]int, noUnknownVars)
		copy(res, gaussRes.Multi[m])

		// start from the last variable = aug[x_n][k_n]
		res[noUnknownVars-1] = augMatrix[noUnknownVars-1][lastCol] ^ augMatrix[noUnknownVars-1][bitCol]

		for i := noUnknownVars - 2; i >= 0; i-- { // looks at every row not all zero
			if Contains(gaussRes.ColNo, i) {
				continue //if we are in an free collum then we skip iteration
				//i.e. we need to have a res with both 0 and 1 for this variable
			}
			if Contains(gaussRes.DepCol, i) {
				continue //if we are in an Dependent collum then we skip iteration
				//i.e. we need to have a res with both 0 and 1 for this variable
			}

			res[i] = augMatrix[i][lastCol]
			for j := i + 1; j < noUnknownVars; j++ {
				if augMatrix[i][j] == 1 {
					res[i] = res[i] ^ res[j]
				}
			}
			res[i] = res[i] ^ augMatrix[i][bitCol]

		}

		gaussRes.Multi[m] = res
	}

	return gaussRes
}

// HandleMulti
// handles if there were empty columns or mutually dependent variables found when during GaussElimination.
// Returns the multiple possible solutions for R4.
func HandleMulti(gauss GaussRes) GaussRes {
	gaussRes := gauss
	gaussRes.Multi = make([][]int, 0)
	noMulti := len(gaussRes.Multi)

	if gaussRes.ResType == Both || gaussRes.ResType == EmptyCol {
		gaussRes = HandleEmptyCol(gaussRes)
		noMulti = len(gaussRes.Multi)
	}

	if gaussRes.ResType == Both || gaussRes.ResType == DepVar {
		gaussRes = HandleDepVar(gaussRes)
		noMulti = len(gaussRes.Multi)

	}

	if noMulti < 1 {
		res := make([]int, len(gaussRes.TempRes[0])-2)
		gaussRes.Multi = append(gaussRes.Multi, res)
		noMulti = len(gaussRes.Multi)
	}

	gaussRes.ResType = Multi
	return gaussRes
}

// HandleEmptyCol
// handles if there were empty columns found during GaussElimination.
func HandleEmptyCol(gauss GaussRes) GaussRes {
	gaussRes := gauss
	noEmptyCol := len(gauss.ColNo)
	bitSlice := MakeBitSlice(noEmptyCol)
	newMulti := make([][]int, 0)
	for i := 0; i < len(bitSlice); i++ {
		res := make([]int, len(gaussRes.TempRes[0])-2)
		for j := 0; j < noEmptyCol; j++ {
			index := gauss.ColNo[j]
			res[index] = bitSlice[i][j]
		}
		newMulti = append(newMulti, res)
	}

	newGauss := GaussRes{
		ResType: gaussRes.ResType,
		TempRes: gaussRes.TempRes,
		DepCol:  gaussRes.DepCol,
		ColNo:   gaussRes.ColNo,
	}
	newGauss.Multi = newMulti
	return newGauss
}

// HandleDepVar
// handles if there were mutually dependent variables found during GaussElimination.
func HandleDepVar(gauss GaussRes) GaussRes {
	gaussRes := gauss
	noDepVar := len(gaussRes.DepCol)
	bitSlice := MakeBitSlice(noDepVar)
	noMulti := len(gaussRes.Multi)
	if noMulti < 1 {
		res := make([]int, (len(gaussRes.TempRes[0]) - 2))
		gaussRes.Multi = append(gaussRes.Multi, res)
		noMulti = len(gaussRes.Multi)
	}

	newMulti := make([][]int, 0)
	for m := 0; m < noMulti; m++ {
		for i := 0; i < len(bitSlice); i++ {
			res := make([]int, (len(gaussRes.TempRes[0]) - 2))
			copy(res, gaussRes.Multi[m])

			for j := 0; j < noDepVar; j++ {
				index := gauss.DepCol[j]
				res[index] = bitSlice[i][j]
			}
			newMulti = append(newMulti, res)
		}
	}

	newGauss := GaussRes{
		ResType: gaussRes.ResType,
		TempRes: gaussRes.TempRes,
		DepCol:  gaussRes.DepCol,
		ColNo:   gaussRes.ColNo,
	}
	newGauss.Multi = newMulti
	return newGauss
}

// Contains
// Takes a list of integer returns bool
// Taken from https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// MakeBitSlice
// takes an integer and outputs a slice of slices with all possible combinations
func MakeBitSlice(number int) [][]int {
	bitSlice := make([][]int, 0)
	noOfDiffComb := int(math.Pow(2, float64(number)))
	for j := 0; j < noOfDiffComb; j++ {
		bit := make([]int, number)
		for i := 0; i < number; i++ {
			bit[i] = (j >> i) & 1 // index 0 becomes least significant bit
		}
		bitSlice = append(bitSlice, bit)
	}

	return bitSlice
}
