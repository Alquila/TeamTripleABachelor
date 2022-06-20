package main

import (
	. "fmt"
	"strconv"
	"strings"
)

// PrettyPrintRegister
// a helping function that Prints registers.
func PrettyPrintRegister(r Register) {
	Printf("%+v", r.RegSlice)
	print("\n")
}

// PrintAllRegisters
// a helping function that Prints all registers.
func PrintAllRegisters() {
	Printf("R1: %+v \n", r1.RegSlice)
	Printf("R2: %+v \n", r2.RegSlice)
	Printf("R3: %+v \n", r3.RegSlice)
	Printf("R4: %+v \n", r4.RegSlice)
}

// PrettyPrintSymRegister
// a helping function that Prints symbolic registers.
func PrettyPrintSymRegister(r SymRegister) {
	rMatrix := r.RegSlice
	rBit := r.SetToOne

	PrettySymPrintSliceBit(rMatrix, rBit)
}

// PrettySymPrintSlice
// a helping function that takes a matrix as input
// and translates the slices into variable names and Prints them.
func PrettySymPrintSlice(slice [][]int) {
	for i := 0; i < len(slice); i++ { //19
		accString := "["
		for j := 0; j < len(slice[0])-1; j++ { //19
			if slice[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += "x" + (str) + " ⨁ "
			}
		}
		accString += strconv.Itoa(slice[i][len(slice[0])-1]) + "]  \n"
		print(accString)
	}
	Println()
}

// PrettySymPrintFrame
// a helping function that takes a matrix as input
// and Prints equations for frames.
func PrettySymPrintFrame(slice [][]int) {
	for i := 0; i < len(slice); i++ { //19
		accString := "["
		for j := 0; j < len(slice[0]); j++ { //19
			if slice[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += "f" + (str) + " ⨁ "
			}
		}
		accString = strings.TrimRight(accString, "⨁ ")
		accString += "]  \n"
		print(accString)
	}
	Println()

}

func PrettySymPrintKeyFrame(slice [][]int) {
	for i := 0; i < len(slice); i++ { //19
		accString := "["
		for j := 0; j < len(slice[0]); j++ { //19
			if slice[i][j] == 1 {
				if j < 64 {
					str := strconv.Itoa(j)
					accString += "k" + (str) + " ⨁ "
				} else {
					str := strconv.Itoa(j - 64)
					accString += "f" + (str) + " ⨁ "
				}
			}
		}
		accString = strings.TrimRight(accString, "⨁ ")
		accString += "]  \n"
		print(accString)
	}
	Println()

}

func PrettySymPrintSliceBit(rMatrix [][]int, bit_entry int) {
	rLength := len(rMatrix)
	for i := 0; i < rLength; i++ {
		accString := "r" + strconv.Itoa(i) + " = "
		for j := 0; j < len(rMatrix[0])-1; j++ {
			if j >= bit_entry {
				if rMatrix[i][j] == 1 {
					str := strconv.Itoa(j + 1)
					accString += " x" + (str) + " ⨁ "
				}
			} else if rMatrix[i][j] == 1 {
				str := strconv.Itoa(j)
				accString += " x" + (str) + " ⨁ "
			}
		}
		accString = accString + strconv.Itoa(rMatrix[i][len(rMatrix[0])-1])
		Printf("")
		println(accString)
	}

}

func Prints(res []int, text string) {
	Printf(text+"%+v \n", res)
}

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		Prints(matrix[i], strconv.Itoa(i))
	}
}
