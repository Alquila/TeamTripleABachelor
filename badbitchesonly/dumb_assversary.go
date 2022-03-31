package main

import (
	"fmt"
	"reflect"
)

//"fmt"

func idk() int {
	return 42
}

func doTheSimpleHack() {
	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.ArrImposter)

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSym(symReg)
	keyStream := SimpleKeyStream(reg)

	// make sym version into [][]int if not allready

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res, orgReg) {
		fmt.Printf("This is fucking wrong\n")
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", reg.ArrImposter)
	}
}

func DoTheKnownPlainTextHack() []int {
	// // Init all four Registers
	// initializeRegisters()
	// SymInitializeRegisters()

	// make stream cipher ?
	b := makeKeyStream()
	A := makeSymKeyStream()

	x := solveByGaussEliminationTryTwo(A, b)

	return x
}

/** TRYING TO USE THE DIFFERENCE IN FRAMENUMBER TO
SEE WETHER THE INDEX IN REGISTER SHOULD BE THE
SAME OR DIFFERENT WHEN INITIALIZING IT		   */

func FindDifferenceOfFrameNumbers(f1 int, f2 int) []int {

	f1_bits := MakeFrameNumberToBits(f1)
	f2_bits := MakeFrameNumberToBits(f2)
	res := XorSlice(f1_bits, f2_bits)

	return res
}

func DescribeNewFrameWithOldVariables(f1 int, f2 int, orgReg [][]int) [][]int {

	diff := FindDifferenceOfFrameNumbers(f1, f2)

	// kan vi bare XOR diff med vores originale register ?

	res := make([][]int, len(orgReg))

	for i := range orgReg {
		res[i] = make([]int, len(orgReg[0]))
		copy(res[i], orgReg[i])
	}

	for i := range diff {
		if diff[i] == 1 {
			// XOR constant-index in expression
			for j := range orgReg {
				res[j][len(orgReg[0])-1] = orgReg[j][len(orgReg[0])-1] ^ 1
			}
		}
	}

	return res
}
