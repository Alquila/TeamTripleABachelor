package main

import (
	"fmt"
	_ "fmt"
	"reflect"
	_ "reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate_G_matrix(t *testing.T) {
	G := CreateGMatrix()

	// fmt.Printf("G Matrix: %d \n", G)
	for i := 0; i < 184; i++ {
		if G[i][i] != 1 {
			t.Fail()
			fmt.Printf("G[%d][%d] er 0 men burde være 1\n", i, i)
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

func TestCiphertextOnlyAttack(t *testing.T) {

	// create message to encrypt
	msg := createRandomMessage(184)
	// make the message a slice slice/matrix
	msgM := SliceToMatrix(msg)

	// create matrix used for error correction
	G := CreateGMatrix()

	// use error-correction on message
	// 'error_corrected_msg' correspons to M in text
	error_corrected_msg := MultiplyMatrix(G, msgM)                                     //456 x 1
	error_corrected_msg2 := MultiplyMatrix(G, SliceToMatrix(createRandomMessage(184))) //456 x 1
	error_corrected_msg3 := MultiplyMatrix(G, SliceToMatrix(createRandomMessage(184))) //456 x 1
	full_msg := append(error_corrected_msg, error_corrected_msg2...)
	full_msg = append(full_msg, error_corrected_msg3...)

	/**	Does the same as TestMAKETEST from dumb_assversary */

	// set frame number
	current_frame_number, original_frame_number = 42, 42
	// session_key is now all 0's
	session_key = make([]int, 64)
	keyStream := make([]int, 0)      // append to this, assert that the length is rigth
	symKeyStream := make([][]int, 0) // same here <3

	// how many frames do we need ?
	for i := 0; i < 6; i++ {
		// handle new frame variables ?
		newKeyStream := makeKeyStream()
		SymInitializeRegisters()
		copy(sr4.ArrImposter, r4_after_init.ArrImposter)
		newSymKeyStream := ClockForKey(sr4)
		assert.Equal(t, sr4.ArrImposter, r4.ArrImposter)
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
	KG_C := MatrixToSlice(MultiplyMatrix(KG, SliceToMatrix(c[:456])))
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
	// fmt.Printf("dims of K_g*k2: %d x %d \n", len(KGk2), len(KGk2[0])) //272 x 657
	full_KGk := append(KGk, KGk2...)
	full_KGk = append(full_KGk, KGk3...)
	fmt.Printf("dims of K_g*k4: %d x %d \n", len(full_KGk), len(full_KGk[0])) //816 x 657
	fmt.Printf("dims of full K_G*C:  %d x 1 \n", len(full_KGC))               //816 x1

	/* Try to solve KG*k = KG*C for V_f*/
	x := solveByGaussEliminationTryTwo(full_KGk, full_KGC)
	println(x.ResType)
	fmt.Printf("Size of multi %d\n", len(x.Multi))
	fmt.Printf("Verifykeystream: %v\n", VerifyKeyStream(x.Multi[0]))
	r1_solved, r2_solved, r3_solved := MakeGaussResultToRegisters(x.Multi[0])
	prints(r1_solved, "r1")
	prints(r2_solved, "r2")
	prints(r3_solved, "r3")

	// if reflect.DeepEqual(r1_solved, r2_solved) {
	// t.Fails

	// }

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
	printmatrix(slice_matrix)
}

func TestMatrixToSlice(t *testing.T) {
	slice := stringToIntArray("0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1")
	matrix := SliceToMatrix(slice)
	backToSlice := MatrixToSlice(matrix)
	assert.Equal(t, slice, backToSlice)
}

func TestTryAllCombinationsOfR4(t *testing.T) {
	r4_found := make([][]int, 0)
	r4_guess := make([]int, 17)

	session_key = make([]int, 64) // FIXME: session_keyis all zeros now
	original_frame_number, current_frame_number = 42, 42
	r4_bin, bin_key, r4_for_test := MakeRealKeyStreamFourFrames(original_frame_number)

	fmt.Printf("This is r4_found: %d\n", r4_found)
	fmt.Printf("This is r4_guess: %d\n", r4_guess)
	fmt.Printf("This is r4_bin: %d\n", r4_bin)
	fmt.Printf("This is bin_key: %d\n", bin_key)
	fmt.Printf("This is r4_for_test: %d\n", r4_for_test)
}
