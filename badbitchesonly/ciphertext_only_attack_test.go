package main

import (
	"fmt"
	_ "fmt"
	"reflect"
	_ "reflect"
	"strconv"
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

	// make the message a slice slice
	longer_msg := make([][]int, 184)
	for i := 0; i < 184; i++ {
		longer_msg[i] = make([]int, 1)
		longer_msg[i][0] = msg[i]
	}

	// use error-correction on message
	G := CreateGMatrix()
	// correspons to M in text
	error_corrected_msg := MultiplyMatrix(G, longer_msg)

	KG := CreateKgMatrix()

	probertyOfInverseTransformation := MultiplyMatrix(KG, error_corrected_msg)
	//fmt.Printf("InvTrans: %d \n", probertyOfInverseTransformation)
	shouldBe := make([][]int, 272)
	for i := 0; i < 272; i++ {
		shouldBe[i] = make([]int, 1)
	}
	assert.Equal(t, probertyOfInverseTransformation, shouldBe)

	/**
	Does the same as TestMAKETEST from dumb_assversary
	*/

	// init r1, r2, r3, r4
	// makeRegisters() REVIEW: happens in makeKeyStream
	// set frame number
	current_frame_number, original_frame_number = 42, 42

	// session_key is now all 0's
	session_key = make([]int, 64)

	// init registers with sesion key and frame number
	// initializeRegisters() REVIEW: happens in makeKeyStream
	// setIndicesToOne() REVIEW: happens in makeKeyStream

	// init sr1, sr2, sr3
	//SymInitializeRegisters()

	keyStream := make([]int, 0)      // append to this, assert that the length is rigth
	symKeyStream := make([][]int, 0) // same here <3

	// how many frames do we need ?
	for i := 0; i < 4; i++ {
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

	/* Does new stuff not included in TestMAKETEST */

	// krypter error-corrected message med keystream

	// cipher text
	c := make([]int, 456)
	acc := "["
	for i := 0; i < 456; i++ {
		c[i] = error_corrected_msg[i][0] ^ keyStream[i]
		acc = acc + " " + strconv.Itoa(error_corrected_msg[i][0])
	}
	acc += "]"

	cipher_matrix := SliceToMatrix(c)
	KG_times_cipher := MultiplyMatrix(KG, cipher_matrix)
	//Vores konkrete bitvektor som skal gives som second argument til Gauss
	KG_C := MatrixToSlice(KG_times_cipher)
	printmatrix(KG_times_cipher)
	prints(KG_C, "KgC")

	// aaa := CalculateKgTimesSymKeyStream(KG, symKeyStream[:456]) //FIXME hvor meget af symKey skal vi have med

	// set up system
	// A := makeAugmentedMatrix(symKeyStream, keyStream)

	// K_gC := MultiplyMatrix(KG)
	// fmt.Printf("this is error corrected: \n%v\n ", acc)
	// fmt.Printf("this is k: \n%d\n", keyStream[:456])
	// fmt.Printf("this is c: \n%d\n", c)

	// kør 'doTheAttack'(?) fra dumb_assversry

	x := solveByGaussEliminationTryTwo(symKeyStream, keyStream)

	r1_solved, r2_solved, _ := MakeGaussResultToRegisters(x.Solved)

	if reflect.DeepEqual(r1_solved, r2_solved) {
		// t.Fails

	}

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
