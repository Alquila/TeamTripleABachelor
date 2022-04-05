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
	fmt.Printf("f1 is: %d \n", f1_bits)
	fmt.Printf("f2 is: %d \n", f2_bits)

	res := XorSlice(f1_bits, f2_bits)

	return res
}

/**
*	Describes the register after initialisation with framenumber 'f2' with the
*	variables used in framenumber 'f1'.
*	Also takes a register with 1 in diagonal ?
 */
func DescribeNewFrameWithOldVariables(f1 int, f2 int, orgReg [][]int) [][]int {

	// gives os bitwise difference of frame numers
	diff := FindDifferenceOfFrameNumbers(f1, f2)
	fmt.Printf("The difference between the two framenumbers are: %d \n", diff)
	// init the predicted new symReg
	length := len(orgReg)
	res := make([][]int, length)

	// for each row in the register
	for i := range orgReg {
		// create the slice that represent the
		res[i] = make([]int, len(orgReg[0]))
		copy(res[i], orgReg[i])
	}

	for i := range diff {
		if diff[i] == 1 { //dvs forskellige frame bits
			fmt.Printf("Diff[%d] is 1\n", i)
			// XOR constant-index in expression
			res[i][len(orgReg[0])-1] = orgReg[i][len(orgReg[0])-1] ^ 1
		}
	}

	return res
}
