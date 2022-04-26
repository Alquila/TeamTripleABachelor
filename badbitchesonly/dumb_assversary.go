package main

import (
	"fmt"
	"math"
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
	if !reflect.DeepEqual(res.Solved, orgReg) {
		fmt.Printf("This is fucking wrong\n")
		fmt.Printf("Res er: %d\n", res.Solved)
		fmt.Printf("reg er: %d\n", reg.ArrImposter)
	}
}

func DoTheKnownPlainTextHack() ([]int, []int, []int) {
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

	r1_solved, r2_solved, r3_solved := MakeGaussResultToRegisters(x.Solved)

	return r1_solved, r2_solved, r3_solved
}

func MakeGaussResultToRegisters(res []int) ([]int, []int, []int) {
	offset := 0

	r1_res := make([]int, r1.Length-1)
	reg_range := r1.Length - 1
	copy(r1_res, res[:reg_range])
	offset = reg_range
	// fmt.Printf("r1 range: %d. len of r1_res: %d \n", reg_range, len(r1_res))
	// prints(r1_res, "r1_res")
	r2_res := make([]int, r2.Length-1)
	reg_range += r2.Length - 1
	copy(r2_res, res[offset:reg_range])
	offset = reg_range

	r3_res := make([]int, r3.Length-1)
	reg_range += r3.Length - 1
	copy(r3_res, res[offset:reg_range])
	offset = reg_range

	// Move r1
	r1_res = putConstantBackInRes(r1_res, sr1.set1)
	r2_res = putConstantBackInRes(r2_res, sr2.set1)
	r3_res = putConstantBackInRes(r3_res, sr3.set1)

	return r1_res, r2_res, r3_res
}

func putConstantBackInRes(arr []int, constantIndex int) []int {
	arr_size := len(arr)
	newarr := make([]int, arr_size+1)
	copy(newarr, arr)

	for i := (arr_size); i > constantIndex-1; i-- {
		newarr[i] = arr[i-1]
	}
	newarr[constantIndex] = 1

	return newarr
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
	for i := range newReg { // for each entry in the outermost array
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

func TryAllReg4() {
	/*
		"For all possible 2^16 values of R4 solve the linearized system of equations that describe the output.
		Most of the 2^16-1 wrong solutions will be found by inconsistensies in Gauss Elimination.
		The solution of the equations will suggest the internal state of R1, R2, and R3.
		If more than one consistent internal state exists then do trial encryptions."
	*/

	r4_found := make([][]int, 0) // append results to this boi
	r4_guess := make([]int, 17)
	r4_guess[10] = 1

	// r4_real := make([]int, 17) //make an r4 that we want to guess
	// for i := 0; i < 17; i++ { r4_real[i] = rand.Intn(2)	}
	// MakeRealKeyStream()		 //make the actual keystream based on this r4 value

	guesses := int(math.Pow(2, 16))

	for i := 0; i < guesses; i++ {
		r4_guess = MakeR4Guess(i) //for all possible value of r4 we need three frames
		r4_guess = putConstantBackInRes(r4_guess, 10)
		//init sr1 sr2 sr3

	}

	if len(r4_found) > 1 {
		// we have multiple plausible solutions
		// somehow try them all and se what works
		for i := 0; i > len(r4_found); i++ {
			r4.ArrImposter = r4_found[i]

			// do makeKeyStream change r4 ???? I sure hope not :))
			// ks := makeKeyStream()
			messageToEncrypt := make([]int, 184)
			messageToEncrypt[5] = 42
			messageToEncrypt[75] = 42
			messageToEncrypt[150] = 42
			messageToEncrypt[129] = 42

		}
	}

	// 	for i := 0; i < 2; i++ {
	// 		r4_guess[0] = i
	// 		for i := 0; i < 2; i++ {
	// 			r4_guess[1] = i
	// 			for i := 0; i < 2; i++ {
	// 				r4_guess[2] = i
	// 				for i := 0; i < 2; i++ {
	// 					r4_guess[3] = i
	// 					for i := 0; i < 2; i++ {
	// 						r4_guess[4] = i
	// 						for i := 0; i < 2; i++ {
	// 							r4_guess[5] = i
	// 							for i := 0; i < 2; i++ {
	// 								r4_guess[6] = i
	// 								for i := 0; i < 2; i++ {
	// 									r4_guess[7] = i
	// 									for i := 0; i < 2; i++ {
	// 										r4_guess[8] = i
	// 										for i := 0; i < 2; i++ {
	// 											r4_guess[9] = i
	// 											for i := 0; i < 2; i++ {
	// 												r4_guess[11] = i
	// 												for i := 0; i < 2; i++ {
	// 													r4_guess[12] = i
	// 													for i := 0; i < 2; i++ {
	// 														r4_guess[13] = i
	// 														for i := 0; i < 2; i++ {
	// 															r4_guess[14] = i
	// 															for i := 0; i < 2; i++ {
	// 																r4_guess[15] = i
	// 																for i := 0; i < 2; i++ {
	// 																	r4_guess[16] = i
	// 																	//do the gauss or whatever
	// 																}
	// 															}
	// 														}
	// 													}
	// 												}
	// 											}
	// 										}
	// 									}
	// 								}
	// 							}
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

}

func MakeR4Guess(number int) []int {
	r4_bit := make([]int, 16)

	for i := 0; i < 16; i++ {
		r4_bit[i] = (number >> i) & 1 // index 0 becomes least significant bit
	}

	return r4_bit
}
