package main

import (
	_ "fmt"
	"testing"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

func TestBitEntry(t *testing.T) {
	reg := InitOneSymRegister()
	PrettySymPrintSlice(reg.ArrImposter)
	Bit_entry(reg)
	PrettySymPrintSlice(reg.ArrImposter)
	for i := 0; i < reg.Length; i++ {
		prints(reg.ArrImposter[i], "")
	}
}

func TestConstPrettyPrint(t *testing.T) {
	reg := InitOneSymRegister()
	Bit_entry(reg)
	PrettyPrint(reg)
}

func TestHowFrames(t *testing.T) {
	reg1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	// ree := make([][]int, 19)
	for i := 0; i < 19; i++ {
		reg1.ArrImposter[i] = make([]int, 22)
		// ree[i][i] = 1
	}
	// printmatrix(reg1.ArrImposter)

	for i := 0; i < 22; i++ {
		SymClock(reg1)
		reg1.ArrImposter[0][i] = 1
		// ree[i][i] = 1
	}

	printmatrix(reg1.ArrImposter)
	PrettySymPrintSlice(reg1.ArrImposter)

}

func printmatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		prints(matrix[i], "")
	}
}
