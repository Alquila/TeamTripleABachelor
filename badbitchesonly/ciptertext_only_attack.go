package main

import (
	"fmt"
	"math/rand"
)

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
	// fmt.Printf("Number of rows in first matrix: %d\n", noRows)
	noCol := len(A[0]) // n
	// fmt.Printf("Number of col in first matrix: %d\n", noCol)

	secNoRows := len(B) // p
	// fmt.Printf("Number of rows in second matrix: %d\n", secNoRows)
	secNoCol := len(B[0]) // q
	// fmt.Printf("Number of columns in second matrix: %d\n", secNoCol)

	if noCol != secNoRows {
		fmt.Println("Error: The matrix cannot be multiplied")
		fmt.Printf("Dimensions are %d x %d * %d x %d\n", noRows, noCol, secNoRows, secNoCol)
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

func createRandomMessage(length int) []int {
	msg := make([]int, length)
	for i := 0; i < length; i++ {
		msg[i] = rand.Intn(2) // returns int in [0,2)
	}

	return msg
}

func SliceToMatrix(slice []int) [][]int {
	lenght := len(slice)
	res := make([][]int, 0)
	for i := 0; i < lenght; i++ {
		res = append(res, []int{slice[i]})
	}
	return res
}

func MatrixToSlice(matrix [][]int) []int {
	length := len(matrix)
	res := make([]int, 0)
	for i := 0; i < length; i++ {
		res = append(res, matrix[i][0])
	}
	return res
}

/*
Takes a K_G of size 272 x 456 and SymKeystream of 456*variables.
Multiplies them to a 272*v matrix where the i'th row becomes xor of the rows in symkey for which K_G[i][j]==1
*/
func CalculateKgTimesSymKeyStream(Kg [][]int, symKeyStream [][]int) [][]int {
	number_of_rows := len(Kg)
	number_of_columns := len(Kg[0])
	number_of_variables := len(symKeyStream[0])
	number_of_key_rows := len(symKeyStream)
	res := make([][]int, number_of_rows)

	if number_of_columns != number_of_key_rows {
		fmt.Printf("Dimensions of the given matrices doesn't match.\n")
		fmt.Printf("Dimension of Kg is: %d x %d \n", number_of_rows, number_of_columns)
		fmt.Printf("Dimension of symKeyStream is: %d x %d \n", number_of_key_rows, number_of_variables)
		fmt.Printf("%d is clearly different from %d \n", number_of_rows, number_of_key_rows)
	}

	for i := 0; i < number_of_rows; i++ {
		res[i] = make([]int, number_of_variables)
		for j := 0; j < number_of_columns; j++ {

			if Kg[i][j] == 1 {
				// the corresponding j'th entry in keyStream is part of i'th entry in output
				// symKeyStream[j] indgår
				res[i] = XorSlice(res[i], symKeyStream[j])
			}
		}
	} //løber alle rows igennem,  dvs ca. 272 gange
	//vi har nu en symbolsk matrix der beskriver hvilke variable fra V_f der indgår i ligningerne
	//i bogen (K_G * S)*V_f
	return res
}
