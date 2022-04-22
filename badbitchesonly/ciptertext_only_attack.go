package main

import "fmt"

// CreateGMatrix
// Creates a matrix with ones in the diagonal
//
func CreateGMatrix() [][]int {
	// Make first slice: 184 columns
	G := make([][]int, 456)

	// Make 184 slices of length 456
	for i := 0; i < 456; i++ {
		colSlice := make([]int, 184)

		// Set diagonal to 1
		if i < 184 {
			colSlice[i] = 1
		}

		G[i] = colSlice
	}

	return G
}

// CreateKgMatrix
// Creates a matrix that multiplied with G returns 0
//
func CreateKgMatrix() [][]int {
	// Make first slice: 184 columns
	KG := make([][]int, 272)

	// Make 184 slices of length 456
	for i := 0; i < 272; i++ {
		colSlice := make([]int, 456)

		// Set diagonal to 1 after 184
		if i > 184 && i < 272 {
			colSlice[i] = 1
		}

		KG[i] = colSlice
	}

	return KG
}

// MultiplyMatrix
// Takes a matrix A and a matrix B and multiplies them.
// Matrix A and B must have dimensions q x n and n x p.
//
func MultiplyMatrix(A [][]int, B [][]int) [][]int {
	noRows := len(A) // m
	fmt.Printf("Number of rows in first matrix: %d\n", noRows)
	noCol := len(A[0]) // n
	fmt.Printf("Number of col in first matrix: %d\n", noCol)

	secNoRows := len(B) // p
	fmt.Printf("Number of rows in second matrix: %d\n", secNoRows)
	secNoCol := len(B[0]) // q
	fmt.Printf("Number of columns in second matrix: %d\n", secNoCol)

	if noCol != secNoRows {
		fmt.Println("Error: The matrix cannot be multiplied")
	}

	// Create result matrix after multiplication
	res := make([][]int, noRows)

	// Initialize inner slices
	for i := 0; i < noRows; i++ {
		res[i] = make([]int, secNoCol)
	}

	midRes := 0
	for i := 0; i < noRows; i++ {
		for j := 0; j < secNoCol; j++ {
			for k := 0; k < secNoRows; k++ {
				midRes = midRes ^ (A[i][k] * B[k][j])
			}
			res[i][j] = midRes
			midRes = 0
		}
	}

	return res
}
