package main

import (
	// "math/rand"
	// "reflect"
	// "strconv"
	// "strings"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/tidwall/assert"
	// "time"
	// //"golang.org/x/tools/go/analysis/passes/nilfunc"
)

// func TestPlaintext(t *testing.T) {
// 	plaintext := MakePlaintext()
// 	fmt.Printf("%d", plaintext)
// }

// func TestEncryptPlaintext(t *testing.T) {
// 	plaintext := MakePlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// 	cipher := EncryptSimplePlaintext(plaintext)
// 	fmt.Printf("%d \n", cipher)
// }

// func TestSymPlaintext(t *testing.T) {
// 	plaintext := MakeSymPlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// }

// func TestSymEncryptPlaintext(t *testing.T) {
// 	plaintext := MakeSymPlaintext()
// 	fmt.Printf("This is the plaintext: %d \n", plaintext)
// 	cipher := EncryptSimpleSymPlaintext()
// }

func NotAllowedBigBangTest(t *testing.T) {
	//plaintext := MakePlaintext()

}

// func TestDoTheSimpleHack(t *testing.T) {
// 	doTheSimpleHack()
// }

func TestPrint2(t *testing.T) {
	print("hello worlds")
}

func TestDoTheSimpleHack(t *testing.T) {
	// TODO: make this a function in dumb_assversary.go plz

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

	// fmt.Printf("Res er: %d\n", res)
	// fmt.Printf("reg er: %d\n", orgReg)

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res, orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("reg er: %d\n", orgReg)
	}
}

func TestDoTheSimpleHackSecondVersion(t *testing.T) {

	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.ArrImposter)

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSymSecondVersion(symReg)
	fmt.Printf("length of symKeyStream[0]: %d\n", len(symKeyStream[0]))
	keyStream := SimpleKeyStreamSecondVersion(reg)
	fmt.Printf("length of KeyStream: %d\n", len(keyStream))

	// make sym version into [][]int if not allready

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)

	fmt.Printf("længden af res er: %d\n", len(res))

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res[0:19], orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res[0:19])
		fmt.Printf("reg er: %d\n", orgReg)
	}
	// fmt.Printf("reg er: %d\n", orgReg)
	// fmt.Printf("Res er: %d\n", res[0:19])
}

func TestPlaintextAttack(t *testing.T) {
	//orgReg := []int{0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0}
	//sr4.ArrImposter = []int{0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1}
	//orgReg := make([]int, 17) [0 0 0 1 1 1 1 1 0 0 0 0 0 0 0 1 1 1 0]
	//copy(orgReg, r4.ArrImposter)

	session_key = make([]int, 64)
	original_frame_number = 55
	current_frame_number = 55

	sr1.ArrImposter = make([][]int, r1.Length)
	sr2.ArrImposter = make([][]int, r2.Length)
	sr3.ArrImposter = make([][]int, r3.Length)
	sr4.ArrImposter = make([]int, r4.Length)

	res1, _, _, res4 := DoTheKnownPlainTextHack()

	//fmt.Printf("len af res er: %d\n", len(res4)) // should be 656 as this is the number of unknown vars

	// offset := r1.Length + r2.Length + r3.Length
	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res4, sr4.ArrImposter) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res4)
		fmt.Printf("reg1 er: %d\n", res1)
		// fmt.Printf("reg er: %d\n", orgReg)
		fmt.Printf("reg er: %d\n", sr4.ArrImposter)
	}
	// fmt.Printf("reg er: %d\n", orgReg)
	// fmt.Printf("Res er: %d\n", res[offset:offset+17])

}

func TestFindDiffOfFrameNumbers(t *testing.T) {

	res := FindDifferenceOfFrameNumbers(1, 2)

	// LEAST significant bit is at index 0, so the bit is kinda 'reversed'
	shouldBe := make([]int, 22)
	shouldBe[0] = 1
	shouldBe[1] = 1

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		fmt.Printf("Res er: %d\n", res)
		fmt.Printf("shouldBe er: %d\n", shouldBe)
	}
}

func TestDescribeNewFrameNumberWithOldVar(t *testing.T) {
	firstSymReg := InitOneSymRegister()
	Bit_entry(firstSymReg)
	// prints(firstSymReg.ArrImposter[15], "række 15")
	// prints(firstSymReg.ArrImposter[0], "række 0")
	// prints(firstSymReg.ArrImposter[16], "række 16")

	res := DescribeNewFrameWithOldVariables(0, 1, firstSymReg)

	// fmt.Printf("res er: \n%d \n", res)
	println("res er")
	for i := 0; i < len(res); i++ {
		prints(res[i], "")
	}
	fmt.Printf("res er %d \n", len(res))
	fmt.Printf("res[0] er %d \n", len(res[0]))
	PrettySymPrintSliceBit(res, 15)

	shouldBe := make([][]int, 19)
	for i := 0; i < 19; i++ {
		if i == 0 {
			shouldBe[i] = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else if i < 15 {
			shouldBe[i] = make([]int, 19)
			shouldBe[i][i] = 1
		} else if i == 15 {
			shouldBe[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else {
			shouldBe[i] = make([]int, 19)
			shouldBe[i][i-1] = 1
		}
	}
	// fmt.Printf("shouldBe: %d \n", shouldBe)
	println("shouldBe er")
	for i := 0; i < len(shouldBe); i++ {
		prints(shouldBe[i], "")
	}
	// shouldBe[0] = []int{""}
	// shouldBe[0]
	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Log("The result is not correct")
	}

}

func TestDescribeNewFrameWithVariables8And15(t *testing.T) {
	firstSymReg := InitOneSymRegister()
	Bit_entry(firstSymReg)

	res := DescribeNewFrameWithOldVariables(8, 15, firstSymReg)
	// fmt.Printf("res is: \n %d \n", res)
	println("res er")
	for i := 0; i < len(res); i++ {
		prints(res[i], "")
	}
	PrettySymPrintSliceBit(res, firstSymReg.set1)
	prints(res[0], "")
	shouldBe := make([][]int, 19)
	for i := 0; i < 19; i++ {
		if i == 0 {
			shouldBe[i] = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else if i == 1 {
			shouldBe[i] = []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else if i == 2 {
			shouldBe[i] = []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else if i < 15 {
			shouldBe[i] = make([]int, 19)
			shouldBe[i][i] = 1
		} else if i == 15 {
			shouldBe[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
		} else {
			shouldBe[i] = make([]int, 19)
			shouldBe[i][i-1] = 1
		}
	}
	PrettySymPrintSliceBit(shouldBe, firstSymReg.set1)
	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Log("The result is not correct")
	}

}

func TestFindDifferenceOfFrameNumbers(t *testing.T) {
	diff := FindDifferenceOfFrameNumbers(136, 1357)
	fmt.Printf("diff:  %d", diff)
}

func TestDescribeFrameWithOldVariables2(t *testing.T) {
	firstSymReg := InitOneSymRegister()
	Bit_entry(firstSymReg)

	res := DescribeNewFrameWithOldVariables(136, 1357, firstSymReg)

	println("res er")
	for i := 0; i < len(res); i++ {
		prints(res[i], "")
	}

	shouldBe := make([][]int, firstSymReg.Length)

	for i := range shouldBe {
		shouldBe[i] = make([]int, len(firstSymReg.ArrImposter[0]))
		copy(shouldBe[i], firstSymReg.ArrImposter[i])

		if i == 3 || i == 4 || i == 5 || i == 7 || i == 11 || i == 13 || i == 14 || i == 15 {
			shouldBe[i][len(firstSymReg.ArrImposter[0])-1] = 1
		}
	}

	println("shouldBe er")
	for i := 0; i < len(res); i++ {
		prints(shouldBe[i], "")
	}

	if !reflect.DeepEqual(res, shouldBe) {
		t.Fail()
		t.Log("The result is wrong")
	}

}

func TestAppendFunction(t *testing.T) {

	A1 := make([][]int, 3)
	A1[0] = []int{1, 1, 1}
	A1[1] = []int{1, 1, 1}
	A1[2] = []int{1, 1, 1}

	A2 := make([][]int, 3)
	A2[0] = []int{2, 2, 2}
	A2[1] = []int{2, 2, 2}
	A2[2] = []int{2, 2, 2}

	// A := make([][]int, 0)
	A := append(A1, A2...)
	// A = append(A, A2...)

	// b := make([]int, 684)
	// b = append(b, b1...)
	// b = append(b, b2...)
	// b = append(b, b3...)

	fmt.Printf("This is A: \n%d\n", A)
	fmt.Printf("This is A1: \n%d\n", A1)
	fmt.Printf("This is A2: \n%d\n", A2)
	// fmt.Printf("This is b: \n%d", b)
}

func TestMAKETEST(t *testing.T) {

	/*
		Make concrete instances of r1, r2, r3 and r4
		Then clock them with the first frame number and copy these registers <- these we want to recover
		initialize the sym registers with the first framenumber (will just be ish I matrix)
		force bits to 1

		run the 99 clocks and 288 clocks with both symbol and actual registers and save

		Take the concrete instances of r1, r2, r3 and r4 and xor with frame_number +1
		Describe sym registers with differences between frames
		Force bits to 1 (how does this affect the symbolic registers exactly?)

		run the 99 clocks and 288 clocks with both symbol and actual registers and save

		repeat

		Stuff it into Gauss
	*/

	/* init r1 r2 r3 r4 */
	makeRegisters()
	/* set frame number */
	current_frame_number = 42
	original_frame_number = 42

	key := make([]int, 64)
	session_key = key

	/* init registers with key and framenumber*/
	initializeRegisters()
	setIndicesToOne()
	// fmt.Printf("This is r1 after init: \n%v\n", r1.ArrImposter)

	/*save initial state registers*/
	old_r1 := make([]int, r1.Length)
	copy(old_r1, r1.ArrImposter)
	old_r2 := make([]int, r2.Length)
	copy(old_r2, r2.ArrImposter)
	old_r3 := make([]int, r3.Length)
	copy(old_r3, r3.ArrImposter)
	old_r4 := make([]int, r4.Length)
	copy(old_r4, r4.ArrImposter)

	fmt.Printf("This is old_r1 after init: \n%v\n", old_r1)
	fmt.Printf("This is old_r2 after init: \n%v\n", old_r2)
	fmt.Printf("This is old_r3 after init: \n%v\n", old_r3)
	fmt.Printf("This is old_r4 after init: \n%v\n", old_r4)

	/*should init the SymRegisters to ~I with bit in the last entry. Sr4 has copy of r4.ArrImposter */
	SymInitializeRegisters()

	assert.Equal(t, old_r1, r1.ArrImposter, "r1")
	assert.Equal(t, old_r2, r2.ArrImposter, "r2")
	assert.Equal(t, old_r3, r3.ArrImposter, "r3")
	assert.Equal(t, old_r4, r4.ArrImposter, "r3")

	keyStream1 := make([]int, 228)
	keyStreamSym1 := make([][]int, 228)
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		clockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		clockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream1[i] = makeFinalXOR()
		keyStreamSym1[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	//Do it all again
	current_frame_number++
	initializeRegisters()
	setIndicesToOne()

	SymInitializeRegisters()

	keyStream2 := make([]int, 228)
	keyStreamSym2 := make([][]int, 228)
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same pt 2")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		clockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same pt 2")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		clockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream2[i] = makeFinalXOR()
		keyStreamSym2[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	//Do it all again
	current_frame_number++
	initializeRegisters()
	setIndicesToOne()

	SymInitializeRegisters()

	keyStream3 := make([]int, 228)
	keyStreamSym3 := make([][]int, 228)
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same pt 3")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		clockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.ArrImposter, sr4.ArrImposter, "R4 and SR4 are not the same pt 3")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		clockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream3[i] = makeFinalXOR()
		keyStreamSym3[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	A := append(keyStreamSym1, keyStreamSym2...)
	A = append(A, keyStreamSym3...)
	// A = append(A, A3...)

	b := append(keyStream1, keyStream2...)
	b = append(b, keyStream3...)

	x := solveByGaussEliminationTryTwo(A, b)

	r1_solved, _, _, _ := MakeGaussResultToRegisters(x)

	if !reflect.DeepEqual(r1_solved, old_r1) {
		t.Fail()
		fmt.Printf("Res er: %d\n", r1_solved)
		fmt.Printf("old er: %d\n", old_r1)
		// fmt.Printf("x er: %d\n", x[0:19])
	}
}
