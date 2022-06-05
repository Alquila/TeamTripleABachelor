package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// DoTheKnownPlainTextHack
// first describes equations from three frame numbers with the same variable names as the base frame.
// Solves the resulting system of equations using Gauss Elimination.
// Returns solved R1, R2 and R3.
func DoTheKnownPlainTextHack() ([]int, []int, []int) {
	sr1.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b1, A1 := RunA5_2()

	current_frame_number++
	sr1.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b2, A2 := RunA5_2()

	current_frame_number++
	sr1.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr1)
	sr2.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr2)
	sr3.RegSlice = DescribeNewFrameWithOldVariables(original_frame_number, current_frame_number, sr3)
	b3, A3 := RunA5_2()

	A := append(A1, A2...)
	A = append(A, A3...)

	b := append(b1, b2...)
	b = append(b, b3...)

	x := solveByGaussEliminationTryTwo(A, b)

	r1_solved, r2_solved, r3_solved := MakeGaussResultToRegisters(x.Solved)

	return r1_solved, r2_solved, r3_solved
}

// MakeGaussResultToRegisters
// receives a Gauss elimination result slice as input.
// Returns the solved R1, R2 and R3.
func MakeGaussResultToRegisters(res []int) ([]int, []int, []int) {
	offset := 0

	r1Res := make([]int, r1.Length-1)
	regRange := r1.Length - 1
	copy(r1Res, res[:regRange])
	offset = regRange

	r2Res := make([]int, r2.Length-1)
	regRange += r2.Length - 1
	copy(r2Res, res[offset:regRange])
	offset = regRange

	r3Res := make([]int, r3.Length-1)
	regRange += r3.Length - 1
	copy(r3Res, res[offset:regRange])
	offset = regRange

	r1Res = PutConstantBackInRes(r1Res, sr1.SetToOne)
	r2Res = PutConstantBackInRes(r2Res, sr2.SetToOne)
	r3Res = PutConstantBackInRes(r3Res, sr3.SetToOne)

	return r1Res, r2Res, r3Res
}

// PutConstantBackInRes
// takes an array and constant index as input and puts 1 in the given index and returns the new array.
func PutConstantBackInRes(arr []int, constantIndex int) []int {
	arrSize := len(arr)
	newArray := make([]int, arrSize+1)
	copy(newArray, arr)

	for i := arrSize; i > constantIndex-1; i-- {
		newArray[i] = arr[i-1]
	}
	newArray[constantIndex] = 1

	return newArray
}

// SimulateClockingR4WithFrameDifference
// creates a new r4 register and initialises it with the difference between the current and original frame number.
// Returns the RegSlice of the register which contains 1's in the place where the bits have been flipped with the new frame.
func SimulateClockingR4WithFrameDifference(original_frame_number int, current_frame int) []int {
	fake_r4 := MakeR4()
	diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame)
	for i := 0; i < 22; i++ {
		Clock(fake_r4)
		fake_r4.RegSlice[0] = fake_r4.RegSlice[0] ^ diff[i]
	}
	return fake_r4.RegSlice
}

//	FindDifferenceOfFrameNumbers
//	trying to use the difference in frame number to
//	see whether the index in a register should be the
//	same or different when initialising it
func FindDifferenceOfFrameNumbers(f1 int, f2 int) []int {
	f1_bits := MakeFrameNumberToBits(f1)
	f2_bits := MakeFrameNumberToBits(f2)
	res := XorSlice(f1_bits, f2_bits)
	return res
}

/*
	DescribeNewFrameWithOldVariables
	Describes the register after initialisation with frame number 'f2' with the
	variables used in frame number 'f1'.
	The provided 'original_reg' should have the last entry as 'compliment' entry in the innermost slice
*/
func DescribeNewFrameWithOldVariables(original_framenum int, current_framenum int, original_reg SymRegister) [][]int {

	// gives os bitwise difference of frame numbers
	diff := FindDifferenceOfFrameNumbers(original_framenum, current_framenum)

	// init the predicted new symReg
	length := len(original_reg.RegSlice)

	/*
		Res is used to simulate what indices gets affected by the difference in frame number.
		Res should be used to determine which indices need to have their 'constant'
		index = 1 after the initialisation process.
	*/
	res := make([]int, length)

	// what to go through every index in frame-number-difference-array
	for i := range diff {

		// new bit is the bit that is placed at index 0
		newbit := 0 // newbit is now zero

		// do the feedback function
		for j := range original_reg.Taps {
			// takes the index corresponding to tap[i] in res and
			// XORs with newbit
			newbit = newbit ^ res[original_reg.Taps[j]]
		}

		// shift each entry one to the right
		for j := len(res) - 1; j > 0; j-- {
			res[j] = res[j-1]
		}

		// place the result of the feedback in the first entry in the resulting array
		res[0] = newbit

		if diff[i] == 1 {
			// the 'newbit' at index 0 gets influenced by the i'th entry
			// in current_framenumber which differs from original_framenumber
			res[0] = res[0] ^ 1
		}
	}

	// this is the register to be returned describing the current
	// frame with variables from previous frame
	newReg := make([][]int, length)
	for i := range newReg { // for each entry in the outermost array
		newReg[i] = make([]int, len(original_reg.RegSlice[0]))

		// copy each 'expression'
		copy(newReg[i], original_reg.RegSlice[i])
	}

	// create the new reg from old variables, based on res
	for i := range res {
		if res[i] == 0 {
			continue
		} else {
			newReg[i][len(original_reg.RegSlice[0])-1] = newReg[i][len(original_reg.RegSlice[0])-1] ^ 1
		}
	}

	newReg[original_reg.SetToOne] = make([]int, len(original_reg.RegSlice[0]))
	newReg[original_reg.SetToOne][len(original_reg.RegSlice[0])-1] = 1

	return newReg
}

// MakeRealKeyStreamFourFrames
// Returns the real r4 that made the keystream.
// Returns a three frame long key. Returns fourth frame key
func MakeRealKeyStreamFourFrames(frame int) ([]int, []int, []int) {
	original_frame_number = frame
	current_frame_number = frame
	key1 := MakeKeyStream()
	r4_real := make([]int, 17)
	copy(r4_real, r4_after_init.RegSlice)

	current_frame_number++
	key2 := MakeKeyStream()
	current_frame_number++
	key3 := MakeKeyStream()

	current_frame_number++
	key4 := MakeKeyStream()

	key := append(key1, key2...)
	key = append(key, key3...)
	return r4_real, key, key4
}

// MakeRealKeyStreamSixFrames
// makes a six frame long key stream and returns it along the initial r4 value.
// Returns two extra frame for test. Returns r4_real, key, extra_key
func MakeRealKeyStreamSixFrames(frame int) ([]int, []int, []int) {
	original_frame_number = frame
	current_frame_number = frame
	key := make([]int, 0)
	key1 := MakeKeyStream()
	r4Real := make([]int, 17)
	copy(r4Real, r4_after_init.RegSlice)
	key = append(key, key1...)

	for i := 0; i < 5; i++ {
		current_frame_number++
		newKeyStream := MakeKeyStream()
		key = append(key, newKeyStream...)
	}

	current_frame_number++
	extraKeyStream := MakeKeyStream()
	current_frame_number++
	extraKeyStream2 := MakeKeyStream()
	extraKeyStream = append(extraKeyStream, extraKeyStream2...)

	return r4Real, key, extraKeyStream
}

// CalculateRealIteration
// splits r4's slice into the first [0,9] bits and the [10, :]
// and converts this binary array to a decimal number
func CalculateRealIteration(r4 []int) int {
	bitSlice := make([]int, 0)
	bitSlice = append(bitSlice, r4[:10]...)
	bitSlice = append(bitSlice, r4[11:]...)
	real_iteration := ConvertBinaryToDecimal(bitSlice)
	return real_iteration
}

// DescribeRegistersFromKey
// creates a symbolic representation of the registers.
// Returns a single matrix with all four symbolic registers.
func DescribeRegistersFromKey() [][]int {
	sre1 := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15)
	sre2 := SymMakeRegister(22, []int{21, 20}, []int{9, 13}, 16, 16)
	sre3 := SymMakeRegister(23, []int{22, 21, 20, 7}, []int{16, 18}, 13, 18)

	for i := 0; i < sre1.Length; i++ {
		sre1.RegSlice[i] = make([]int, 64)
	}

	for i := 0; i < sre2.Length; i++ {
		sre2.RegSlice[i] = make([]int, 64)
	}

	for i := 0; i < sre3.Length; i++ {
		sre3.RegSlice[i] = make([]int, 64)
	}

	reg4 := SymMakeRegister(17, []int{16, 11}, []int{12, 15}, 14, 10)
	for i := 0; i < 17; i++ {
		reg4.RegSlice[i] = make([]int, 64)
	}

	for i := 0; i < 64; i++ {
		SymClock(sre1)
		SymClock(sre2)
		SymClock(sre3)
		SymClock(reg4)
		sre1.RegSlice[0][i] = 1
		sre2.RegSlice[0][i] = 1
		sre3.RegSlice[0][i] = 1
		reg4.RegSlice[0][i] = 1
	}

	symbolicDescription := make([][]int, 0)
	symbolicDescription = append(symbolicDescription, sre1.RegSlice...)
	symbolicDescription = append(symbolicDescription, sre2.RegSlice...)
	symbolicDescription = append(symbolicDescription, sre3.RegSlice...)
	symbolicDescription = append(symbolicDescription, reg4.RegSlice...)

	return symbolicDescription
}

// KnownPlaintextAttack
// Make 2^16 guesses and performs the Known Plaintext Attack.
// Prints the result
func KnownPlaintextAttack() {
	/*
		"For all possible 2^16 values of R4 solve the linearized system of equations that describe the output.
		Most of the 2^16-1 wrong solutions will be found by inconsistencies in Gauss Elimination.
		The solution of the equations will suggest the internal state of R1, R2, and R3.
		If more than one consistent internal state exists then do trial encryptions." - E. Barkan, E. Biham, and N. Keller
	*/

	r4Found := make([][]int, 0) // append results to this boi
	r4Guess := make([]int, 17)

	session_key = make([]int, 64) //all zero session key
	original_frame_number = 42
	r4Real, realKey, r4ForTest := MakeRealKeyStreamFourFrames(original_frame_number)

	realIteration := CalculateRealIteration(r4Real)
	lower := realIteration - 150
	upper := realIteration + 150
	fmt.Printf("real: %d, lower: %d, upper: %d\n", realIteration, lower, upper)

	guesses := int(math.Pow(2, 16))
	println(guesses)
	for i := lower; i < upper; i++ {
		// for i := 0; i < guesses; i++ {
		if i%100 == 0 {
			fmt.Printf("iteration %d \n", i)
		}
		if i == realIteration {
			fmt.Printf("iteration %d\n", realIteration)
		}
		if i == realIteration+1 {
			fmt.Printf("iteration %d\n", realIteration+1)
		}
		original_frame_number = 42 // reset the frame number for the symbolic version
		current_frame_number = 42

		r4Guess = MakeR4Guess(i) // for all possible value of r4 we need three frames
		r4Guess = PutConstantBackInRes(r4Guess, 10)

		// do this such that r4 guess can be copied into sr4 in SymSetRegisters()
		r4 = MakeR4()
		copy(r4.RegSlice, r4Guess)
		key1 := MakeSymKeyStream() // this clocks sr4 which has r4Guess as its array

		current_frame_number++

		// update r4Guess with new frame value
		r4 = MakeR4()
		copy(r4.RegSlice, r4Guess)

		frameInfluencedBits := SimulateClockingR4WithFrameDifference(original_frame_number, current_frame_number)
		r4.RegSlice = XorSlice(frameInfluencedBits, r4Guess)

		r4.RegSlice[10] = 1

		key2 := MakeSymKeyStream() //this will now copy the updated r4_regSlice into sr4

		current_frame_number++
		r4 = MakeR4()
		fakeR4 := MakeR4()
		copy(r4.RegSlice, r4Guess)
		diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame_number)

		//fakeR4.RegSlice clockes således at det er [...1..] de steder hvor diff påvirker indgangene
		for i := 0; i < 22; i++ {
			Clock(fakeR4)
			fakeR4.RegSlice[0] = fakeR4.RegSlice[0] ^ diff[i]
		}

		r4.RegSlice = XorSlice(fakeR4.RegSlice, r4.RegSlice)
		r4.RegSlice[10] = 1
		key3 := MakeSymKeyStream()
		current_frame_number++

		key := append(key1, key2...)
		key = append(key, key3...)

		// this returns a gauss struct
		gauss := solveByGaussEliminationTryTwo(key, realKey)

		if gauss.ResType == Error {
			continue
		} else if gauss.ResType == Multi {
			fmt.Printf("found multi in %d of lenght %d \n", i, len(gauss.Multi))
			for i := 0; i < len(gauss.Multi); i++ {
				if VerifyKeyStream(gauss.Multi[i]) {
					// If a solution is verified it is added to a list of verified guesses
					r4Found = append(r4Found, r4Guess)
				}
			}
		}
	}

	// 'trial encryptions'
	correctR4 := make([]int, len(r4Guess))
	numberOfValidR4 := len(r4Found)
	if numberOfValidR4 <= 0 {
		fmt.Printf("No solutions were found \n")
	} else if numberOfValidR4 > 1 {
		// we have multiple plausible solutions to r4
		for i := 0; i > numberOfValidR4; i++ {
			r4.RegSlice = r4Found[i]
			// We make an extra keystream with frame number: base frame + 4
			current_frame_number = original_frame_number + 4
			ks := MakeKeyStream()
			if reflect.DeepEqual(ks, r4ForTest) {
				fmt.Printf("This is the right one: %d\n", r4ForTest)
				correctR4 = r4Found[i]
				break
			}
		}
	} else {
		correctR4 = r4Found[0]
	}

	fmt.Printf("This is original r4:       %d\n", r4Real)
	for i := range r4Found {
		fmt.Printf("This is %d'th found r4:    %d\n", i, r4Found[i])
	}
	fmt.Println("Have we found the right r4?")
	if reflect.DeepEqual(correctR4, r4Real) {
		fmt.Println("Fuck yes we found it gutterne")
	} else {
		fmt.Println("RIP we dit not")
	}
}

func MakeR4Guess(number int) []int {
	r4Bit := make([]int, 16)

	for i := 0; i < 16; i++ {
		r4Bit[i] = (number >> i) & 1 // index 0 becomes least significant bit
	}

	return r4Bit
}

// VerifyKeyStream
// compares the found vars with the products that involves them and check that they match up.
func VerifyKeyStream(key []int) bool {
	// [vars1 | vars2 | vars3 | prod1 | prod2 | prod3 | b ]
	vars1_len := r1.Length - 1
	vars2_len := r2.Length - 1
	vars3_len := r3.Length - 1

	prod1_len := vars1_len * (vars1_len - 1) / 2
	prod2_len := vars2_len * (vars2_len - 1) / 2
	prod3_len := vars3_len * (vars3_len - 1) / 2

	prod1 := key[vars1_len+vars2_len+vars3_len : vars1_len+vars2_len+vars3_len+prod1_len]
	prod2 := key[vars1_len+vars2_len+vars3_len+prod1_len : vars1_len+vars2_len+vars3_len+prod1_len+prod2_len]
	prod3 := key[vars1_len+vars2_len+vars3_len+prod1_len+prod2_len : vars1_len+vars2_len+vars3_len+prod1_len+prod2_len+prod3_len]

	VerifiesProducts(key[0:vars1_len], prod1)
	VerifiesProducts(key[vars1_len:vars1_len+vars2_len], prod2)
	VerifiesProducts(key[vars1_len+vars2_len:vars1_len+vars2_len+vars3_len], prod3)

	return true
}

// VerifiesProducts
// checks if products in solved solutions fits the found variables.
func VerifiesProducts(vars []int, prods []int) bool {
	acc := 0
	for i := 0; i < len(vars); i++ {
		var_1 := vars[i]
		for j := i + 1; j < len(vars); j++ {
			var_2 := vars[j] // i and j runs over the vars variables
			if var_2*var_1 != prods[acc] {
				return false
			}
			acc++ //acc runs over the index in prod1
		}
	}

	return true
}

// ConvertBinaryToDecimal
// Returns an integer corresponding to the input bit slice
func ConvertBinaryToDecimal(number []int) int {
	bin_num := ""

	for i := len(number) - 1; i >= 0; i-- {
		bin_num = bin_num + strconv.Itoa(number[i])
	}

	num, err := strconv.ParseInt(bin_num, 2, 64)

	if err != nil {
		panic(err)
	}
	return int(num)
}

func RetrieveSessionKey(registers []int) []int {
	// FIXME: ???

	skey := make([]int, 0)

	symkey := DescribeRegistersFromKey() //
	gauss := solveByGaussEliminationTryTwo(symkey, registers)
	println(gauss.ResType)
	if gauss.ResType == Multi {
		skey = gauss.Multi[0]
	}

	return skey
}
