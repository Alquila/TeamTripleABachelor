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
func DescribeNewFrameWithOldVariables(f1 int, f2 int, orgReg SymRegister) [][]int {

	// gives os bitwise difference of frame numbers
	diff := FindDifferenceOfFrameNumbers(f1, f2)
	fmt.Printf("The difference between the two framenumbers are: %d \n", diff)
	// init the predicted new symReg
	length := len(orgReg.ArrImposter)
	res := make([]int, length)

	/*
		Here res is initialised. Res is used to simulate what indices gets 
		affected by the difference in frame number.
		Res should be used to determine which indices need to have their 
		'constant' index = 1 after the initialisation process. 
	*/
	// for each row in the register
	// for i := range orgReg.ArrImposter {
	// 	// create a slice with all zeroes
	// 	res[i] = make([]int, len(orgReg.ArrImposter[0]))
	// }
	fmt.Printf("This is res after init: %d \n", res)

	// what to go through every indice in frame-number-difference-array
	for i := range diff {
		
		// this is copied from cipher_sym.SymCalculateNewBit
		// new bit is the bit that is placed at index 0
		newbit := 0 // newbit is now zerom.  jLANOÅQFWoåNKDVz

		// do the feedback function
		for i := range orgReg.Tabs {
			// print(i)

			// takes the index corresponding to tab[i] in res and 
			// XOR with newbit
			newbit = newbit ^ res[orgReg.Tabs[i]]
		}

		// this is copied from cipher.Clock
		for i := len(res) - 1; i > 0; i-- { 
			res[i] = res[i-1]
		}

		res[0] = newbit 

		if diff[i] == 1 { //dvs forskellige frame number bits
			fmt.Printf("Diff[%d] is 1\n", i)
			// XOR constant-index in expression
			// REVIEW: This is wrong
			res[0] = res[0] ^ 1
		}
	}

	// this is the register to be returned desribind the current 
	// frame with varibales from previous frame
	newReg := make([][]int, length)
	for i := range newReg {
		newReg[i] = make([]int, len(orgReg.ArrImposter[0])) 
		copy(newReg[i], orgReg.ArrImposter[i])
	}


	// create the new reg from old variables, based on res
	for i := range res {
		if res[i] == 0 {
			continue
		} else {
			newReg[i][len(orgReg.ArrImposter[0])-1] = 1
		}
	}

	newReg[orgReg.set1] = make([]int, len(orgReg.ArrImposter[0]))
	newReg[orgReg.set1][len(orgReg.ArrImposter[0]) - 1] = 1

	return newReg
}
