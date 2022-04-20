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

func DoTheKnownPlainTextHack() ([]int, []int, []int, []int) {
	// // Init all four Registers
	// initializeRegisters()
	// SymInitializeRegisters()

	// make stream cipher ?
	sr1.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b1, A1 := RunA5_2()

	current_frame_number++
	sr1.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b2, A2 := RunA5_2()

	current_frame_number++
	sr1.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.ArrImposter = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b3, A3 := RunA5_2()

	// A := make([][]int, 684)
	A := append(A1, A2...)
	A = append(A, A3...)
	// A = append(A, A3...)

	b := append(b1, b2...)
	b = append(b, b3...)

	x := solveByGaussEliminationTryTwo(A, b)

	// I think we should handle stuff here
	// on index set_1, set it to 1 and move rest of slice one
	// do this or each of the registers ?

	r1_solved, r2_solved, r3_solved, r4_solved := MakeGaussResultToRegisters(x)

	return r1_solved, r2_solved, r3_solved, r4_solved
}

func MakeGaussResultToRegisters(res []int) ([]int, []int, []int, []int) {
	offset := 0

	r1_res := make([]int, r1.Length)
	reg_range := r1.Length - 1
	copy(r1_res, res[:reg_range])
	offset = reg_range

	r2_res := make([]int, r2.Length)
	reg_range += r2.Length - 1
	copy(r2_res, res[offset:reg_range])
	offset = reg_range

	r3_res := make([]int, r3.Length)
	reg_range += r3.Length - 1
	copy(r3_res, res[offset:reg_range])
	offset = reg_range

	r4_res := make([]int, r4.Length)
	reg_range += r4.Length - 1
	copy(r4_res, res[offset:reg_range])
	offset = reg_range

	// Move r1
	putConstantBackInRes(r1_res, sr1.set1)
	putConstantBackInRes(r2_res, sr2.set1)
	putConstantBackInRes(r3_res, sr3.set1)
	putConstantBackInRes(r4_res, 10) // hardcoded for register 4 as this has no symbolic representation

	return r1_res, r2_res, r3_res, r4_res
}

func putConstantBackInRes(arr []int, constantIndex int) []int {
	arr_size := len(arr)

	for i := (arr_size - 1); i > constantIndex-1; i-- {
		arr[i] = arr[i-1]
	}
	arr[constantIndex] = 1

	return arr
}

/** TRYING TO USE THE DIFFERENCE IN FRAMENUMBER TO
SEE WETHER THE INDEX IN REGISTER SHOULD BE THE
SAME OR DIFFERENT WHEN INITIALIZING IT		   */
func FindDifferenceOfFrameNumbers(f1 int, f2 int) []int {

	f1_bits := MakeFrameNumberToBits(f1)
	f2_bits := MakeFrameNumberToBits(f2)
	// fmt.Printf("f1 is: %d \n", f1_bits)
	// fmt.Printf("f2 is: %d \n", f2_bits)

	res := XorSlice(f1_bits, f2_bits)

	return res
}

/**
*	Describes the register after initialisation with framenumber 'f2' with the
*	variables used in framenumber 'f1'.
*	Also takes a register with 1 in diagonal ?
*	The provided 'original_reg' should have the last entry as 'compliment' entry in the innermost slice
 */
func DescribeNewFrameWithOldVariables(original_framenum int, current_framenum int, original_reg SymRegister) [][]int {

	// gives os bitwise difference of frame numbers
	diff := FindDifferenceOfFrameNumbers(original_framenum, current_framenum)
	
	// init the predicted new symReg
	length := len(original_reg.ArrImposter)

	/*	Res is used to simulate what indices gets
		affected by the difference in frame number.
		Res should be used to determine which indices need to have their
		'constant' index = 1 after the initialisation process.
	*/
	res := make([]int, length)

	// what to go through every indice in frame-number-difference-array
	for i := range diff {

		// this is copied from cipher_sym.SymCalculateNewBit
		// new bit is the bit that is placed at index 0
		newbit := 0 // newbit is now zero

		// do the feedback function
		for j := range original_reg.Tabs {
			// takes the index corresponding to tab[i] in res and
			// XOR with newbit
			newbit = newbit ^ res[original_reg.Tabs[j]]
		}

		// this is copied from cipher.Clock
		// shift each entry one to the right 
		for j := len(res) - 1; j > 0; j-- {
			res[j] = res[j-1]
		}

		// place the result of the feedback in the first entry in the 
		// resulting array
		res[0] = newbit

		if diff[i] == 1 { //dvs forskellige frame number bits
			// the 'newbit' at index 0 gets influenced by the i'th entry 
			// in current_framenum which differs from original_framenum
			res[0] = res[0] ^ 1
		}
	}

	// this is the register to be returned describing the current
	// frame with varibales from previous frame
	newReg := make([][]int, length)
	for i := range newReg {		// for each entry in the outermost array
		newReg[i] = make([]int, len(original_reg.ArrImposter[0]))
		// copy each 'expression' 
		copy(newReg[i], original_reg.ArrImposter[i])
	}

	// create the new reg from old variables, based on res
	for i := range res {
		if res[i] == 0 {
			continue
		} else {
			newReg[i][len(original_reg.ArrImposter[0])-1] = newReg[i][len(original_reg.ArrImposter[0])-1] ^ 1
		}
	}

	newReg[original_reg.set1] = make([]int, len(original_reg.ArrImposter[0]))
	newReg[original_reg.set1][len(original_reg.ArrImposter[0])-1] = 1

	return newReg
}
