package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
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

/*
	Simulates how the frame difference influences the entries of a register.

	@diff_arr is the frame difference slice
	@register is the register
*/
func simulateClockingWithFrameDifference(diff_arr []int, register Register) []int {
	he := make([]int, register.Length)

	for i := 0; i < 22; i++ {
		yas := 0
		for j := 0; j < len(register.Tabs); j++ {
			yas = he[register.Tabs[i]] ^ yas
		}

		he[0] = diff_arr[i] ^ yas
	}

	return he
}

//Creates a new r4 register and initialises it with the difference between the current and original frame number.
//Returns the ArrImposter of the register wich contains 1's in the place where the bits have been flipped with the new frame.
func simulateClockingR4WithFrameDifference(original_frame_number int, current_frame int) []int {
	fake_r4 := makeR4()
	diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame)
	for i := 0; i < 22; i++ {
		Clock(fake_r4)
		fake_r4.ArrImposter[0] = fake_r4.ArrImposter[0] ^ diff[i]
	}
	return fake_r4.ArrImposter
}

/**
FindDifferenceOfFrameNumbers
TRYING TO USE THE DIFFERENCE IN FRAMENUMBER TO
SEE WETHER THE INDEX IN REGISTER SHOULD BE THE
SAME OR DIFFERENT WHEN INITIALIZING IT
*/
func FindDifferenceOfFrameNumbers(f1 int, f2 int) []int {

	f1_bits := MakeFrameNumberToBits(f1)
	f2_bits := MakeFrameNumberToBits(f2)
	// fmt.Printf("f1 is: %d \n", f1_bits)
	// fmt.Printf("f2 is: %d \n", f2_bits)

	res := XorSlice(f1_bits, f2_bits)

	return res
}

/**
DescribeNewFrameWithOldVariable
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

//Returns the real r4 that made the keystream. Returns three frame long key. Returns fourth frame key
func MakeRealKeyStreamFourFrames(frame int) ([]int, []int, []int) {
	original_frame_number = frame
	current_frame_number = frame
	key1 := makeKeyStream()
	// prints(r4.ArrImposter, "r4 after makeKeyStream:     					")
	r4_real := make([]int, 17)
	copy(r4_real, r4_after_init.ArrImposter)

	current_frame_number++
	key2 := makeKeyStream()
	// r4_second := r4.ArrImposter
	// prints(r4_after_init.ArrImposter, "r4 after second init:       ") //[0 1 0 1 0 0 1 0 1 1 1 0 0 0 0 0 1]
	current_frame_number++
	key3 := makeKeyStream()

	current_frame_number++
	key4 := makeKeyStream()
	// prints(r4_after_init.ArrImposter, "r4 after third init:       ")

	key := append(key1, key2...)
	key = append(key, key3...)
	return r4_real, key, key4
}

/*Makes a six frame long key stream and returns it along the initial r4 value. returns two extra frame for test.
Returns r4_real, key, extra_key */
func MakeRealKeyStreamSixFrames(frame int) ([]int, []int, []int) {
	original_frame_number = frame
	current_frame_number = frame
	key := make([]int, 0)
	key1 := makeKeyStream()
	// prints(r4.ArrImposter, "r4 after makeKeyStream:     					")
	r4_real := make([]int, 17)
	copy(r4_real, r4_after_init.ArrImposter)
	key = append(key, key1...)

	for i := 0; i < 5; i++ {
		current_frame_number++
		newKeyStream := makeKeyStream()
		key = append(key, newKeyStream...)
	}

	current_frame_number++
	extra_key_stream := makeKeyStream()
	current_frame_number++
	extra_key_stream2 := makeKeyStream()
	extra_key_stream = append(extra_key_stream, extra_key_stream2...)

	return r4_real, key, extra_key_stream
}

func CalculateRealIteration(r4 []int) int {
	reeee := make([]int, 0)
	reeee = append(reeee, r4[:10]...)
	reeee = append(reeee, r4[11:]...)
	real_iteration := convertBinaryToDecimal(reeee)
	return real_iteration
}

func DescribeRegistersFromKey() [][]int {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.ArrImposter[i] = make([]int, 64)
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.ArrImposter[i] = make([]int, 64)
	}

	for i := 0; i < 64; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
		sre1.ArrImposter[0][i] = 1 //should this be xor? <- no den påvirkes kun af den i'te bit én gang
		sre2.ArrImposter[0][i] = 1
		sre3.ArrImposter[0][i] = 1
		reg4.ArrImposter[0][i] = 1
	}

	symbolicDescription := make([][]int, 0)
	symbolicDescription = append(symbolicDescription, sre1.ArrImposter...)
	symbolicDescription = append(symbolicDescription, sre2.ArrImposter...)
	symbolicDescription = append(symbolicDescription, sre3.ArrImposter...)
	symbolicDescription = append(symbolicDescription, reg4.ArrImposter...)

	return symbolicDescription

}

func TryAllReg4() {
	/*
		"For all possible 2^16 values of R4 solve the linearized system of equations that describe the output.
		Most of the 2^16-1 wrong solutions will be found by inconsistensies in Gauss Elimination.
		The solution of the equations will suggest the internal state of R1, R2, and R3.
		If more than one consistent internal state exists then do trial encryptions."
	*/

	r4_found := make([][]int, 0) // append results to this boi
	// solved := make([][]int, 0)
	r4_guess := make([]int, 17)

	session_key = make([]int, 64) //all zero session key
	// makeSessionKey()              //Make a random session key
	original_frame_number = 42
	r4_real, real_key, r4_for_test := MakeRealKeyStreamFourFrames(original_frame_number)
	//FIXME: we need to make a 4'th real key stream for testing if the found r4 values are correct

	// current_frame_number++

	real_iteration := CalculateRealIteration(r4_real)
	lower := real_iteration - 150
	upper := real_iteration + 150
	fmt.Printf("real: %d, lower: %d, upper: %d\n", real_iteration, lower, upper)
	//[0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1] <- dette er R4 som vi skal frem til når der ikke er noget random
	//[0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1] <- 33114 omgang

	guesses := int(math.Pow(2, 16))
	println(guesses)
	// for i := lower; i < upper; i++ {
	for i := 0; i < guesses; i++ { //FIXME ind og udkommenter de to headers her for at skifte -AK
		if i%100 == 0 {
			fmt.Printf("iteration %d \n", i)
		}
		if i == real_iteration {
			fmt.Printf("iteration %d\n", real_iteration)
		}
		if i == real_iteration+1 {
			fmt.Printf("iteration %d\n", real_iteration+1)
		}
		original_frame_number = 42 //reset the framenumber for the symbolic version
		current_frame_number = 42

		r4_guess = MakeR4Guess(i) //for all possible value of r4 we need three frames
		r4_guess = putConstantBackInRes(r4_guess, 10)
		// prints(r4_guess, "r4_guess")

		//do this such that r4 guess can be copied into sr4 in SymSetRegisters()
		r4 = makeR4()
		copy(r4.ArrImposter, r4_guess)
		// prints(r4.ArrImposter, "r4_guess1 ")
		key1 := makeSymKeyStream() //this clocks sr4 which has r4_guess as its array
		// prints(sr4.ArrImposter, "sr4 after makeSymkey  						")

		current_frame_number++

		//update r4_guess with new frame value //we want it to be clean right..??
		// prints(r4_guess, "r4_guess")
		r4 = makeR4()
		// fake_r4 := makeR4()
		copy(r4.ArrImposter, r4_guess)

		frame_influenced_bits := simulateClockingR4WithFrameDifference(original_frame_number, current_frame_number)
		r4.ArrImposter = XorSlice(frame_influenced_bits, r4_guess)

		// diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame_number)
		// for i := 0; i < 22; i++ {
		// 	Clock(fake_r4)
		// 	fake_r4.ArrImposter[0] = fake_r4.ArrImposter[0] ^ diff[i]
		// } //fake_r4.ArrImposter er nu clocked således at det er [...1..] de steder hvor diff påvirker indgangene
		// r4.ArrImposter = XorSlice(fake_r4.ArrImposter, r4.ArrImposter)
		// fake_r4.ArrImposter[10] = 1 //FIXME
		r4.ArrImposter[10] = 1
		//FIXME
		//FIXME
		//FIXME

		//we want this -> [0 1 0 1 0 0 1 0 1 1 1 0 0 0 0 0 1]
		// prints(r4.ArrImposter, "sr4_guess init ")
		key2 := makeSymKeyStream() //this will now copy the updated r4_arrimposter into sr4
		// prints(r4_second, "r4_second ")
		// prints(sr4.ArrImposter, "sr4_after2")

		current_frame_number++
		r4 = makeR4()
		fake_r4 := makeR4()
		copy(r4.ArrImposter, r4_guess)
		diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame_number)
		for i := 0; i < 22; i++ {
			Clock(fake_r4)
			fake_r4.ArrImposter[0] = fake_r4.ArrImposter[0] ^ diff[i]
		} //fake_r4.ArrImposter er nu clocked således at det er [...1..] de steder hvor diff påvirker indgangene
		r4.ArrImposter = XorSlice(fake_r4.ArrImposter, r4.ArrImposter)
		r4.ArrImposter[10] = 1
		//prints(r4.ArrImposter, " sr4 after third")
		key3 := makeSymKeyStream()
		current_frame_number++

		key := append(key1, key2...)
		key = append(key, key3...)

		// this returns a gauss struct
		gauss := solveByGaussEliminationTryTwo(key, real_key)

		if gauss.ResType == Error {
			continue
		} else if gauss.ResType == Multi {
			fmt.Printf("found multi in %d of lenght %d \n", i, len(gauss.Multi))
			for i := 0; i < len(gauss.Multi); i++ {
				if VerifyKeyStream(gauss.Multi[i]) { ///what do we do here
					r4_found = append(r4_found, r4_guess)
					// fmt.Printf("found in ")
					// solved = append(solved, gauss.Multi[i])
					// solved = append(solved, []int{42})

				}
			}
		}
		//init sr1 sr2 sr3
		//make first sym-keystream based on r4 guess and symbol registers
		//key1 := makeSymKeyStream()
		//framenumber ++
		//init sr1 sr2 sr3 again with the new framenumber
		//init r4_guess with the new frame number
		//key2 := makeSymKeyStream()
		//framenumber ++
		//init sr1 sr2 sr3 again with the new framenumber
		//init r4_guess with the new frame number
		//key3 := makeSymKeyStream()

		// gauss: based on response add to found
	}

	// FIXME: this might not work ?
	// 'trial encryptions'
	correct_r4 := make([]int, len(r4_guess))
	number_of_valid_r4 := len(r4_found)
	if number_of_valid_r4 <= 0 {
		fmt.Printf("We didn't find any at all \n")
	} else if number_of_valid_r4 > 1 {
		// we have multiple plausible solutions to r4
		// somehow try them all and se what works
		for i := 0; i > number_of_valid_r4; i++ {
			r4.ArrImposter = r4_found[i] // is this how its supposed to happend?
			//what should frame_number be ? original + 4
			current_frame_number = original_frame_number + 4
			ks := makeKeyStream()
			if reflect.DeepEqual(ks, r4_for_test) {
				fmt.Printf("This should be the right one: %d\n", r4_for_test)
				correct_r4 = r4_found[i]
				break
			}

		}
	} else {
		correct_r4 = r4_found[0]
	}

	fmt.Printf("This is original r4:       %d\n", r4_real)
	for i := range r4_found {
		fmt.Printf("This is %d'th found r4:    %d\n", i, r4_found[i])
		// fmt.Printf("This is %d'th found solved:  %d \n", i, solved[i])
	}
	fmt.Println("Have we found the right r4?")
	// for i := range r4_found {
	if reflect.DeepEqual(correct_r4, r4_real) { // 'correct_r4' used to be 'r4_found[i]'
		fmt.Println("Fuck yes we found it gutterne")
	}
	// }

}

func MakeR4Guess(number int) []int {
	r4_bit := make([]int, 16)

	for i := 0; i < 16; i++ {
		r4_bit[i] = (number >> i) & 1 // index 0 becomes least significant bit
	}

	return r4_bit
}

//VerifyKeyStream compares the found vars with the products that involves them and check that they match up.
func VerifyKeyStream(key []int) bool {
	//[vars1 | vars2 | vars3 | prod1 | prod2 | prod3 | b ]
	vars1_len := r1.Length - 1
	vars2_len := r2.Length - 1
	vars3_len := r3.Length - 1

	prod1_len := vars1_len * (vars1_len - 1) / 2
	prod2_len := vars2_len * (vars2_len - 1) / 2
	prod3_len := vars3_len * (vars3_len - 1) / 2
	// fmt.Printf("vars1_len : %d  vars2_len %d, vars3_len: %d, prod1_len: %d, prod2: %d, prod3: %d \n", vars1_len, vars2_len, vars3_len, prod1_len, prod2_len, prod3_len)
	prod1 := key[vars1_len+vars2_len+vars3_len : vars1_len+vars2_len+vars3_len+prod1_len]
	prod2 := key[vars1_len+vars2_len+vars3_len+prod1_len : vars1_len+vars2_len+vars3_len+prod1_len+prod2_len]
	prod3 := key[vars1_len+vars2_len+vars3_len+prod1_len+prod2_len : vars1_len+vars2_len+vars3_len+prod1_len+prod2_len+prod3_len]
	// prints(prod1, "prod1")
	// print("\n")
	// prints(prod2, "prod2")
	// print("\n")
	// prints(prod3, "prod3")

	helper(key[0:vars1_len], prod1)
	helper(key[vars1_len:vars1_len+vars2_len], prod2)
	helper(key[vars1_len+vars2_len:vars1_len+vars2_len+vars3_len], prod3)

	return true

}

func helper(vars []int, prods []int) bool {
	acc := 0
	for i := 0; i < len(vars); i++ {
		var_1 := vars[i]
		for j := i + 1; j < len(vars); j++ {
			var_2 := vars[j] // i and j runs over the vars variables
			if var_2*var_1 != prods[acc] {
				return false
			}
			// fmt.Printf("var1: %d * var2: %d = prod1[%d]: %d \n", var_1, var_2, acc, prods[acc])
			// fmt.Printf(" %d * %d = %d \n", var_1, var_2, prods[acc])
			acc++ //acc runs over the index in prod1
		}
	}

	return true
}

func convertBinaryToDecimal(number []int) int {
	bin_num := ""
	// prints(number, "reee")

	for i := len(number) - 1; i >= 0; i-- {
		// println(i)
		// println(number[i])
		bin_num = bin_num + strconv.Itoa(number[i])
	}

	num, err := strconv.ParseInt(bin_num, 2, 64)

	if err != nil {
		panic(err)
	}
	return int(num)
}

func convertBinaryToDecimal2(number []int) int {
	bin_num := ""
	// prints(number, "reee")

	for i := 0; i < len(number); i++ {
		// println(i)
		// println(number[i])
		bin_num = bin_num + strconv.Itoa(number[i])
	}

	num, err := strconv.ParseInt(bin_num, 2, 64)

	if err != nil {
		panic(err)
	}
	return int(num)
}

func RetrieveSessionKey(registers []int) []int {

	skey := make([]int, 0)

	symkey := DescribeRegistersFromKey() //
	gauss := solveByGaussEliminationTryTwo(symkey, registers)
	println(gauss.ResType)
	if gauss.ResType == Multi {
		skey = gauss.Multi[0]
	}

	return skey
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
