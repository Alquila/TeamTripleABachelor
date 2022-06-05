package main

import (
	// "math/rand"
	// "reflect"
	// "strconv"
	// "strings"
	_ "encoding/binary"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

func TestDoTheSimpleHack1(t *testing.T) {
	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	reg := InitOneRegister()
	orgReg := make([]int, 19)
	copy(orgReg, reg.RegSlice)
	Bit_entry(symReg)

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSym(symReg)
	keyStream := SimpleKeyStream(reg)

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)
	print("Type is: " + res.ResType + "\n")
	r1_res := putConstantBackInRes(res.Solved, 15)

	fmt.Printf("length of res is: %d\n", len(res.Solved))
	// fmt.Printf("length of sym")
	// fmt.Printf("Res er: %d\n", res)
	// fmt.Printf("reg er: %d\n", orgReg)

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(r1_res, orgReg) {
		t.Fail()
		fmt.Printf("r1_res er: %d\n", r1_res)
		fmt.Printf("reg er:    %d\n", orgReg)
	}
}

func TestDoTheSimpleHackSecondVersion(t *testing.T) {

	// init one register, in both OG and sym version
	symReg := InitOneSymRegister()
	// Bit_entry(symReg) REVIEW: I moved this - Am
	reg := InitOneRegister()
	// orgReg is init, has entry for each variable, including the one set to 1
	orgReg := make([]int, 19)
	copy(orgReg, reg.RegSlice)
	Bit_entry(symReg)

	assert.Equal(t, orgReg, reg.RegSlice, "orgReg and reg are not the same")

	// make output keystream in both
	symKeyStream := SimpleKeyStreamSymSecondVersion(symReg)
	fmt.Printf("length of symKeyStream[0]: %d\n", len(symKeyStream[0]))
	keyStream := SimpleKeyStreamSecondVersion(reg)
	fmt.Printf("length of KeyStream: %d\n", len(keyStream))

	// make sym version into [][]int if not allready

	// use gauss to solve equations
	//res := solveByGaussElimination(symKeyStream, keyStream)
	// fmt.Printf("symKeyStream: \n%d\n", symKeyStream)
	res := solveByGaussEliminationTryTwo(symKeyStream, keyStream)
	fmt.Printf("Res Type: %v \n", res.ResType)

	r1_res := putConstantBackInRes(res.Solved[0:18], 15)

	fmt.Printf("længden af res er: %d\n", len(res.Solved))

	// compare if found res is equal to init registers
	if !reflect.DeepEqual(r1_res, orgReg) {
		t.Fail()
		fmt.Printf("Res er: %d\n", r1_res)
		fmt.Printf("reg er: %d\n", orgReg)
		t.Log("The result is wrong :(")
	}
	// fmt.Printf("reg er: %d\n", orgReg)
	// fmt.Printf("Res er: %d\n", res[0:19])
}

func TestPlaintextAttack(t *testing.T) {
	//orgReg := []int{0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0}
	//sr4.RegSlice = []int{0, 0, 0, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1}
	//orgReg := make([]int, 17) [0 0 0 1 1 1 1 1 0 0 0 0 0 0 0 1 1 1 0]
	//copy(orgReg, r4.RegSlice)

	session_key = make([]int, 64)
	original_frame_number = 55
	current_frame_number = 55

	sr1.ArrImposter = make([][]int, r1.Length)
	sr2.ArrImposter = make([][]int, r2.Length)
	sr3.ArrImposter = make([][]int, r3.Length)
	sr4.RegSlice = make([]int, r4.Length)

	res1, _, _ := DoTheKnownPlainTextHack()

	//fmt.Printf("len af res er: %d\n", len(res4)) // should be 656 as this is the number of unknown vars

	// offset := r1.Length + r2.Length + r3.Length
	// compare if found res is equal to init registers
	if !reflect.DeepEqual(res1, sr1.ArrImposter) {
		t.Fail()
		fmt.Printf("Res1 er   : %d\n", res1)
		// fmt.Printf("reg er: %d\n", orgReg)
		fmt.Printf("sr1.Arr er: %d\n", sr1.ArrImposter)
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

	// prints(firstSymReg.RegSlice[15], "række 15")
	// prints(firstSymReg.RegSlice[0], "række 0")
	// prints(firstSymReg.RegSlice[16], "række 16")

	firstSymReg.ArrImposter = DescribeNewFrameWithOldVariables(0, 1, firstSymReg)

	Bit_entry(firstSymReg)
	res := firstSymReg.ArrImposter
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
			shouldBe[i] = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		} else if i == 2 || i == 3 || i == 4 || i == 7 {
			shouldBe[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
			shouldBe[i][i] = 1
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
	println("res er")
	for i := 0; i < len(res); i++ {
		prints(res[i], "")
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
		if i == 0 || i == 2 || i == 4 || i == 5 || i == 6 || i == 7 {
			shouldBe[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
			shouldBe[i][i] = 1
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
	MakeRegisters()
	/* set frame number */
	current_frame_number = 42
	original_frame_number = 42

	// key := make([]int, 64)
	MakeSessionKey()

	/* init registers with key and framenumber*/
	InitializeRegisters()
	SetIndicesToOne()
	// fmt.Printf("This is r1 after init: \n%v\n", r1.RegSlice)

	/*save initial state registers*/
	old_r1 := make([]int, r1.Length)
	copy(old_r1, r1.RegSlice)
	old_r2 := make([]int, r2.Length)
	copy(old_r2, r2.RegSlice)
	old_r3 := make([]int, r3.Length)
	copy(old_r3, r3.RegSlice)
	old_r4 := make([]int, r4.Length)
	copy(old_r4, r4.RegSlice)

	fmt.Printf("This is old_r1 after init: \n%v\n", old_r1)
	fmt.Printf("This is old_r2 after init: \n%v\n", old_r2)
	fmt.Printf("This is old_r3 after init: \n%v\n", old_r3)
	fmt.Printf("This is old_r4 after init: \n%v\n", old_r4)

	/*should init the SymRegisters to ~I with bit in the last entry. Sr4 has copy of r4.RegSlice */
	SymInitializeRegisters()

	assert.Equal(t, old_r1, r1.RegSlice, "r1")
	assert.Equal(t, old_r2, r2.RegSlice, "r2")
	assert.Equal(t, old_r3, r3.RegSlice, "r3")
	assert.Equal(t, old_r4, r4.RegSlice, "r3")

	keyStream1 := make([]int, 228)
	keyStreamSym1 := make([][]int, 228)
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		ClockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		ClockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream1[i] = MakeFinalXOR()
		keyStreamSym1[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	//Do it all again
	current_frame_number++
	InitializeRegisters()
	SetIndicesToOne()

	SymInitializeRegisters()

	keyStream2 := make([]int, 228)
	keyStreamSym2 := make([][]int, 228)
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same pt 2")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		ClockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same pt 2")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		ClockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream2[i] = MakeFinalXOR()
		keyStreamSym2[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}
	fmt.Printf("keystreamsym %d \n", len(keyStreamSym2[0]))
	//Do it all again
	current_frame_number++
	InitializeRegisters()
	SetIndicesToOne()

	SymInitializeRegisters()

	keyStream3 := make([]int, 228)
	keyStreamSym3 := make([][]int, 228)
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same pt 3")

	/* Run A5/2 for 99 clocks and ignore output */
	for i := 0; i < 99; i++ {
		ClockingUnit(r4)
		Clock(r4)
		SymClockingUnit(sr4)
		Clock(sr4)
	}
	assert.Equal(t, r4.RegSlice, sr4.RegSlice, "R4 and SR4 are not the same pt 3")

	/* Run A5/2 for 228 clocks and use outputs as key-stream */
	for i := 0; i < 228; i++ {
		ClockingUnit(r4)
		SymClockingUnit(sr4)
		Clock(r4)
		Clock(sr4)
		keyStream3[i] = MakeFinalXOR()
		keyStreamSym3[i] = SymMakeFinalXOR(sr1, sr2, sr3)
	}

	A := append(keyStreamSym1, keyStreamSym2...)
	A = append(A, keyStreamSym3...)
	// A = append(A, A3...)

	b := append(keyStream1, keyStream2...)
	b = append(b, keyStream3...)

	x := solveByGaussEliminationTryTwo(A, b)
	println(x.ResType)
	// prints(x[0:20], "res 20xx")
	println(len(x.Multi))
	assert.Equal(t, true, VerifyKeyStream(x.Multi[0]), "VerifyKeyStream returned false")
	// println(x.Multi[0])
	r1_solved, r2_solved, r3_solved := MakeGaussResultToRegisters(x.Multi[0])

	// after_init := []int{1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1}
	// assert.Equal(t, after_init, old_r1)
	assert.Equal(t, r1_solved[15], 1)
	assert.Equal(t, r2_solved[16], 1)
	assert.Equal(t, r3_solved[18], 1)
	if !reflect.DeepEqual(r1_solved, old_r1) {
		t.Fail()
		fmt.Printf("r1_solved er: %d\n", r1_solved)
		fmt.Printf("old_r1 er   : %d\n", old_r1)
		// fmt.Printf("x er: %d\n", x[0:19])
	}

	if !reflect.DeepEqual(r2_solved, old_r2) {
		fmt.Printf("r2_solved er: %d\n", r2_solved)
		fmt.Printf("old_r2 er   : %d\n", old_r2)
		t.Fail()
	}

	if !reflect.DeepEqual(r3_solved, old_r3) {
		fmt.Printf("r3_solved er: %d\n", r3_solved)
		fmt.Printf("old_r3 er   : %d\n", old_r3)
		t.Fail()
	}

}

func TestDescribeSimpleSymWithFrame(t *testing.T) {

	sreg := SymMakeRegister(19, []int{18, 17, 16, 13}, []int{12, 15}, 14, 15) // equvalent to reg1
	for i := 0; i < 19; i++ {
		// reg.RegSlice[i] = make([]int, 19)
		sreg.ArrImposter[i][i] = 1 // each entry in the diagonal set to 1 as x_i is only dependent on x_i when initialized
	}

}

func TestMakeGaussResToRegisters(t *testing.T) {
	MakeRegisters()
	SymSetRegisters()

	res := make([]int, 0, 61)

	for i := 0; i < 15; i++ {
		res = append(res, i)
	}
	for i := 16; i < 19; i++ {
		res = append(res, i)
	}

	for i := 0; i < sr2.set1; i++ {
		res = append(res, i)
	}
	for i := sr2.set1 + 1; i < r2.Length; i++ {
		res = append(res, i)
	}

	for i := 0; i < sr3.set1; i++ {
		res = append(res, i)
	}
	for i := sr3.set1 + 1; i < r3.Length; i++ {
		res = append(res, i)
	}

	for i := 0; i < 10; i++ {
		res = append(res, i)
	}
	for i := 11; i < r4.Length; i++ {
		res = append(res, i)
	}
	// prints(res, "")
	r1s, r2s, r3s := MakeGaussResultToRegisters(res)
	// prints(r1s, "")	// prints(r2s, "")	// prints(r3s, "")
	r1shouldbe := stringToIntArray("0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 1 16 17 18")
	r2shouldbe := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 1, 17, 18, 19, 20, 21}
	r3shouldbe := stringToIntArray("0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 1 19 20 21 22")
	if !reflect.DeepEqual(r2s, r2shouldbe) {
		t.Fail()
		fmt.Printf("r2s er   : %d\n", r2s)
		fmt.Printf("shouldbe: %d\n", r2shouldbe)
	}
	if !reflect.DeepEqual(r1s, r1shouldbe) {
		t.Fail()
		fmt.Printf("r1s er   : %d\n", r1s)
		fmt.Printf("shouldbe: %d\n", r1shouldbe)
	}
	if !reflect.DeepEqual(r3s, r3shouldbe) {
		t.Fail()
		fmt.Printf("r2s er   : %d\n", r3s)
		fmt.Printf("shouldbe: %d\n", r3shouldbe)
	}

}

func TestPutConstantBackInRes(t *testing.T) {
	arr := make([]int, 10)

	assert.Equal(t, arr[2], 0)
	arr = putConstantBackInRes(arr, 2)
	assert.Equal(t, arr[2], 1)

	assert.Equal(t, arr[5], 0)
	arr = putConstantBackInRes(arr, 5)
	assert.Equal(t, arr[5], 1)

	MakeRegisters()
	SymSetRegisters()

	res := make([]int, 0, 19)

	for i := 0; i < 15; i++ {
		res = append(res, i)
	}
	for i := 16; i < 19; i++ {
		res = append(res, i)
	}
	// prints(res, "")
	res = putConstantBackInRes(res, 15)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 1, 16, 17, 18}, res)
	// prints(res, "")
}

// func TestMakeR4Guess(t *testing.T) { //outcommented becuase it doesn't really test anything, just to look in terminal

// 	// number := 0

// 	// for i := 0; i < int(math.Pow(2, 16)); i++ {
// 	// 	r4 := MakeR4Guess(i)
// 	// 	prints(r4, strconv.Itoa(i))
// 	// }
// 	r4 := MakeR4Guess(0)
// 	prints(r4, strconv.Itoa(0))
// 	prints(putConstantBackInRes(r4, 10), "with constant")

// 	r4 = MakeR4Guess(int(math.Pow(2, 16)) - 1)
// 	prints(r4, strconv.Itoa(0))
// 	prints(putConstantBackInRes(r4, 10), "with constant")

// }

func stringToIntArray(s string) []int {
	strs := strings.Split(s, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	return ary
}

func TestVerifyKeyStream(t *testing.T) {
	key := MakeLongIntSlice()
	// prints(key, "key")
	// fmt.Printf("len of key: %d \n", len(key))

	// VerifyKeyStream(key)
	fmt.Printf("%d \n", key[16])
	fmt.Printf("%d \n", key[17])
	fmt.Printf("%d \n", key[18+19])
	fmt.Printf("%d \n", key[18+20])
	// print(key[18+22+21+152+209])
	fmt.Printf("%d \n", key[18+21+22+153+209])

	//The above was for printing purpose only

	vars := []int{0, 1, 1, 0, 1, 0}
	prods := []int{0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 0}
	res := helper(vars, prods)
	if !res {
		t.Fail()
		t.Log("The products did not macth the vars")

	}

	MakeRegisters()
	key = make([]int, 655)
	key[16] = 1           //x17 = 1
	key[17] = 1           //x18 = 1
	key[18+21+22+152] = 1 //the place for x17*x18

	key[18+19] = 1            //y20
	key[18+20] = 1            //y21
	key[18+21+22+153+209] = 1 //the place for y20*y21

	key[18+21+20] = 1
	key[18+21+21] = 1
	key[18+21+22+153+210+230] = 1

	res = VerifyKeyStream(key)
	if !res {
		t.Fail()
		t.Log("The products did not macth the vars")

	}

}

func TestTryAllReg4(t *testing.T) {
	x := 33114
	prints(MakeR4Guess(x), "")
	prints(putConstantBackInRes(MakeR4Guess(x), 10), "")
	TryAllReg4()
}

func TestWhy(t *testing.T) {
	r4_sec_real := stringToIntArray("0 1 0 1 0 0 1 0 1 1 1 0 0 0 0 0 1")
	r4_sec_fake := stringToIntArray("0 1 1 0 1 0 0 0 0 1 1 0 0 0 1 1 1")
	r4_first_real := stringToIntArray("0 1 0 1 1 0 1 0 1 0 1 0 0 0 0 0 1")

	acc := make([]int, len(r4_sec_real))
	for i := 0; i < len(r4_sec_fake); i++ {
		acc[i] = r4_first_real[i] ^ r4_sec_real[i]
	}
	// diff between fake and real sec [0 0 1 1 1 0 1 0 1 0 0 0 0 0 1 1 0]
	prints(acc, "diff")

	original_frame_number = 42
	current_frame_number = 43
	diff := FindDifferenceOfFrameNumbers(original_frame_number, current_frame_number)
	prints(diff, "diff")
	iiii := MakeR4()
	for i := 0; i < 22; i++ {
		Clock(iiii)
		iiii.RegSlice[0] = iiii.RegSlice[0] ^ diff[i]
		// prints(iiii.RegSlice, strconv.Itoa(i))
	}
	prints(iiii.RegSlice, "will this work ")
	prints(XorSlice(r4_first_real, iiii.RegSlice), "?")
}

func TestBinaryConverter(t *testing.T) {
	x := 33114
	slice := (MakeR4Guess(x))
	intt := convertBinaryToDecimal(slice)
	assert.Equal(t, x, intt)

	slice = MakeR4Guess(10)
	intt = convertBinaryToDecimal(slice)
	assert.Equal(t, 10, intt)

	slice = MakeR4Guess(15)
	intt = convertBinaryToDecimal(slice)
	assert.Equal(t, 15, intt)

}

func TestDeep(t *testing.T) {
	slice := SliceToMatrix(MakeR4Guess(200))
	slice2 := SliceToMatrix(MakeR4Guess(200))
	if !reflect.DeepEqual(slice, slice2) {
		t.Fail()
	}
}
