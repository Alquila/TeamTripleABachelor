package main

import (
	"fmt"
	_ "fmt"
	"math"
	"reflect"
	_ "reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate_G_matrix(t *testing.T) {
	G := CreateGMatrix()

	for i := 0; i < 184; i++ {
		if G[i][i] != 1 {
			t.Fail()
			fmt.Printf("G[%d][%d] er 0 men burde være 1\n", i, i)
		}
	}
}

func TestCreate_K_Matrix(t *testing.T) {
	Kg := CreateKgMatrix()
	PrintMatrix(Kg)
	for i := 0; i < len(Kg[0]); i++ {
		if Kg[0][i] == 1 {
			fmt.Printf("1 at %d \n", i)
		}
	}
}

func TestMultiplyMatrix(t *testing.T) {
	A := make([][]int, 3)

	A[0] = []int{0, 1, 0, 1, 0}
	A[1] = []int{0, 0, 0, 1, 1}
	A[2] = []int{1, 1, 1, 0, 0}

	B := make([][]int, 5)

	B[0] = []int{1, 1, 1}
	B[1] = []int{0, 0, 0}
	B[2] = []int{1, 0, 1}
	B[3] = []int{0, 1, 0}
	B[4] = []int{0, 1, 1}

	res := MultiplyMatrix(A, B)
	shouldBe := make([][]int, 3)
	shouldBe[0] = []int{0, 1, 0}
	shouldBe[1] = []int{0, 0, 1}
	shouldBe[2] = []int{0, 1, 0}

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Logf("The result is not correct. Res is: %d", res)
	}
}

func TestMultiplyMatrix2(t *testing.T) {
	A := CreateKgMatrix()
	B := CreateGMatrix()

	res := MultiplyMatrix(A, B)

	shouldBe := make([][]int, 272)
	for i := 0; i < 272; i++ {
		shouldBe[i] = make([]int, 184)
	}
	fmt.Printf("Res size: %d x %d \n", len(res), len(res[0]))
	fmt.Printf("ShouldBe size: %d x %d \n", len(shouldBe), len(shouldBe[0]))

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Logf("The result is not correct. Res is: %d", res)
	}

}

func TestCiphertextOnlyAttackGauss(t *testing.T) {

	// create message to encrypt
	msg := CreateRandomMessage(184)
	// make the message a slice slice/matrix
	msgM := SliceToMatrix(msg)

	// create matrix used for error correction
	G := CreateGMatrix()

	// use error-correction on message
	// 'error_corrected_msg' correspons to M in text
	error_corrected_msg := MultiplyMatrix(G, msgM)                                     //456 x 1
	error_corrected_msg2 := MultiplyMatrix(G, SliceToMatrix(CreateRandomMessage(184))) //456 x 1
	error_corrected_msg3 := MultiplyMatrix(G, SliceToMatrix(CreateRandomMessage(184))) //456 x 1
	full_msg := append(error_corrected_msg, error_corrected_msg2...)
	full_msg = append(full_msg, error_corrected_msg3...)

	/**	Does the same as TestMAKETEST from dumb_assversary */

	// set frame number
	current_frame_number, original_frame_number = 42, 42
	// session_key is now all 0's
	session_key = make([]int, 64)
	// MakeSessionKey()
	keyStream := make([]int, 0)      // append to this, assert that the length is rigth
	symKeyStream := make([][]int, 0) // same here <3

	for i := 0; i < 6; i++ {
		newKeyStream := MakeKeyStream()
		SymInitializeRegisters()
		copy(sr4.RegSlice, r4_after_init.RegSlice)
		newSymKeyStream := ClockForKey(sr4)
		keyStream = append(keyStream, newKeyStream...)
		symKeyStream = append(symKeyStream, newSymKeyStream...)
		current_frame_number++
	}

	fmt.Printf("dims of msg:    %d x 1\n", len(full_msg))
	fmt.Printf("dims of Symkey: %d x %d\n", len(symKeyStream), len(symKeyStream[0]))
	fmt.Printf("dims of key	%d x 1\n", len(keyStream))

	c := make([]int, len(full_msg)) // cipher text = key xor msg
	for i := 0; i < len(c); i++ {
		c[i] = full_msg[i][0] ^ keyStream[i]
	}

	/* Create KG and multiply it with C */
	KG := CreateKgMatrix()
	probertyOfInverseTransformation := MultiplyMatrix(KG, error_corrected_msg) // sanity check
	shouldBe := make([][]int, 272)
	for i := 0; i < 272; i++ {
		shouldBe[i] = make([]int, 1)
	}
	assert.Equal(t, probertyOfInverseTransformation, shouldBe)
	assert.Equal(t, MultiplyMatrix(KG, error_corrected_msg2), shouldBe)
	assert.Equal(t, MultiplyMatrix(KG, error_corrected_msg3), shouldBe)
	fmt.Printf("dims of C	%d x 1\n", len(c))                //1368
	fmt.Printf("dims of KG	%d x %d\n", len(KG), len(KG[0])) //272 x 456

	/* Vores konkrete bitvektor som skal gives som second argument til Gauss */
	Prints(c[:456], "c")
	KG_C := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[:456])))
	Prints(KG_C, "KGc")
	Prints(c[184:456], "c[184:456]")
	print(len(c[184:456]), "c[184:456]")
	fmt.Printf("dims of K_G*C:  %d x 1 \n", len(KG_C)) //272 x 1
	KG_C2 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[456:912])))
	KG_C3 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[912:])))
	full_KGC := append(KG_C, KG_C2...)
	full_KGC = append(full_KGC, KG_C3...)

	/* Multiply KG with the SymbolicKeyStream to make KGK */
	KGk := CalculateKgTimesSymKeyStream(KG, symKeyStream[:456])
	KGk2 := CalculateKgTimesSymKeyStream(KG, symKeyStream[456:912]) //
	KGk3 := CalculateKgTimesSymKeyStream(KG, symKeyStream[912:])    //
	fmt.Printf("dims of K_g*k:  %d x %d \n", len(KGk), len(KGk[0])) //272 x 657
	full_KGk := append(KGk, KGk2...)
	full_KGk = append(full_KGk, KGk3...)
	fmt.Printf("dims of K_g*k4: %d x %d \n", len(full_KGk), len(full_KGk[0])) //816 x 657
	fmt.Printf("dims of full K_G*C:  %d x 1 \n", len(full_KGC))               //816 x1

	/* Try to solve KG*k = KG*C for V_f*/
	x := SolveByGaussElimination(full_KGk, full_KGC)
	println(x.ResType)
	fmt.Printf("Size of multi %d\n", len(x.Multi))
	fmt.Printf("Verifykeystream: %v\n", VerifyKeyStream(x.Multi[0]))
	r1_solved, r2_solved, r3_solved := MakeGaussResultToRegisters(x.Multi[0])
	Prints(r1_solved, "r1")
	Prints(r2_solved, "r2")
	Prints(r3_solved, "r3")

}

func TestCalculateKgTimesSymKeyStream(t *testing.T) {
	KG := make([][]int, 2)
	KG[0] = []int{1, 0, 1}
	KG[1] = []int{0, 1, 0}

	symkey := make([][]int, 3)
	symkey[0] = []int{1, 1, 1, 0}
	symkey[1] = []int{0, 1, 0, 1}
	symkey[2] = []int{1, 0, 1, 1}

	res := CalculateKgTimesSymKeyStream(KG, symkey)

	shouldBe := make([][]int, 2)
	shouldBe[0] = []int{0, 1, 0, 1}
	shouldBe[1] = []int{0, 1, 0, 1}

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Log("fuck")
	}

}

func TestSliceToMatrix(t *testing.T) {
	slice := stringToIntArray("0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1")
	slice_matrix := SliceToMatrix(slice)
	PrintMatrix(slice_matrix)
}

func TestMatrixToSlice(t *testing.T) {
	slice := stringToIntArray("0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1")
	matrix := SliceToMatrix(slice)
	backToSlice := MatrixToSlice(matrix)
	assert.Equal(t, slice, backToSlice)
}

func TestCiphertextOnlyAttack(t *testing.T) {
	start := time.Now()

	r4_found := make([][]int, 0)
	r4_guess := make([]int, 17)

	session_key = make([]int, 64) // FIXME: session_keyis all zeros now
	// MakeSessionKey()
	original_frame_number, current_frame_number = 42, 42
	// should have eight frames
	r4_bin, bin_key, key_for_test := MakeRealKeyStreamSixFrames(original_frame_number)

	real_iteration := CalculateRealIteration(r4_bin)
	lower := real_iteration - 100
	upper := real_iteration + 100
	fmt.Printf("real: %d, lower: %d, upper: %d\n", real_iteration, lower, upper)

	/* calculate ciphertext */
	c := CalculateXFrameCiphertext(bin_key, 6)
	c_for_test := CalculateXFrameCiphertext(key_for_test, 2)
	println(c[0] + c_for_test[0])
	/* Calculate KG*C */
	KG := CreateKgMatrix()
	KG_C := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[:456])))
	KG_C2 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[456:912])))
	KG_C3 := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[912:])))
	full_KGC := append(KG_C, KG_C2...)
	full_KGC = append(full_KGC, KG_C3...)

	guesses := int(math.Pow(2, 16))
	println(guesses)
	//for i := lower; i < upper; i++ {
	for i := 0; i < guesses; i++ {
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

		r4_guess = MakeR4Guess(i)
		r4_guess = PutConstantBackInRes(r4_guess, 10)

		symKeyStream := make([][]int, 0)

		for i := 0; i < 6; i++ { //Make six frame long sym-keystream for the guess
			frame_influenced_bits := SimulateClockingR4WithFrameDifference(original_frame_number, current_frame_number)
			r4.RegSlice = XorSlice(frame_influenced_bits, r4_guess)
			r4.RegSlice[10] = 1
			key1 := MakeSymKeyStream() //this clocks sr4 which has r4_guess as its array
			symKeyStream = append(symKeyStream, key1...)
			current_frame_number++
		}

		/* Multiply KG with the SymbolicKeyStream to make KGK */
		KGk := CalculateKgTimesSymKeyStream(KG, symKeyStream[:456])
		KGk2 := CalculateKgTimesSymKeyStream(KG, symKeyStream[456:912])
		KGk3 := CalculateKgTimesSymKeyStream(KG, symKeyStream[912:])
		full_KGk := append(KGk, KGk2...)
		full_KGk = append(full_KGk, KGk3...)

		x := SolveByGaussElimination(full_KGk, full_KGC)
		println(x.ResType)

		if x.ResType == Multi {
			//do stuff
			for i := 0; i < len(x.Multi); i++ {
				if VerifyKeyStream(x.Multi[i]) {
					r4_found = append(r4_found, r4_guess)
				}
			}

		}
		if x.ResType == Error {
			continue
		}

	}

	fmt.Printf("This is r4_found: %v\n", r4_found)
	// fmt.Printf("This is r4_guess: %v\n", r4_guess)
	fmt.Printf("This is r4_bin: %v\n", r4_bin)
	fmt.Printf("This is bin_key: %v\n", bin_key[:2])
	fmt.Printf("This is key_for_test: %v\n", key_for_test[:2])

	executionTime := time.Since(start)
	fmt.Printf("Ciphertext-only Attack took: %s", executionTime)
}

func TestCalculateXFramCiphertext(t *testing.T) {
	session_key = make([]int, 64)
	original_frame_number = 42
	current_frame_number = 42
	_, key, _ := MakeRealKeyStreamSixFrames(42)

	c := CalculateXFrameCiphertext(key, 6)
	assert.Equal(t, 1368, len(c))
	Prints(c, "c")

	c = CalculateXFrameCiphertext(key, 2)
	assert.Equal(t, 456, len(c))

}
